package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func returnr(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "%s\n", string(requestDump))
}

func main() {
	port := "8080"
	http.HandleFunc("/returnr", returnr)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
