package main

import "fmt"
import "net/http"


func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World - running image testworld:#{Build.BuildId}#</h1>")
}


func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Service Health Check : Passed</h1>")
}

func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/health",health)
	fmt.Println("Server Listening at 8080...")
	http.ListenAndServe(":8080",nil)
}