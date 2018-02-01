// package main
package main

import (
	"fmt"
	"strconv"
	"time"
	"os"
	"sync"
    "sync/atomic"
)

var atomicInt int64 = 0

func main() {

	var goroutines int = 10

	if len(os.Args) >= 2 {
		goroutines, _ = strconv.Atoi(os.Args[1])
	}

	startTime := time.Now();

	fmt.Printf(("%s Start %d goroutines\n"), startTime.String(), goroutines)

	var wg sync.WaitGroup

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&atomicInt, 1)
			wg.Done()
		}()
	}

	endTime := time.Now();

	fmt.Printf(("%s %d goroutines started up\n"), endTime.String(),  goroutines)

	wg.Wait()

	fmt.Printf("Done, atomicInt = %d\n", atomicInt)

}

