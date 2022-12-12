const express = require("express");
const app = express();
const http = require("http");
const cors = require("cors")
const server = http.createServer(app)
const { Kafka } = require('kafkajs')

app.use(cors());

server.listen(3001, () => {
    console.log("SERVER IS RUNNING ON 3001")
})

app.get("/", (req, res) => {
    res.send("The server is up!");
})

app.get("/poll", (req, res) => {
    var dummy_object = [{"low_product":"1","state":"NY","city":"New York"},{"low_product":"1","state":"CA","city":"San Diego"},{"low_product":"1","state":"TX","city":"Austin"},{"low_product":"4","state":"CO","city":"Boulder"},{"low_product":"2","state":"WA","city":"Seattle"},{"low_product":"3","state":"CO","city":"Boulder"},{"low_product":"3","state":"WA","city":"Seattle"},{"low_product":"4","state":"CA","city":"San Diego"},{"low_product":"4","state":"TX","city":"Austin"},{"low_product":"4","state":"MN","city":"Minneapolis"},{"low_product":"3","state":"TX","city":"Austin"},{"low_product":"1","state":"CO","city":"Boulder"},{"low_product":"3","state":"CA","city":"San Diego"},{"low_product":"2","state":"MN","city":"Minneapolis"},{"low_product":"3","state":"MN","city":"Minneapolis"},{"low_product":"1","state":"MN","city":"Minneapolis"},{"low_product":"3","state":"NY","city":"New York"},{"low_product":"4","state":"NY","city":"New York"},{"low_product":"2","state":"CO","city":"Boulder"},{"low_product":"1","state":"WA","city":"Seattle"},{"low_product":"4","state":"WA","city":"Seattle"},{"low_product":"2","state":"CA","city":"San Diego"},{"low_product":"2","state":"TX","city":"Austin"},{"low_product":"2","state":"NY","city":"New York"}]

    res.json(dummy_object);

});
