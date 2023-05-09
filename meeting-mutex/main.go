package main

import (
	"fmt"
	"sync"
	"time"
)

var mapita map[int]int = make(map[int]int)
var mutex sync.Mutex

func getOrCreate(key int, value int) int {
	//fmt.Printf("Invocación a getOrCreate(%d,%d)\n", key, value)
	//mutex.Lock()
	_, exist := mapita[key]
	fmt.Printf("getOrCreate(%d,%d), verificación de clave: %t\n", key, value, exist)
	if !exist {
		mapita[key] = value
		fmt.Printf("getOrCreate(%d,%d), modificación de clave\n", key, value)
	}
	//mutex.Unlock()
	return mapita[key]
}

func main() {

	go getOrCreate(1, 1)
	go getOrCreate(1, 2)
	go getOrCreate(1, 3)
	go getOrCreate(1, 4)

	time.Sleep(time.Second * 2)
	fmt.Printf("mapita:%+v\n", mapita)
}
