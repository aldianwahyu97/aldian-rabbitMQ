package main

import (
	"log"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string){
	if err != nil{
		log.Fatalf("%s: %s",msg,err)
	}
}

func main(){
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Gagal Terkoneksi dengan RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err,"Gagal Membuka Channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"message-from-golang", //Message Name
		false, // durable
		false, // delete when unused
		false, //exclusive
		false, //no-wait
		nil, // arguments
	)

	failOnError(err,"Gagal mendeklarasikan Queue")

	body := "Hello, This Message from Golang"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body: []byte(body),
		},
	)
	failOnError(err, "gagal mengirim pesan")
}