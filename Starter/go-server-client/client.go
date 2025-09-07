package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	msg := "Hello from client!"
	resp, err := http.Post("http://localhost:8080/message", "text/plain", bytes.NewBuffer([]byte(msg)))
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Server response:", string(body))
}
