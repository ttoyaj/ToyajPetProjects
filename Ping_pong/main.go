package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Serve static files (index.html will be the root page)

	// Serve static files from /static/
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Serve index.html at root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// API endpoint
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
