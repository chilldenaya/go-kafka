# go-kafka
Run Kafka locally via docker
```
docker compose up

// or

docker-compose up
```

Simulate microservices interaction to the kafka by doing this:
```
// run the producer
go run order/main.go

// run the first consumer
go run fulfillment/main.go

// run the second consumer
go run stock/main.go
```