// http client demo
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	// http send get request
	res, err := http.Get("http://ifconfig.io/ip")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
	fmt.Println("********************************************************")
	// http send post request
	var json = strings.NewReader(`{"name":"John", "age":30, "foo":"bar"}`)
	res2, err2 := http.Post("http://httpbin.org/post", "application/json", json)
	if err2 != nil {
		log.Fatal(err2)
	}
	robots2, err2 := ioutil.ReadAll(res2.Body)
	res2.Body.Close()
	fmt.Printf("%s", robots2)
}