package scavenger

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/tenebresus/muninn/pkg/muninnmq"
)

func OnMsg(d amqp.Delivery) {

    var msg muninnmq.MunninmqData
    json.Unmarshal(d.Body, &msg)

    ip := string(msg.Ip)

    fmt.Println("Received scan action for: " + ip)
    dmidecode := scavenge(ip)

    send := muninnmq.MunninmqData{
        Action: "store",
        Ip: ip,
        Data: []byte(dmidecode),
    }
    sendable, _ := json.Marshal(send)

    incubatorq := muninnmq.Init("muninn.incubator.queue")
    incubatorq.Send(sendable)
    fmt.Println("Send dmidecode to incubator")

}

func scavenge(ip string) string {

    resp, _ := http.Get("http://" + ip +":8080/dmidecode")
    body, _ := io.ReadAll(resp.Body)
    return string(body)

}   
