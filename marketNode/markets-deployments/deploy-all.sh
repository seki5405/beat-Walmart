#!/bin/zsh

# This script is used to deploy the market nodes to the cloud.
# It is called by the deploy.sh script in the root directory.

# The script assumes that the following environment variables are set:
# - KAFKA_BROKER
# - CITY
# - STATE
# - BIAS
# - MAX_BUYING
# - DEFAULT_INVENTORY

kubectl apply -f boulder-co-deployment.yaml
kubectl apply -f austin-tx-deployment.yaml
kubectl apply -f mpls-mn-deployment.yaml
kubectl apply -f seattle-wa-deployment.yaml
kubectl apply -f newyork-ny-deployment.yaml
kubectl apply -f sandiego-ca-deployment.yaml