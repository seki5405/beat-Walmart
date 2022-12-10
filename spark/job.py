from pyspark.sql import SparkSession
from kafka import KafkaProducer, KafkaConsumer
import time

prod = KafkaProducer(bootstrap_servers="sw-kafka.kafka-ns.svc.cluster.local:9092")
consumer = KafkaConsumer("market-kafka-test", bootstrap_servers="sw-kafka.kafka-ns.svc.cluster.local:9092")

with open("kafka_output.txt", "w") as outfile:
    counter = 0
    for msg in consumer:
        msg = [str(x) for x in msg]
        outfile.write(",".join(msg) + "\n")
        counter += 1

        if counter == 10:
            break

prod.send("logs", b"starting workers ...")
print("Hello 1")

spark = SparkSession \
    .builder \
    .appName("test") \
    .master("spark://my-release-spark-master-svc:7077") \
    .getOrCreate()

df = spark \
    .readStream \
    .format("kafka") \
    .option("kafka.bootstrap.servers", "sw-kafka.kafka-ns.svc.cluster.local:9092") \
    .option("subscribe", "market-kafka-test") \
    .option("includeHeaders", "true") \
    .option("startingOffsets", "latest") \
    .option("spark.streaming.kafka.maxRatePerPartition", "10") \
    .load()
 
var = df.writeStream.trigger(once=True).start(queryName='that_query', outputMode="append", format='memory')

df.show()
var.stop()
rows = df.collect()

with open("kafka_output_2.txt", "w") as outfile:
    for row in rows:
        outwrite = ", ".join([str(x) for x in row])
        outfile.write(outwrite + "\n")

print("Hello here before the logs")
prod.send("logs", b"ending workers ...")

print(df)