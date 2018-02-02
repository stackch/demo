package main

import (
	"sync"
	"net/http"
	"fmt"
	"strconv"
	"time"
	"os"
)

func main() {
	var startPort int = 8000
	var endPort int = 8001

	if len(os.Args) >= 3 {
		startPort, _ = strconv.Atoi(os.Args[1])
		endPort, _ = strconv.Atoi(os.Args[2])
	}
	called := make(chan string)

	startTime := time.Now();

	fmt.Printf(("%s Start Web Servers for each port between %d and %d\n"), startTime.String(), startPort, endPort)

	var wg sync.WaitGroup

	for i := startPort; i <= endPort; i++ {
		wg.Add(1)
		go func(ipPort int) {
			mux := http.NewServeMux()
			mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				//fmt.Printf("URL.Path  = %q\n", r.URL.Path)
				//fmt.Fprintf(w, "URL.Path  = %q\n", r.URL.Path)
				called <- "ipPort "+ strconv.Itoa(ipPort) + " http call on " + r.URL.Path
			})
			wg.Done()
			http.ListenAndServe("localhost:"+strconv.Itoa(ipPort), mux)
		}(i)
	}

	wg.Wait()

	endTime := time.Now();

	fmt.Printf(("%s Web Servers started for each port between %d and %d\n"), endTime.String(),  startPort, endPort)
	fmt.Printf("Web Server startup time (nanos) = %d\n", (endTime.UnixNano() - startTime.UnixNano()))

	for true {
		select {
		case <-called:
			fmt.Print("+")
		}
	}

	fmt.Println("Done")

}

