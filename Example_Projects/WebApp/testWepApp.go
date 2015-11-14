/* 
net/http : allows you to write web applications
Step1:
	go run testWebApp.go
Step2:
	Open the browser and run
		localhost:7272/Programming

Step3: 
	Result would be
	'Hi there, I love Programming'
*/
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":7272", nil)
}