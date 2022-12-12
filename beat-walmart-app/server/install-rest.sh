kubectl create ns rest-ns
kubectl apply -n rest-ns -f server-deployment.yaml
kubectl apply -n rest-ns -f server-service.yaml