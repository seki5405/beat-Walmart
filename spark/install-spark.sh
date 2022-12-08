helm repo add bitnami https://charts.bitnami.com/bitnami
kubectl create ns spark-ns
helm install my-release bitnami/spark --namespace spark-ns