# Run make mount before deploy to avoid permission issues

.PHONY: mount deploy port-forward test

mount:
	@minikube mount data:/tmp/testdata --uid=10001 --gid=10001

deploy:
	@helm upgrade --install agent \
		--namespace monitoring \
		--create-namespace \
		-f ../../charts/agent/values.test.yaml \
		../../charts/agent

port-forward:
	@kubectl port-forward --namespace monitoring services/agent 24317:4317

test:
	@find data -name "*.json" | xargs rm
	@OTEL_EXPORTER_OTLP_ENDPOINT=127.0.0.1:24317 \
		go test -v -count 1 ./...