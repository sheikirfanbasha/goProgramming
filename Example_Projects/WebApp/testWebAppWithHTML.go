/* 
net/http : allows you to write web applications
html/template : allows us to serve html pages as response

The notation {{variable/expression}} is used to embed dynamic data into html

Step1:
	go run testWebAppWithHTML.go
Step2:
	Open the browser and run
		localhost:7272/edit/edit

Step3: 
	Result would be
	edit.html content served as form
*/
package main

import (
    "html/template"
    "net/http"
    "io/ioutil"
    "fmt"
)
type Page struct {
	Title string
	Body  []byte
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    fmt.Println(title)
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    t, _ := template.ParseFiles("edit.html")
    t.Execute(w, p)
 //    body, err := ioutil.ReadFile("edit.txt")
 //    if err != nil {
	// 	panic(err)
	// }
	//fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title,p.Body)
}

func main() {
    http.HandleFunc("/edit/", editHandler)
    http.ListenAndServe(":7272", nil)
}