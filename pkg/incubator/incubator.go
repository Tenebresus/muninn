package incubator

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/tenebresus/muninn/pkg/muninnmq"
)

var storageDir string = "/tmp/muninn_storage"

func OnMsg(d amqp.Delivery) {

    var msg muninnmq.MunninmqData
    json.Unmarshal(d.Body, &msg)

    do(msg)

}

func do(msg muninnmq.MunninmqData) {

    action := msg.Action
    ip := msg.Ip
    data := msg.Data

    switch action {

    case "store":
        incubate(ip, data)

    case "getHostScanByTS":
        getHostScanByTS(ip, data)

    case "getHostScans":
        getHostScans(ip)

    case "getHosts":
        getHosts()

    }

}

func getHostScans(ip string) {

    validateDir(ip)
    ipDir := storageDir + "/" + ip + "/"

    ret := ""

    scans, _ := os.ReadDir(ipDir)

    for _, scan := range scans {
        ret += scan.Name() + "\n"
    }

    apiserverq := muninnmq.Init("muninn.apiserver.queue")
    send := muninnmq.MunninmqData{
        Data: []byte(ret),
    }
    sendable, _ := json.Marshal(send)

    apiserverq.Send(sendable)
    fmt.Println("Sent data to apiserver")
}

func getHosts() {

    items, _ := os.ReadDir(storageDir)

    ret := ""

    for _, item := range items {
        ret += item.Name() + "\n"
    }

    apiserverq := muninnmq.Init("muninn.apiserver.queue")

    send := muninnmq.MunninmqData{
        Data: []byte(ret),
    }
    sendable, _ := json.Marshal(send)

    apiserverq.Send(sendable)
    fmt.Println("Sent data to apiserver")

}

func getHostScanByTS(ip string, data []byte) {

    validateDir(ip)
    tsFile := storageDir + "/" + ip + "/" + string(data)

    content, err := os.ReadFile(tsFile) 

    if err != nil {
        content = []byte("File not found")
    }

    apiserverq := muninnmq.Init("muninn.apiserver.queue")

    send := muninnmq.MunninmqData{
        Data: []byte(content),
    }
    sendable, _ := json.Marshal(send)

    apiserverq.Send(sendable)
    fmt.Println("Sent data to apiserver")

}

func incubate(ip string, dmidecode []byte) {

    validateDir(ip)
    ipDir := storageDir + "/" + ip + "/"
    time := strconv.Itoa(int(time.Now().Unix()))

    incubateFile := ipDir + time 

    file, _ := os.Create(incubateFile)
    file.Write(dmidecode)
    fmt.Println("Succesfully wrote data to: " + incubateFile)

}   

func validateDir(ip string) {

    _, err := os.ReadDir(storageDir)

    if err != nil {
        os.Mkdir(storageDir, 0777)
    }

    ipDir := storageDir + "/" + ip
    _, err = os.ReadDir(ipDir)

    if err != nil {
        os.Mkdir(ipDir, 0777)
    }

}
