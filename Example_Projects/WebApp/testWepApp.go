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
    fmt.Fprintf(w, "<h1>%s %s!</h1>","Hi there, I love", r.URL.Path[1:])
}

func loginHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>You are in login Page</h1>");
}

func main() {
    http.HandleFunc("/", handler)
     http.HandleFunc("/login", loginHandler)
    http.ListenAndServe(":7272", nil)
}