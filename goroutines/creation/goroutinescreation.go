package main

import (
	"fmt"
	"strconv"
	"time"
	"os"
	"sync"
)

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
			s := ""
			for i := 0; i < 20000; i++ {
				s += strconv.Itoa(i)
				//time.Sleep(time.Millisecond)
			}
			fmt.Print(".")
			wg.Done()
		}()
	}

	endTime := time.Now();

	fmt.Printf(("%s %d goroutines started up\n"), endTime.String(),  goroutines)

	wg.Wait()

	fmt.Println("Done")

}

