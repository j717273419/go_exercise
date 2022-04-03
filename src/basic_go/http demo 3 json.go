package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello World!"))
}

// Output: set header
// received get request & params
// received post request & params
func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("X-My-Custom-Header", "some value")
	if r.Method == "GET" {
		// w.Write([]byte("Users! get " + r.URL.Query()["q"][0]))
		params := make(map[string]string)
		for k, v := range r.URL.Query() {
			params[k] = v[0]
		}
		fmt.Println("received get request & params: %v", params)
		w.Write([]byte("Users! get q:" + params["q"] + " p:" + params["p"]))
	} else {
		reqBody, err := ioutil.ReadAll(r.Body)
		 if err != nil {
		 	log.Fatal(err)
		 }		 
		 fmt.Printf("%s\n", reqBody)		
		w.Write([]byte("Users! post " + string(reqBody)))
	}
}

// Output: json
func projects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var users = `{"hello": "world"}`
	w.Write([]byte(users))
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/users", users)
	http.HandleFunc("/projects", projects)
	http.ListenAndServe(":8000", nil)
}
