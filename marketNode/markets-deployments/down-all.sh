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

kubectl delete -f boulder-co-deployment.yaml
kubectl delete -f austin-tx-deployment.yaml
kubectl delete -f mpls-mn-deployment.yaml
kubectl delete -f seattle-wa-deployment.yaml
kubectl delete -f newyork-ny-deployment.yaml
kubectl delete -f sandiego-ca-deployment.yaml