package main

import (
	"fmt"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header { //assign headers from a map to its name
		for _, h := range headers { //ignore the name and ensure all headers, h, are accessed
			fmt.Fprintf(w, "%v: %v\n", name, h) //Write to responsewriter, w, "Headername: Header\n"
		}
	}
}

func end(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "Closing down server.\n")

	os.Exit(0) //Probably bad practice TODO Figure out good practice for remote shutdown of httpServers
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/exit", end)

	http.ListenAndServe(":8090", nil) //localhost:8090, nil tells to use default router we set up.
}
