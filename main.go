package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/mehdieidi/hey"
)

func main() {
	workDuration := flag.Int("w", 25, "how much minutes do you work?")
	breakDuration := flag.Int("b", 2, "how much minutes do you rest?")
	flag.Parse()

	fmt.Printf("work: %d mins, break: %d mins\n", *workDuration, *breakDuration)

	ticker := time.NewTicker(time.Duration(*workDuration) * time.Minute)

	count := 1

	for {
		select {
		case <-ticker.C:
			fmt.Println("Break number ", count)
			count++

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
		Duration: hey.DefaultDuration,
	})
}
