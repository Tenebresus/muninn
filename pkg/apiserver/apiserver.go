package apiserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tenebresus/muninn/pkg/Muninnmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

type getReq struct {

    Ip string

}

var qmessages = make(chan amqp.Delivery)

func Run() {

    apiserverq := muninnmq.Init("muninn.apiserver.queue")
    go apiserverq.Listen(OnMsg)

    http.HandleFunc("GET /scavenge", getHosts)
    http.HandleFunc("GET /scavenge/{ip}", getHostScans)
    http.HandleFunc("POST /scavenge/{ip}", scanHost)
    http.HandleFunc("GET /scavenge/{ip}/{timestamp}", getHostScanByTS)

    http.ListenAndServe(":8081", nil)
   
}

func OnMsg(d amqp.Delivery) {

    qmessages <-d

}

func getHostScans(w http.ResponseWriter, r *http.Request) {

    incubatorq := muninnmq.Init("muninn.incubator.queue")

    ip := r.PathValue("ip")

    send := muninnmq.MunninmqData{

        Action: "getHostScans",
        Ip: ip,

    }

    fmt.Println("Got request to list all the hosts that muninn is watching")

    sendable, _ := json.Marshal(send)
    incubatorq.Send(sendable)

    retstruct := <-qmessages

    fmt.Println("Received response from incubator")

    var msg muninnmq.MunninmqData
    json.Unmarshal(retstruct.Body, &msg)

    fmt.Println(string(msg.Data))

    w.Write(msg.Data)

}

func getHosts(w http.ResponseWriter, r *http.Request) {
    
    incubatorq := muninnmq.Init("muninn.incubator.queue")

    send := muninnmq.MunninmqData{

        Action: "getHosts",

    }

    fmt.Println("Got request to list all the hosts that muninn is watching")

    sendable, _ := json.Marshal(send)
    incubatorq.Send(sendable)

    retstruct := <-qmessages

    fmt.Println("Received response from incubator")

    var msg muninnmq.MunninmqData
    json.Unmarshal(retstruct.Body, &msg)

    fmt.Println(string(msg.Data))

    w.Write(msg.Data)
}

func getHostScanByTS(w http.ResponseWriter, r *http.Request) {

    ip := r.PathValue("ip")
    timestamp := r.PathValue("timestamp")

    incubatorq := muninnmq.Init("muninn.incubator.queue")

    send := muninnmq.MunninmqData{

        Action: "getHostScanByTS",
        Ip: ip,
        Data: []byte(timestamp),

    }

    fmt.Println("Got request to retrieve dmidecode for: " + send.Ip)

    sendable, _ := json.Marshal(send)
    incubatorq.Send(sendable)

    fmt.Println("Send request to incubator")

    retstruct := <-qmessages

    fmt.Println("Received response from incubator")

    var msg muninnmq.MunninmqData
    json.Unmarshal(retstruct.Body, &msg)

    fmt.Println(string(msg.Data))

    w.Write(msg.Data)
}

func scanHost(w http.ResponseWriter, r *http.Request) {

    ip := r.PathValue("ip")

    fmt.Println("Got scan request for: " + ip)

    scavengerq := muninnmq.Init("muninn.scavenger.queue")

    send := muninnmq.MunninmqData{
        
        Action: "scan",
        Ip: ip,
        Data: []byte(""),

    }

    sendable, _ := json.Marshal(send)
    scavengerq.Send(sendable)

    fmt.Println("Sent scan action to the scavenger")

}
