# Messaging demo

Simple demo to demostrate how to send and receive messages with go 
using RabbitMQ as broker.

## Configure

First you need to configure the broker, the connection strings have to
be updated with your broker's data:

```
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
``` 

Once the credentials and host-URL are configured for your broker, you
can proceed to send messages via:

```
go run send.go
```

And you can start message receivers/consumers with:

```
go run receive.go
```

