package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/vituchon/labora-golang-course/milis-to-time-units/resolution/home"
	"github.com/vituchon/labora-golang-course/milis-to-time-units/resolution/stackoverflow"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Print("Enter seconds (A valid int) or CTR+C to quit:")
		hasMoreInput := scanner.Scan()
		if !hasMoreInput {
			break
		}
		inputStr := scanner.Text()
		durationInMilis, err := strconv.Atoi(inputStr)
		if err == nil {
			fmt.Printf("Result (using stackoverflow implementation):%+v\n", stackoverflow.SplitSecondsInDaysHoursMinutesAndSeconds(durationInMilis))
			fmt.Printf("Result (using home made implementation):%+v\n", home.SplitSecondsInDaysHoursMinutesAndSeconds(durationInMilis))
		} else {
			fmt.Printf("IOooooo: '%v'", err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
