# Kafka Node

This is the node that will be used to run the Kafka cluster. It will be used to run the Kafka brokers and Zookeeper.

## How to deploy the Kafka Node and the Zookeeper
`./install-kafka.sh`

## How to deploy sample client (It could be used as a producer or consumer)
`./sample-kafka-client.sh`

## Sample commands to use on client container are explained in 
`sample-commands.txt`

## Workflow
#### 1. Deploy Kafka and Zookeeper
`./install-kafka.sh`

#### 2. Deploy sample client
`./sample-kafka-client.sh`

#### 3. Create topic (For now, it's `market-kafka-test`. Have to change later) -- (Run this on sample-client)
`kafka-topics.sh --create --topic market-kafka-test --bootstrap-server sw-kafka.kafka-ns.svc.cluster.local:9092`

#### 4. List topics to check if it's created properly -- (Run this on sample-client)
`kafka-topics.sh --list --bootstrap-server sw-kafka.kafka-ns.svc.cluster.local:9092`