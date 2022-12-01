#!/bin/sh

# Add bitnami repository just in case it is not already there
helm repo add bitnami https://charts.bitnami.com/bitnami
helm search repo | grep -E "kafka|zookeeper"

# Create namespace for kafka
kubectl create ns kafka-ns

# Install kafka w/ helm
helm install sw bitnami/kafka --namespace kafka-ns

