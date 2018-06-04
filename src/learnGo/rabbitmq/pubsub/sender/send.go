/**
 * @author songgl
 * @since 2018-05-28 18:02
 */
package main

import (
    "github.com/streadway/amqp"
    "learnGo/rabbitmq/pubsub"
    "os"
    "log"
)

func main() {

    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
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
    body := pubsub.BodyFrom(os.Args[1:])
    err = ch.Publish(
        "logs",
        "",
        false,
        false,
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(body),
        },
    )
    pubsub.FailOnError(err, "Failed to publish a message")
    log.Printf("[x] send %s", body)
}