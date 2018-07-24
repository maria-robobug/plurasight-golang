package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// Should not execute before 7am
// 8am - Good Morning
// 2pm - Good Afternoon
// 7pm - Good Evening

func main() {
	hourOfDay := time.Now().Hour()
	greeting, err := getGreeting(hourOfDay)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(greeting)
}

func getGreeting(hour int) (string, error) {
	var message string

	if hour < 7 {
		err := errors.New("too early for greetings")
		return message, err
	} else if hour < 12 {
		message = "Good morning"
	} else if hour < 18 {
		message = "Good afternoon"
	} else {
		message = "Good Evening"
	}

	return message, nil
}
