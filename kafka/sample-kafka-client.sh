#!/bin/sh

# Delete the kafka client if it exists
kubectl -n kafka-ns delete pod sw-kafka-client
sleep 5

# Run the client container and sleep forever (to keep it running)
kubectl run sw-kafka-client --restart='Never' --image docker.io/bitnami/kafka:2.6.0-debian-10-r0 --namespace kafka-ns --command -- sleep infinity
sleep 5

# Get the container's shell
kubectl -n kafka-ns exec -it sw-kafka-client -- bash