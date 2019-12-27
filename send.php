<?php
    require_once __DIR__ . '/vendor/autoload.php';
    use PhpAmqpLib\Connection\AMQPStreamConnection;
    use PhpAmqpLib\Message\AMQPMessage;

    $connection = new AMQPStreamConnection('localhost', 5672, 'guest', 'guest');
    $channel = $connection->channel();

    $channel->queue_declare('message-from-php', false, false, false, false);

    $msg = new AMQPMessage('Hello, This Message From PHP');
    $channel->basic_publish($msg, '', 'message-from-php');

    echo " [x] Sent 'Hello, this message From PHP!'\n";

    $channel->close();
    $connection->close();
?>