package main

import (
	"fmt"
	"sort"
	"sync"
)

var j int = 0

func main() {
	var wg sync.WaitGroup
	// extraidos de : https://sharbeargle.gitbooks.io/golang-notes/content/nested-functions.html
	var i int = 0
	wg.Add(2)
	for i < 2 {
		go func(iCopy int) {
			var local int = 3
			fmt.Println(iCopy, i, j, local)
			wg.Done()
		}(i) //
		i++
	}
	fmt.Println("termine de iterar, i vale ", i)

	var numbers []int = []int{1, 2, 3}
	sort.Slice(numbers, func(i int, j int) bool {
		var misOtrosNumbers []int = []int{1, 2, 3}
		misOtrosNumbers = misOtrosNumbers
		return numbers[i] < numbers[j]
	})

	wg.Wait()
	//time.Sleep(2 * time.Second)
}

func main2() {

	var funcs []func(i int)
	var i int = 0
	for i < 2 {
		anonymousFunc := func(i int) {
			fmt.Println(i)
		}
		funcs = append(funcs, anonymousFunc)
		i++
	}

	for _, f := range funcs {
		f(i)
	}
}
