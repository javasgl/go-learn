/**
 * @author songgl
 * @since 2018-05-28 18:01
 */
package main

import (
    "github.com/streadway/amqp"
    "learnGo/rabbitmq/pubsub"
    "log"
)

func main() {

    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    pubsub.FailOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    ch, err := conn.Channel()
    pubsub.FailOnError(err, "Failed to open a channel")
    defer ch.Close()

    err = ch.ExchangeDeclare(
        "logs",
        "fanout",
        true,
        false,
        false,
        false,
        nil,
    )
    pubsub.FailOnError(err, "Failed to declare a exchange")

    queue, err := ch.QueueDeclare(
        "",
        false,
        false,
        true,
        false,
        nil,
    )
    pubsub.FailOnError(err, "Failed to declare a queue")
    err = ch.QueueBind(
        queue.Name,
        "",
        "logs",
        false,
        nil,
    )
    pubsub.FailOnError(err, "Failed to bind a queue")
    msgs, err := ch.Consume(
        queue.Name,
        "",
        true,
        false,
        false,
        false,
        nil,
    )

    pubsub.FailOnError(err, "Failed to register a cousumer")
    forever := make(chan bool)
    go func() {
        for m := range msgs {
            log.Printf("[x] %s", m.Body)
        }
    }()

    log.Printf("[*] Waiting for logs.To exit press CTRL+C")
    <-forever
}
