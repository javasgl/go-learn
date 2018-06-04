/**
 * @author songgl
 * @since 2018-05-25 16:47
 */
package main

import (
    "github.com/streadway/amqp"
    "learnGo/rabbitmq/hello"
    "os"
)

func main() {
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    hello.FailOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()
    ch, err := conn.Channel()
    hello.FailOnError(err, "Failed to open a channel")
    defer ch.Close()

    queue, err := ch.QueueDeclare("hello", false, false, false, false, nil)
    hello.FailOnError(err, "Failed to declare a queue")

    body := hello.BodyFrom(os.Args)
    err = ch.Publish(
        "",
        queue.Name,
        false,
        false,
        amqp.Publishing{
            DeliveryMode: amqp.Persistent,
            ContentType:  "text/plain",
            Body:         []byte(body),
        },
    )
    hello.FailOnError(err, "Failed to publish a message")
}
