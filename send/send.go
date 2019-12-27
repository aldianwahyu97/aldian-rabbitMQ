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
	// ====================================
	// Open Connection
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Gagal Terkoneksi dengan RabbitMQ")
	defer conn.Close()
	// ====================================

	// ====================================
	// Open Channel
	ch, err := conn.Channel()
	failOnError(err,"Gagal Membuka Channel")
	defer ch.Close()
	// ====================================

	// ====================================
	// Channel Pertama digunakan untuk kirim pesan ke GOLANG
	q, err := ch.QueueDeclare(
		"message-from-golang", //Message Name
		false, // durable
		false, // delete when unused
		false, //exclusive
		false, //no-wait
		nil, // arguments
	)
	failOnError(err,"Gagal mendeklarasikan Queue")
	// ====================================

	// ====================================
	// Channel Kedua digunakan untuk kirim pesan ke PHP
	q2, err2 := ch.QueueDeclare(
		"message-from-php", //Message Name
		false, // durable
		false, // delete when unused
		false, //exclusive
		false, //no-wait
		nil, // arguments
	)
	failOnError(err2,"Gagal mendeklarasikan Queue")
	// ====================================

	// ====================================
	// Isi pesan yang akan didistribusikan 
	body := "Hello, This Message from Golang To GOLANG"
	body2 := "Hello, This Message from Golang To PHP"
	// ====================================

	// ====================================
	// Error Handling Channel 1 (Pesan untuk Golang)
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
	// ====================================

	// ====================================
	// Error Handling Channel 2 (Pesan untuk PHP)
	err2 = ch.Publish(
		"",     // exchange
		q2.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body: []byte(body2),
		},
	)
	failOnError(err, "gagal mengirim pesan")
	// ====================================
}