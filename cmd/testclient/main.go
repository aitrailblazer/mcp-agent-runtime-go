package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	url := flag.String("url", "http://localhost:8080/invoke", "The MCP server URL")
	file := flag.String("input-file", "./test/prompt.json", "Path to the input JSON file")
	flag.Parse()

	data, err := os.ReadFile(*file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read file: %v", err)
		os.Exit(1)
	}

	resp, err := http.Post(*url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Fprintf(os.Stderr, "HTTP request failed: %v", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Status: %s Response:%s", resp.Status, body)
}
