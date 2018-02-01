package main

import (
	"net/http"
	"fmt"
	"strconv"
	"time"
	"os"
	"sync"
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

	for i := startPort; i <= endPort; i++ {
		go func(ipPort int) {
			mux := http.NewServeMux()
			mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				//fmt.Printf("URL.Path  = %q\n", r.URL.Path)
				//fmt.Fprintf(w, "URL.Path  = %q\n", r.URL.Path)
				called <- "ipPort "+ strconv.Itoa(ipPort) + " http call on " + r.URL.Path
			})
			http.ListenAndServe("localhost:"+strconv.Itoa(ipPort), mux)
		}(i)
	}

	endTime := time.Now();

	fmt.Printf(("%s Web Servers started for each port between %d and %d\n"), endTime.String(),  startPort, endPort)
	fmt.Printf("Web Server startup time (nanos) = %d\n", (endTime.UnixNano() - startTime.UnixNano()))

	// run client thread randomized

	var wg sync.WaitGroup

	go func() {
		for i := startPort; i <= endPort; i++ {
			ipPort := i
			wg.Add(1)
			go func() {
				//fmt.Printf("http.Get on %s\n", "http://localhost:" + strconv.Itoa(ipPort) + "/" + strconv.Itoa(ipPort))
				response, err := http.Get("http://localhost:" + strconv.Itoa(ipPort) + "/" + strconv.Itoa(ipPort))
				if err != nil {
					fmt.Printf("%s\n", err)
				} else {
					defer response.Body.Close()
					fmt.Print(".")
				}
				wg.Done()
			}()
		}
	}()

	go func() {
		fmt.Printf("%s Waiting for workers to finish\n", time.Now().String())
		wg.Wait()
		fmt.Printf("%s We terminate now in 15s\n", time.Now().String())
		time.Sleep(15 * time.Second)
		os.Exit(0)
	}()

	for true {
		select {
		case <-called:
			fmt.Print("+")
		}
	}
	fmt.Println("Done")

}

