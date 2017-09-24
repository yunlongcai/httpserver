package main

import (
    "fmt"
    "net/http"
)

func sayHello(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("hello"))
}

func main() {
    http.HandleFunc("/hello", sayHello)
    err := http.ListenAndServe("0.0.0.0:8001", nil);
    if err != nil {
	fmt.Printf("error %s\n", err.Error())
	return
    }
    select{}
}
