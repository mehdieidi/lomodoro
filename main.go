package main

import (
	"flag"
	"time"

	"github.com/mehdieidi/hey"
)

func main() {
	workDuration := flag.Int("w", 25, "how much minutes do you work?")
	breakDuration := flag.Int("b", 2, "how much minutes do you rest?")
	flag.Parse()

	ticker := time.NewTicker(time.Duration(*workDuration) * time.Minute)

	for {
		select {
		case <-ticker.C:
			topic := "break"
			text := "enough man! go get some rest."
			notify(topic, text)
		}

		time.Sleep(time.Duration(*breakDuration) * time.Minute)
	}
}

func notify(topic, text string) {
	hey.Push(hey.Notification{
		Title:    topic,
		Body:     text,
		AppName:  "lomodoro",
		Duration: 5 * time.Second,
	})
}
