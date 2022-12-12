cd kafka
install-kafka.sh
cd ..

sleep 15

cd spark
install-spark.sh

sleep 15

kubectl cp --namespace spark-ns job.py my-release-spark-master-0:/opt/bitnami/spark/job.py
kubectl exec -ti --namespace spark-ns my-release-spark-master-0 -- pip install py4j confluent-kafka
kubectl exec -ti --namespace spark-ns my-release-spark-master-0 -- spark-submit --master spark://my-release-spark-master-svc:7077 --packages org.apache.spark:spark-sql-kafka-0-10_2.12:3.3.1 /opt/bitnami/spark/job.py

sleep infinity