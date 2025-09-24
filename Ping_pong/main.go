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

	if hour < 19 {
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

	http.HandleFunc("/getTime", func(w http.ResponseWriter, r *http.Request) {
		ist, err := time.LoadLocation("Asia/Kolkata")
		if err != nil {
			http.Error(w, "Error loading timezone", http.StatusInternalServerError)
			return
		}
		currentTime := time.Now().In(ist).Format("2006-01-02 15:04:05 MST")
		fmt.Fprintf(w, "Current time (IST): %s", currentTime)
	})

	fmt.Println("Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
		os.Exit(1)
	}
}
