# kafka
Created: 2022-11-30 10:21

## General Concepts
### Description
 - High throughput
 - Low latency (less then 2ms)
 - Scalable
 - Storage
 - High availability
 - Diversity of integrations and libs
 - Open-source
 - Not a Queue
 - Partitioned Log
### Dynamic
 - Has producer that owns a mensage
 - Has consumers is interested in mensages
 - Kafka is a set/cluster of machines/nodes/brokers
 - Each broker has a database
 - Kafka don't follow/distribute messages like pub sub
 - Consumers need to get the messages from queue
 - Zookeeper or the Kafka himself
	 - manager how this clusters communicate with each other
	 - healthy check
	 - manager containers
	 - service discovery
### Topics
 - it's not like a queue/exchange
 - communication channel responsible to receive and make available data sent to Kafka
 - Topic it is like a Log, a Partitioned Log
 - each message has an offset (id)
 - there is no problem consumer has another order
 - you can reprocess, messages can read and reread
 - Registry (e.g. offset 0)
	 - Headers -> metadata
	 - Key -> context of type of message group
	 - Value -> message content
	 - Timestamp -> creation date
### Partitions
 - divide reponsabilities
 - ensure distribution with resilience
 - don't put all eggs in same basket
 - if some partitions fall
	 - there others that will be elect
 - made load balancing
 - make fast speed reading
	 - parallel consuming
	 - more machines will manage consumer to read in partition/broker
	 - use round-robin
### Distributed Partitions
 - each broker has separated partition
 - Disavantage: If a broker fall, a partition is lost forever
 - Replicator Factor = 2
 - Advantage: ensure resilience if a broker fall a partition still available in others brokers
 - this decrease risk
 - ensure high availability
### Partition Leadership
 - the consumer will always read from leader partition
 - the copies is to ensure availabilitie
 - when a broker is down an election is made
 - an previous follower becomes new leader
 - if a replica is dow
### Message Delivery Ensurance
 - Ack 0 is a param that you don't need to return to make it right
	 - Fire and Forget (FF), None
	 - has some risks of lost messages
	 - can process more messages
	 - it's like UDP protocol idea
	 - Performatic
 - Ack 1 - Leader
	 - Replication and delivery the confirmation of message is sent from consumer to producer only by the leader
	 - Mid performer
 - Ack 1 - All
	 - All the replicas sent the delivery mensage, so you ensure that all the mensages are sent and replicated
	 - Poor in performance
### Keys
 - ensure delivery order
### Delivery Types
 - At most once
	 - best performance
	 - can lose some mensages
 - At least once
	 - mid performer
	 - may duplicate messages
 - Exactly once:
	 - worst performance
	 - dont duplicate
	 - ensure delivery of all messages
![[Screenshot 2022-12-07 at 06.32.21.png]]
### Producer: Indepontency
 - Indepontency is OFF
	 - Producer sent the same message duplicated
	 - Consumer read 2 times the same message
 - Producer is ON or indepotent
	 - Broker will discard the duplicated message
	 - Ensure Order,
	 - worst perfomance
### Consumer
 - 1 producer to 1 consumer
	 - slow in performance
 - Consumer Groups
	 - split partition between members of group
	 - common thing to have groups
 - Best performance:
	 - a consumer for each partition
 - Wasting resources:
	 - more consumer than partition, the rest will be iddle
### Create partitions in cli
 - Create
```sh
docker exec -it kafka-kafka-1 bash
kafka-topics --create --topic=test --bootstrap-server=localhost:9092 --partitions=3
```
 - List
```sh
kafka-topics --list --bootstrap-server=localhost:9092
```
 - get topic
```sh
kafka-topics --bootstrap-server=localhost:9092 --topic=test --describe
```
```
Topic: test     PartitionCount: 3       ReplicationFactor: 1    Configs: 
        Topic: test     Partition: 0    Leader: 1       Replicas: 1     Isr: 1
        Topic: test     Partition: 1    Leader: 1       Replicas: 1     Isr: 1
        Topic: test     Partition: 2    Leader: 1       Replicas: 1     Isr: 1
```
 - create a consumer:
```sh
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=test
```
 - create a producer, in another contaiter:
```sh
docker exec -it kafka-kafka-1 bash
kafka-console-producer --bootstrap-server=localhost:9092 --topic=test
```
 - get all the messages if a producer sent a message when consumer is down 
	 - without sort in messages
	 - don't have key
	 - each message go to different partitions
```sh
kafka-console-producer --bootstrap-server=localhost:9092 --topic=test --from-begining
```
 - create consumers with a groups:
```sh
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=test --group=x
```
 - monitor a consumer group
```sh
kafka-consumer-groups --bootstrap-server==localhost:9092 --group=x --describe
```
```
GROUP           TOPIC           PARTITION  CURRENT-OFFSET  LOG-END-OFFSET  LAG             CONSUMER-ID                                       HOST            CLIENT-ID
x               test            0          22              22              0               consumer-x-1-9ba01610-9f5e-4ef1-8eee-d6ddd83ea2ce /172.18.0.3     consumer-x-1
x               test            1          19              19              0               consumer-x-1-9f002ebb-9c29-4877-8311-72307ebd56c7 /172.18.0.3     consumer-x-1
x               test            2          26              26              0               consumer-x-1-ea92e38d-1477-4129-8344-08f1d643867f /172.18.0.3     consumer-x-1
```

## Integration 
 - [project repo](http://github.com/vinicamposdev/devops-study)
 - for the sake of example, will be use extra_hosts in docker
	 - you need to have 127.0.0.1 docker.host.internel inside of your /etc/hosts
 - access the go container
```sh
docker exec -it gokafka bash 
```
 - init go dependencies manager
```sh
go mod init github.com/vinicamposdev/devops-study
```
## Ecosystem
### ksqlDB
### Schema Registry
### Kafka Streams
### Kafka Connect
 - Integrate systems
 - Centralized Data Hub for sample integrations between 
	 - databases
	 - key-value stores
	 - search indexes
	 - file systems
 - What it is not
	 - ETL
	 - Complex Integration
	 - Apache Nifi
 - Time and effort that safes
#### Connector
 - Data Sources
	 - get data from some where
	 - example
		 - mysql
		 - mongo
		 - salesforce
 - Sinks
	 - where will throw the information
	 - example
		 - jdbc
		 - elastic search
		 - aws lambda
#### Standalone Workers
 - Its a node/machines
#### Distributed Workers
 - Kafka Connect Cluster
 - All partitions belongs to a 
 - All the clients helps to consume
 - Group Id
	 - will divide tasks to process faster
#### Converters
 - change data format for read/write on kafka
	 - Avro
	 - Protobuf
	 - JsonSchema
	 - Json
	 - String
	 - ByteArray
#### DLQ - Dead Letter Queue
 - Not on the format
 - Processing error
 - What to do when error occurs
 - property errors.tolerance
 - connect of Sync type
	 - none: make task failed imediatly
	 - all: error are ignored, process just continue
	 - errors.deadletterqueue.topic.name = \<topic-name\>
		 - on the message headers, get the message
		 - task is not stoped

## Misc
 - [[kafka-job-requirements]]
 - [[kafka-interview-questions]]
## References
1. Kafka
	1. [Confluent Hub Plugins](https://www.confluent.io/hub/)
	2. [Microservices with kafka](https://assets.confluent.io/m/3c05a1eaa11d258/original/20201028-WP-Event_Driven_Microservices.pdf)
2. Github
	1. * [microservice kafka java](https://github.com/ewolff/microservice-kafka)
	2. ** [kafka akka scala](https://github.com/lightbend/kafka-with-akka-streams-kafka-streams-tutorial)
	3. ** [kafka grpc go](https://github.com/AleksK1NG/Go-Kafka-gRPC-MongoDB-microservice)
	4. * [graphql kafka python](https://github.com/AlecAivazis/graphql-over-kafka)
	5. *** [clean arch kafka rabbit c#](https://github.com/phongnguyend/Practical.CleanArchitecture)
	6. ** [kafka cqrs grpc go](https://github.com/AleksK1NG/Go-CQRS-EventSourcing-Microservice)
	7. ** [footbal kafka cqrs java](https://github.com/djarza/football-events)
	8. ** [restaurant kafka kotlin](https://github.com/idugalic/digital-restaurant)
	9. * [kafka quarkus java](https://github.com/rmarting/kafka-clients-quarkus-sample)
	10. ** [kafka quarkus java coffee](https://github.com/cescoffier/reactive-coffeeshop-demo)
	11. *** [awesome kafka](https://github.com/semantalytics/awesome-kafka)
	12. *** [another awesome kafka](https://github.com/infoslack/awesome-kafka)
	13. *** [awesome arch kafka](https://awesome-architecture.com/messaging/kafka/#books)
	14. ** [kafka streams examples java](https://github.com/confluentinc/kafka-streams-examples)
	15. ** [kafka eventbus c#](https://github.com/mizrael/SuperSafeBank)
3. Misc
	1. #see-later [kafka blockchain](https://dzone.com/articles/apache-kafka-and-blockchain-friends-enemies-frenem)
	2. #see-later [aws msk](https://aws.amazon.com/msk/features/msk-serverless/)
	3. #see-later [aws mks](https://docs.aws.amazon.com/msk/latest/developerguide/serverless.html)
	4. 