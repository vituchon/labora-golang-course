package main

import (
	"fmt"
	//"github.com/vituchon/labora-golang-course/test-things/util"
)

var channel chan bool = make(chan bool)

func main() { // ver en playground
	go meEjecutaranEnOtraGorutina()

	fmt.Println("Soy la gorutina principal")
	<-channel // recibo un valor del canal
	fmt.Println("FIN main")
}

func meEjecutaranEnOtraGorutina() {
	fmt.Println("Soy otra gorutina")
	channel <- true // envío un valor al canal
	fmt.Println("FIN meEjecutaranEnOtraGorutina")
}
