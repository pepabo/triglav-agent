package main

import (
	"flag"
	"github.com/pepabo/triglav-agent"
)

var oneShot = flag.Bool(
	"one-shot",
	false,
	"Flag to toggle command mode or daemon mood (Default: false [daemon mode])",
)
var updateHostInterval = flag.Int(
	"update-host-interval",
	60 * 60,
	"Interval time of updating the info of the host (Default: 3,600 seconds)",
)

func init() {
	flag.Parse()
}

func main() {
	options := make(map[string]interface{})
	options["one-shot"] = *oneShot
	options["update-host-interval"] = *updateHostInterval

	agent := &triglav.Agent{}
	agent.Run(options)
}
