package main

import (
	"fmt"
	"log"
	"net/http"

	opensearch "github.com/opensearch-project/opensearch-go"
)

func main() {
	// Create a new OpenSearch client
	client, err := opensearch.NewClient(opensearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		// Username: "your-username", 
		// Password: "your-password", 
		Transport: &http.Transport{
			// Skip TLS verification if needed (not recommended for production)
			// TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	})
	if err != nil {
		log.Fatalf("Error creating OpenSearch client: %s", err)
	}

	// Ping the OpenSearch cluster to check the connection
	res, err := client.Info()
	if err != nil {
		log.Fatalf("Error getting response from OpenSearch: %s", err)
	}
	defer res.Body.Close()

	fmt.Println("OpenSearch Info:")
	fmt.Println(res)

	// Example: Create an index
	/*
	indexName := "test-index-2"
	createIndexResponse, err := client.Indices.Create(indexName)
	if err != nil {
		log.Fatalf("Error creating index: %s", err)
	}
	defer createIndexResponse.Body.Close()

	fmt.Printf("Index %s created successfully: %v\n", indexName, createIndexResponse)
	*/
}
