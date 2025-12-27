package main

import (
	"fmt"
	"time"
)

func getGreeting() string {
	hour := time.Now().Hour()

	var greeting string

	switch {
	case hour >= 0 && hour < 6:
		greeting = "Good night!"
	case hour >= 6 && hour < 12:
		greeting = "Good morning!"
	case hour >= 12 && hour < 18:
		greeting = "Good afternoon!"
	case hour >= 18 && hour < 23:
		greeting = "Good evening!"
	default:
		greeting = "Hello!"
	}

	return greeting
}

func main() {

	projectName := "Better HTTP"

	greeting := getGreeting()

	log := fmt.Sprintf("%s Running %s...", greeting, projectName)

	fmt.Println(log)
}
