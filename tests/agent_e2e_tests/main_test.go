package main

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"git.time2go.tech/go/telemetry"
	"git.time2go.tech/go/telemetry/log"
	"git.time2go.tech/go/telemetry/sdk/sdktrace"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
)

var telemetryData Telemetry

func TestMain(m *testing.M) {
	testId := uuid.NewString()
	fileName := fmt.Sprintf("data/otel-%s.json", testId)
	var err error = nil

	telemetryData, err = writeWaitReadData(testId, fileName, 15*time.Second)
	if err != nil {
		panic(err)
	}

	m.Run()
}

func writeWaitReadData(testId, fileName string, timeout time.Duration) (Telemetry, error) {
	err := writeData(testId)
	if err != nil {
		return Telemetry{}, err
	}

	err = waitData(fileName, timeout)
	if err != nil {
		return Telemetry{}, err
	}

	return readData(fileName)
}

func writeData(testId string) error {
	ctx := context.Background()
	tel, err := telemetry.New(
		ctx,
		telemetry.WithServiceName("test_service_name"),
		telemetry.WithResourceAttributes(attribute.String("test_id", testId)),
	)
	if err != nil {
		return err
	}

	tel = tel.Named("test")
	defer tel.Close(ctx)

	tel.Info(ctx, "test_log", log.Int("test_revision", 123))

	ctr, err := tel.Int64Counter("test_count")
	if err != nil {
		return err
	}
	ctr.Add(ctx, 1)

	ctx, span, tel := tel.Start(
		sdktrace.ContextWithSpanSampling(ctx, true),
		"test_span",
	)
	span.End()

	return nil
}

func waitData(fileName string, timeout time.Duration) error {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	timeoutChan := time.After(timeout)

	for {
		select {
		case <-timeoutChan:
			return errors.New(fmt.Sprintf("no data received after %d seconds", int(timeout.Seconds())))
		case <-ticker.C:
			stat, err := os.Stat(fileName)

			_ = stat
			if err == nil {
				return nil
			}
		}
	}
}

func readData(fileName string) (Telemetry, error) {
	jsonFile, err := os.Open(fileName)
	defer jsonFile.Close()
	if err != nil {
		return Telemetry{}, err
	}

	scanner := bufio.NewScanner(jsonFile)
	var logs []ResourceLog
	var metrics []ResourceMetric
	var spans []ResourceSpan

	for scanner.Scan() {
		var telemetryLine Telemetry
		err = json.Unmarshal(scanner.Bytes(), &telemetryLine)
		if err != nil {
			return Telemetry{}, err
		}

		if telemetryLine.ResourceLogs != nil {
			logs = append(logs, telemetryLine.ResourceLogs...)
		}

		if telemetryLine.ResourceMetrics != nil {
			metrics = append(metrics, telemetryLine.ResourceMetrics...)
		}

		if telemetryLine.ResourceSpans != nil {
			spans = append(spans, telemetryLine.ResourceSpans...)
		}
	}

	return Telemetry{
		ResourceLogs:    logs,
		ResourceMetrics: metrics,
		ResourceSpans:   spans,
	}, nil
}
