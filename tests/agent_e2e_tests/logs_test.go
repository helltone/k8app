package main

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestLogResourceAttributes(t *testing.T) {
	attributes := telemetryData.ResourceLogs[0].Resource.Attributes

	assert.Equal(t, "test_deployment_environment", attributes["deployment.environment"])
	assert.Equal(t, "test_project", attributes["project"])
	assert.Equal(t, "test_license", attributes["license"])

	assert.Equal(t, "test_service_name", attributes["service"])
	assert.Equal(t, "test_service_name", attributes["service.name"])
	assert.Equal(t, "test_project", attributes["service.namespace"])
	assert.Equal(t, "dev", attributes["service_version"])
	assert.Regexp(t, regexp.MustCompile("test_service_name-\\w{8}"), attributes["service_instance_id"])

	assert.Equal(t, "test_project/test_license/test_deployment_environment", attributes["cluster"])

	assert.Equal(t, "127.0.0.1", attributes["k8s.pod.ip"])

	assert.Equal(t, "logfmt", attributes["loki.format"])
	assert.Equal(t, "aws_region, container, deployment_environment, license, namespace, pod, project, service", attributes["loki.resource.labels"])
}
