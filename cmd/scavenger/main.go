package main

import (
    "github.com/tenebresus/muninn/pkg/scavenger"
	"github.com/tenebresus/muninn/pkg/muninnmq"
)

func main() {

    scavengerq := muninnmq.Init("muninn.scavenger.queue")
    go scavengerq.Listen(scavenger.OnMsg)

    var forever chan struct{}
    <-forever
}
