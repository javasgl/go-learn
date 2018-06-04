/**
 * @author songgl
 * @since 2018-05-25 17:09
 */
package main

import (
    "github.com/streadway/amqp"
    "learnGo/rabbitmq/hello"
    "log"
    "bytes"
    "time"
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

    msgs, err := ch.Consume(queue.Name, "", true, false, false, false, nil)
    hello.FailOnError(err, "Failed to register a Consumer")
    forever := make(chan bool)

    go func() {

        for m := range msgs {
            log.Printf("Received a message: %s", m.Body)
            dotCount := bytes.Count(m.Body, []byte("."))
            t := time.Duration(dotCount)
            time.Sleep(t * time.Second)
            log.Printf("Done")
            m.Ack(false)
        }
    }()

    log.Printf("[*]Waiting for messages.To eixt press CTRL+C")

    <-forever
}
