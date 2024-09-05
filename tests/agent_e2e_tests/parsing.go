package main

import (
	"encoding/json"
	"strconv"
	"time"
)

type UnixTime time.Time

type Attributes map[string]interface{}

type Resource struct {
	Attributes Attributes `json:"attributes"`
}

type ResourceLog struct {
	Resource Resource `json:"resource"`
}

type ResourceMetric struct {
	Resource Resource `json:"resource"`
}

type ResourceSpan struct {
	Resource Resource `json:"resource"`
}

type Telemetry struct {
	ResourceLogs    []ResourceLog    `json:"resourceLogs"`
	ResourceMetrics []ResourceMetric `json:"resourceMetrics"`
	ResourceSpans   []ResourceSpan   `json:"resourceSpans"`
}

func (a *Attributes) UnmarshalJSON(b []byte) error {
	var attributes []interface{}
	var result = Attributes{}

	err := json.Unmarshal(b, &attributes)
	if err != nil {
		return err
	}

	for _, attrInterface := range attributes {
		attr, ok := attrInterface.(map[string]interface{})
		if !ok {
			panic("can't unmarshal attributes")
		}

		result[attr["key"].(string)] = attr["value"].(map[string]interface{})["stringValue"].(string)
	}
	*a = result

	return nil
}

func (t *UnixTime) UnmarshalJSON(b []byte) error {
	var ts string

	err := json.Unmarshal(b, &ts)
	if err != nil {
		return err
	}

	i, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return err
	}

	*t = UnixTime(time.Unix(0, i))

	return nil
}
