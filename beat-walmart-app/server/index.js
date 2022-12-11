const express = require("express");
const app = express();
const http = require("http");
const cors = require("cors")
const server = http.createServer(app)
const { Kafka } = require('kafkajs')

const kafka = new Kafka({
  clientId: 'restserver',
  brokers: ['sw-kafka.kafka-ns.svc.cluster.local:9092'],
})

notifications = [];
seen_messages = [];
var start_index = 0;

const consumer = kafka.consumer({ groupId: 'test-group' + Math.random() })
const run = async () => {
    consumer.connect()
    consumer.subscribe({ topic: 'notifications', fromBeginning: true })
    consumer.run({
        eachMessage: async ({ topic, partition, message }) => {
            str_msg = message.value.toString();

            if (!seen_messages.includes(str_msg)) {
                console.log(str_msg);
                msg_components = str_msg.split(",");
                formatted_msg = {
                    "low_product": msg_components[2],
                    "state": msg_components[1],
                    "city": msg_components[0]
                }
                notifications.push(formatted_msg);
                seen_messages.push(str_msg);
            }
        },
    })
}

run().catch(e => console.error("Could not connect to the kafka bootstrap server ..."))


app.use(cors());

server.listen(3001, () => {
    console.log("SERVER IS RUNNING ON 3001")
})

app.get("/", (req, res) => {
    res.send("The server is up!");
})

app.get("/poll", (req, res) => {
    console.log(notifications);
    var length = notifications.length;

    res.json(notifications.slice(start_index, length));
    start_index = length;
});