from pyspark.sql import SparkSession

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
    .load()

with open("/tmp/someone.txt", "w"):
    pass

df.write.csv("./someoutput.csv")

print(df)