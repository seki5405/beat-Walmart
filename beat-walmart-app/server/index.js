const express = require("express");
const app = express();
const http = require("http");
const {Server} = require("socket.io")
const cors = require("cors")
const server = http.createServer(app)
const { Kafka } = require('kafkajs')

const kafka = new Kafka({
  clientId: 'restserver',
  brokers: ['sw-kafka.kafka-ns.svc.cluster.local:9092'],
})

const consumer = kafka.consumer({ groupId: 'test-group' })
consumer.connect()
consumer.subscribe({ topic: 'notifications', fromBeginning: false })

consumer.run({
    eachMessage: async ({ topic, partition, message }) => {
        console.log({
        value: message.value.toString(),
        })
    },
})

app.use(cors());

const io = new Server(server, {
    cors: {
        origin: "http://localhost:3000",
        methods: ["GET", "POST"]
    },
});

server.listen(3001, () => {
    console.log("SERVER IS RUNNING ON 3001")
})

app.get("/", (req, res) => {
    res.send("Something here")
})