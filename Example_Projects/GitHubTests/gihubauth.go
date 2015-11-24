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
    "fmt"
    "encoding/json"
    "net/url"
    "io/ioutil"
    "strings"

)


func baseHandler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("index.html")
    t.Execute(w, nil)
}

func oauthHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	acctoken_url := "https://github.com/login/oauth/access_token"
	resp, err := http.PostForm(acctoken_url,
	url.Values{"client_id": {"65d2dc71007894d51596"}, "client_secret": {"1e072e5ec68456ad77ae43dad6b7b3b622ceb2e0"}, "code":{code}})
	if err == nil{
		defer resp.Body.Close()
	    body, _ := ioutil.ReadAll(resp.Body)
	    //fmt.Println("response Body:", string(body))
	    msg := string(body)
	    m_array := strings.Split(msg, "=")
	    m1_array := strings.Split(m_array[1], "&")
	    fmt.Println("access_token is : ", m1_array[0])
	    resp, err = http.Get("https://api.github.com/user?access_token="+m1_array[0]);
	    body, _ = ioutil.ReadAll(resp.Body)
	   // fmt.Println("response Body:", string(body))
	    var res map[string]interface{}
	    json.Unmarshal([]byte(string(body)), &res)
	    
	    for k, v := range res {
	    	fmt.Println(k, "<- key value->",v)

	    }
	    if str, ok := res["login"].(string); ok {
		    /* act on str */
		    fmt.Fprintf(w, "<h1>Welcome %s</h1>", str)
		} else {
		    /* not string */
		}
	    
	}else{
		fmt.Println(err)
	}
}

func main() {
    http.HandleFunc("/", baseHandler)
    http.HandleFunc("/oauth", oauthHandler)
    http.ListenAndServe(":7272", nil)
    fmt.Println("Listening on 7272")
}