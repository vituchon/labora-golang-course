// main.go
package main

import (
	"fmt"
)

func main() {
	var val int
	val = sumPairsMethod1(11)
	fmt.Println(val)
	val = sumPairsMethod2(11)
	fmt.Println(val)
	val = sumPairsMethod3(11)
	fmt.Println(val)
}

// uses mod
func sumPairsMethod1(n int) int {
	sum := 0
	loops := n * 2
	for i := 1; i <= loops; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	return sum
}

// uses multiply
func sumPairsMethod2(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		pairNum := i * 2
		sum += pairNum
	}
	return sum
}

// uses add
func sumPairsMethod3(n int) int {
	sum := 0
	pairNum := 2
	for i := 1; i <= n; i++ {
		sum += pairNum
		pairNum += 2
	}
	return sum
}
