package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	reversor "github.com/vituchon/labora-golang-course/reverse-string"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Print("Enter an string or CTR+C to quit:")
		hasMoreInput := scanner.Scan()
		if !hasMoreInput {
			break
		}
		inputStr := scanner.Text()
		fmt.Printf("Result (using golang implementation):%+v\n", reversor.ReverseRunes(inputStr))
		fmt.Printf("Result (using home made implementation):%+v\n", reversor.Reverse(inputStr))
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
