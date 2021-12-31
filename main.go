package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	day := time.Now().Weekday()

	switch day {
	case time.Monday:
		fmt.Println("Happy Monday!")
	case time.Tuesday:
		fmt.Println("Have a good Tuesday!")
	case time.Wednesday:
		fmt.Println("Hope the week is going well! It's Wednesday")
	case time.Thursday:
		fmt.Println("Have a good Thursday!")
	case time.Friday:
		fmt.Println("Congrats it's Friday!!")
	case time.Saturday:
		fmt.Println("Happy Saturday!")
	case time.Sunday:
		fmt.Println("Happy Sunday!")
	default:
		fmt.Println("Bad Day")
	}

	tasks := getDaysTasks(day)

	for _, task := range tasks {
		fmt.Printf(" > %s\n", task)
	}
}

func getDaysTasks(day time.Weekday) []string {
	data, err := os.ReadFile("tasks/" + day.String() + ".txt")
	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), "\n")
}
