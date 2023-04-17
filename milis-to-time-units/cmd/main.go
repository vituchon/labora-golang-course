package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	resolution_stack_overflow "github.com/vituchon/labora-golang-course/milis-to-time-units/resolution-1"
	resolution_home_made "github.com/vituchon/labora-golang-course/milis-to-time-units/resolution-2"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputStr := scanner.Text()
		durationInMilis, err := strconv.Atoi(inputStr)
		if err == nil {
			fmt.Printf("Result:%+v", resolution_stack_overflow.SplitSecondsInDaysHoursMinutesAndSeconds(durationInMilis))
			fmt.Printf("Result:%+v", resolution_home_made.SplitSecondsInDaysHoursMinutesAndSeconds(durationInMilis))
		} else {
			fmt.Printf("IOooooo: '%v'", err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
