package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/vituchon/labora-golang-course/segment-value-by-ranges/segmentor"
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
	fmt.Println(segmentor.SegmentarValorPorRangos(x))
}
