/*a simple web server*/
package main

import (
	"path/filepath"
	"os"
	"net/http"
	"log"
)

func main() {
	port := "8080"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	log.Println("std.ch Web Server is starting up")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:" + port, nil))
	log.Println("std.ch  Web Server terminates")
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	workingPath, _ := os.Getwd()
	realPath := filepath.Join(workingPath, path)
	stat, err := os.Stat(realPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("std.ch file " + realPath + " does not exist")
			return
		}
	}
	if stat.IsDir() {
		// check for index.html file
		realPath = filepath.Join(realPath, "index.html")
		if stat, err = os.Stat(realPath); err != nil {
			if os.IsNotExist(err) {
				log.Println("std.ch file " + realPath + " not found")
			}
		}					
	}
	log.Println("std.ch serve file " + realPath)
    http.ServeFile(w, r, realPath)	
}
