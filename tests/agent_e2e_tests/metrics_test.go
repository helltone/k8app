package main

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestMetricResourceAttributes(t *testing.T) {
	attributes := telemetryData.ResourceMetrics[0].Resource.Attributes

	assert.Equal(t, "test_deployment_environment", attributes["deployment.environment"])
	assert.Equal(t, "test_project", attributes["project"])
	assert.Equal(t, "test_license", attributes["license"])

	assert.Equal(t, "test_service_name", attributes["service"])
	assert.Equal(t, "test_service_name", attributes["service.name"])
	assert.Equal(t, "test_project", attributes["service.namespace"])
	assert.Equal(t, "dev", attributes["service.version"])
	assert.Regexp(t, regexp.MustCompile("test_service_name-\\w{8}"), attributes["service.instance_id"])

	assert.Equal(t, "test_project/test_license/test_deployment_environment", attributes["cluster"])
}
