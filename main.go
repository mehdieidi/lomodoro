package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/gen2brain/beeep"
)

func main() {
	workDuration := flag.Int("w", 25, "how much minutes do you work?")
	breakDuration := flag.Int("b", 2, "how much minutes do you rest?")
	flag.Parse()

	fmt.Printf("work: %d mins, break: %d mins\n", *workDuration, *breakDuration)

	ticker := time.NewTicker(time.Duration(*workDuration) * time.Minute)

	count := 1

	for {
		<-ticker.C
		fmt.Println("Break number ", count)
		count++

		topic := "break"
		text := "enough man! go get some rest."
		notify(topic, text)

		time.Sleep(time.Duration(*breakDuration) * time.Minute)
		topic = "work"
		text = "let's make money and get rich!"
		notify(topic, text)

		ticker = time.NewTicker(time.Duration(*workDuration) * time.Minute)
	}
}

func notify(topic, text string) {
	err := beeep.Notify(topic, text, "")
	if err != nil {
		panic(err)
	}
}
