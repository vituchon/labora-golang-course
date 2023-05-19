// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	//var mutex sync.Mutex = sync.Mutex{}
	var wg sync.WaitGroup = sync.WaitGroup{}
	var saldo int32 = 0

	wg.Add(2)
	go func() {
		time.Sleep(1 * time.Nanosecond) // imagenos que se hace alguna operación larga
		atomic.AddInt32(&saldo, 100)
		wg.Done()
	}()
	go func() {
		time.Sleep(1 * time.Nanosecond) // imagenos que se hace alguna operación larga
		atomic.AddInt32(&saldo, 100)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("saldo final", saldo)
}
