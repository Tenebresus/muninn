package main

import (
	"github.com/tenebresus/muninn/pkg/muninnmq"
    "github.com/tenebresus/muninn/pkg/incubator"
)

func main() {

    incubatorq := muninnmq.Init("muninn.incubator.queue")
    go incubatorq.Listen(incubator.OnMsg)

    var forever chan struct{}
    <-forever
}
