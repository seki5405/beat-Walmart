from pyspark.sql import SparkSession
from confluent_kafka import Consumer, Producer
import json

def kafka_message_loop(spark_session, consumer, producer, WINDOW, SCHEMA, THRESHOLD):
    while True:
        print(f"Reading {WINDOW} messages ...")

        msg_array = consumer.consume(WINDOW)
        data_array = [json.loads(x.value().decode("utf-8")) for x in msg_array]

        rdd_data = []

        for data in data_array:
            current_city = data["City"]
            current_state = data["State"]

            for product in data["Cart"]:
                rdd_data.append((current_city, current_state, int(product), int(data["Cart"][product]), int(data["Inventory"][product])))

        df = spark_session.createDataFrame(rdd_data, SCHEMA)

        df = df.filter(df["PRODUCT"] != 0)
        df = df.groupBy(["CITY", "STATE", "PRODUCT"]).min("INVENTORY")
        df = df.filter(df["min(INVENTORY)"] <= THRESHOLD).select("CITY", "STATE", "PRODUCT")
        notifications = df.collect()

        print("Showing output ...")
        df.show()

        for notification in notifications:
            producer.produce("notifications", key="Key-B", value=",".join([str(x) for x in notification]))

        producer.flush()
        

def main():
    # Set up the configuration for the message queue using KAFKA
    CONF = {
        "bootstrap.servers": "sw-kafka.kafka-ns.svc.cluster.local:9092", 
        "group.id": "pyspark"
    }
    
    WINDOW = 50
    SCHEMA = ["City", "STATE", "PRODUCT", "SALE", "INVENTORY"]
    THRESHOLD = 9900

    spark = SparkSession \
    .builder \
    .appName("test") \
    .master("spark://my-release-spark-master-svc:7077") \
    .getOrCreate()


    consumer = Consumer(CONF)
    consumer.subscribe(["market-kafka-test"])

    producer = Producer(CONF)

    kafka_message_loop(spark, consumer, producer, WINDOW, SCHEMA, THRESHOLD)

main()