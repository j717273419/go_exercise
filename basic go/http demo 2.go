package main
import ("net/http")

func helloWorld(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello World!"))
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("X-My-Custom-Header", "some value")
	w.Write([]byte("Users!"))
}

func projects(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Projects!"))
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/users", users)
	http.HandleFunc("/projects", projects)
	http.ListenAndServe(":8000", nil)
}
