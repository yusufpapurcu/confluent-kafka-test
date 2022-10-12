## ==== Tilt make commands ====
cluster:
	ctlptl apply -f ./kafka/kind-cluster.yaml

tilt-up: cluster
	tilt up

tilt-down:
	tilt down

## ==== Kafka make commands ====
TOPIC=empty-topic-name

kafka-producer:
	@echo "🚀 Entering kafka-client..."
	kubectl exec -it deploy/kafka-client -- kafka-console-producer.sh --topic $(TOPIC)  --broker-list  kafka-0.kafka-headless.default.svc.cluster.local:9092

purge-topic:
	 - @kubectl exec -it deploy/kafka-client -- kafka-topics.sh --create --topic $(TOPIC) --bootstrap-server  kafka-0.kafka-headless.default.svc.cluster.local:9092
	 - @kubectl exec -it deploy/kafka-client -- kafka-topics.sh --delete --topic $(TOPIC) --bootstrap-server  kafka-0.kafka-headless.default.svc.cluster.local:9092
	 - @kubectl exec -it deploy/kafka-client -- kafka-topics.sh --create --topic $(TOPIC) --bootstrap-server  kafka-0.kafka-headless.default.svc.cluster.local:9092

create-topic:
	 - @kubectl exec -it deploy/kafka-client -- kafka-topics.sh --create --topic $(TOPIC) --bootstrap-server  kafka-0.kafka-headless.default.svc.cluster.local:9092

delete-topic:
	 - @kubectl exec -it deploy/kafka-client -- kafka-topics.sh --delete --topic $(TOPIC) --bootstrap-server  kafka-0.kafka-headless.default.svc.cluster.local:9092
