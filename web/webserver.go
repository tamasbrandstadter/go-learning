package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayName(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	// send data to client side
	if _, err := fmt.Fprintf(w, "Hello %s!", r.Form.Get("name")); err != nil {
		log.Println(err)
	}
}

func main() {
	// set router
	http.HandleFunc("/", sayName)

	// set listen port
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
