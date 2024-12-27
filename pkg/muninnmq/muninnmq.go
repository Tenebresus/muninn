package muninnmq

import (

	amqp "github.com/rabbitmq/amqp091-go"
)

var cacheq = make(map[string]Munninmq)

type MunninmqData struct {
    
    Action string
    Ip string
    Data []byte

}

type Munninmq struct {

    queue amqp.Queue
    channel *amqp.Channel
    connection *amqp.Connection

}

func (m Munninmq) Listen(OnMsg func(amqp.Delivery)) {

    msgs, _ := m.channel.Consume(m.queue.Name, "", true, false, false, false, nil)
    
    for d := range msgs {
        OnMsg(d)
    }

}

func (m Munninmq) Send(message []byte) {

    m.channel.Publish("", m.queue.Name, false, false, amqp.Publishing{
        ContentType: "text/plain",
        Body: message,
    })    

}

func Init(queueName string) Munninmq {

    queue, exists := cacheq[queueName]

    if (exists) {
        return queue
    }

    conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
    ch, _ := conn.Channel()
    q, _ := ch.QueueDeclare(queueName, false, false, false, false, nil)

    m := Munninmq{
        queue: q,
        channel: ch,
        connection: conn,
    }

    cacheq[queueName] = m

    return m

}

