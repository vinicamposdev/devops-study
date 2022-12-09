# Go Kafka App
## Project Setup
Start consumer:
```sh
docker exec -it gokafka bash 
go run cmd/consumer/main.go
```
Start producer:
```sh
docker exec -it gokafka bash 
go run cmd/producer/main.go
```

## Concepts
You can find out more on [concepts page](concepts/kafka.md)