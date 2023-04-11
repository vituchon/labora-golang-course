package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	x, err := strconv.Atoi(text[:len(text)-1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(SegmentarValorPorRangos(x))
}
