package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {

	// primero solo descomenten executionWithoutConcurrency y corranlo
	//executionWithoutConcurrency()
	// luego solo descomenten executionWithConcurrency y sientan la concurrencia en su salsa!
	executionWithConcurrency()

	/*fmt.Println("Sin concurrencia")
	longRunningTasksWithoutConcurrency()
	fmt.Println("Con concurrencia")
	longRunningTasksWithoutConcurrency()*/
}

/*
func executionWithConcurrencyUsingChannels() {

	// TIP (más a la derecha... VER SOLO si no se te cae una idea.................................................. seguro que querés ver? tampoco es la gran cosa... .................................... bueno, necesitamos un canal para valores enteros  `var channel chan int = make(chan int)`)

	go func() {
		sum := doAnExpensiveSum()
	}()
	go func() {
		number := getNumberFromStdin()
	}()

	// lo que quiero es imprimir aca tanto `sum` e `number`
}*/

func executionWithConcurrency() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		sum := doAnExpensiveSum()
		fmt.Println(sum)
		wg.Done()
	}()
	go func() {
		input := getNumberFromStdin()
		fmt.Println(input)
		wg.Done()
	}()

	wg.Wait()
}

func executionWithoutConcurrency() {
	sum := doAnExpensiveSum()
	fmt.Println(sum)

	input := getNumberFromStdin()
	fmt.Println(input)
}

func doAnExpensiveSum() int {
	ac := 0
	for i := 1; i < 3000000000; i++ {
		ac += i
	}
	return ac
}

func getNumberFromStdin() int {
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Println("Enter valid number (or press CTR+D abort):")
		hasInput := scanner.Scan()
		if !hasInput {
			break
		}
		text := scanner.Text()
		number, err := strconv.Atoi(text)
		if err == nil {
			return number
		}
	}
	return 0
}

func sumConsecutiveInteger(from int, to int) int {
	ac := 0
	for i := from; i <= to; i++ {
		ac += i
	}
	return ac
}

func longRunningTasksWithoutConcurrency() {
	sum := sumConsecutiveInteger(0, 3000000000)
	fmt.Println(sum)

	sum2 := sumConsecutiveInteger(10000, 4000000000)
	fmt.Println(sum2)
}

func longRunningTasksWithConcurrency() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		sum1 := sumConsecutiveInteger(0, 3000000000)
		fmt.Println(sum1)
		wg.Done()
	}()
	go func() {
		sum2 := sumConsecutiveInteger(10000, 4000000000)
		fmt.Println(sum2)
		wg.Done()
	}()

	wg.Wait()
}
