package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	url := "https://stream.wikimedia.org/v2/stream/recentchange"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error fetching stream:", err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "data: ") {
			data := strings.TrimPrefix(line, "data: ")
			fmt.Println("Received data:", data)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading from stream:", err)
	}
}
