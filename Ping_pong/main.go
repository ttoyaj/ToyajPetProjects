package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// letReplyToPing decides response based on current IST time
func letReplyToPing() string {
	// Load IST timezone
	ist, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		return "Error loading timezone"
	}

	now := time.Now().In(ist)
	hour := now.Hour()

	if hour < 21 {
		return "PONG from Toyaj"
	} else {
		return "PONG from Adithi"
	}
}

func main() {
	// Serve static files from /static/
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Serve static/index.html at root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	// API endpoint with time-based response
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		reply := letReplyToPing()
		fmt.Fprintln(w, reply)
	})

	fmt.Println("Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
		os.Exit(1)
	}
}
