package main

import (
	"math/rand"
	"time"

	"github.com/Gimulator-Games/paper-soccer-random-agent/agent"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	agent, err := agent.NewAgent()
	if err != nil {
		panic(err)
	}
	agent.Listen()
}
