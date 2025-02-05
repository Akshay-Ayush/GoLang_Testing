package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is for testing")
}

func main() {
	// http://localhost:8080/hello	
	http.HandleFunc("/hello", helloHandler)
	// http://localhost:8080/testing
	http.HandleFunc("/testing", testHandler)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)  // Start server on port 8080
}
