package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	current_month := currentTime.Month()
	current_date := currentTime.Day()

	// Catppuccin Colors
	// var Red = "\x1b[38;2;243;139;168m"
	var text = "\x1b[38;2;205;214;244m"

	// Main Logic
	switch current_month.String() {
	case "June":
		fmt.Println(text + saints.june[current_date])
	case "July":
		fmt.Println(text + july[current_date])
	}

}
