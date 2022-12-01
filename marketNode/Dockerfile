FROM golang:1.17.6
WORKDIR /marketnode
COPY . /marketnode
# RUN go mod init github.com/seki5405/kafka-go
RUN go mod tidy
RUN go build -o main .
CMD ["/marketnode/main"]