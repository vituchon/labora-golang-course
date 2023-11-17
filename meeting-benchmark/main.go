package main

import (
	"fmt"
	"time"
)

func main () {
	printWhenIsSaturdaySwitch()
	printWhenIsSaturdayIf()
}

func printWhenIsSaturdaySwitch() {
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		//fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func printWhenIsSaturdayIf() {
	today := time.Now().Weekday()
	if (time.Saturday == today +0) {
			fmt.Println("Today.")
		} else if (time.Saturday == today +1) {
			//fmt.Println("Tomorrow.")
		} else if (time.Saturday == today + 2) {
			fmt.Println("In two days.")
		} else {
			fmt.Println("Too far away.")
		}
}


