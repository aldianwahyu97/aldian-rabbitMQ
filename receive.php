<?php
    require_once __DIR__ . '/vendor/autoload.php';
    use PhpAmqpLib\Connection\AMQPStreamConnection;

    // Menghubungkan PHP dengan RabbitMQ: Parameter->IP, Port, Username, Password
    $connection = new AMQPStreamConnection('localhost', 5672, 'guest', 'guest'); 
    $channel = $connection->channel();

    // "message-from-golang" : Merupakan key dari pesan yang akan diterima dari mana 
    // Key pada Receiver harus sama dan sesuai dengan Key yang diberikan oleh Sender
    // Notes: Sender boleh bersalah dari program dengan bahasa apa pun, tidak mesti harus satu bahasa
    // PENTING! Sesuaikan Key Sender dengan Receiver agar pesan diterima oleh Receiver
    $channel->queue_declare('message-from-golang', false, false, false, false);

    echo " [*] Waiting for messages. To exit press CTRL+C\n";
    $callback = function ($msg) {
        echo ' [x] Received a Message: ', $msg->body, "\n";
    };
      
    $channel->basic_consume('message-from-golang', '', false, true, false, false, $callback);
    while ($channel->is_consuming()) {
        $channel->wait();
    }
?>