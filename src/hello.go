package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Printf("hello go")
	http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		s := fmt.Sprintf("你好，世界！-- Time:%s", time.Now().String())
		fmt.Fprintf(w, "%v\n", s)
		log.Printf("%v\n", s)
	})
	if err := http.ListenAndServe(":12345", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
