const express = require("express");
const app = express();
const http = require("http");
const {Server} = require("socket.io")
const cors = require("cors")
const server = http.createServer(app)

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