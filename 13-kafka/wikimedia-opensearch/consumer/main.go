package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/IBM/sarama"
	"github.com/google/uuid"

	opensearch "github.com/opensearch-project/opensearch-go"
)

func ConnectConsumer(brokers []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.AutoCommit.Enable=true
	config.Consumer.Offsets.AutoCommit.Interval = 5 * time.Second
	return sarama.NewConsumer(brokers, config)
}

func main() {
	indexName := "wikimedia"
	topicName := "wikimedia.recentchange"
	msgCnt := 0

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

	// Check if the index exist
	res, err := client.Indices.Exists([]string{indexName}, client.Indices.Exists.WithContext(context.Background()))
	if err != nil {
		log.Fatalf("Error checking if index exists: %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		fmt.Printf("Index %s exists.\n", indexName)
	} else if res.StatusCode == 404 {
		createIndexResponse, err := client.Indices.Create(indexName)
		if err != nil {
			log.Fatalf("Error creating index: %s", err)
		}
		defer createIndexResponse.Body.Close()
		fmt.Printf("Index %s created successfully: %v\n", indexName, createIndexResponse)
	}

	worker, err := ConnectConsumer([]string{"localhost:9092"})
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := worker.Close(); err != nil {
			fmt.Println("Error closing worker:", err)
		}
	}()

	consumer, err := worker.ConsumePartition(topicName, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			fmt.Println("Error closing consumer:", err)
		}
	}()

	fmt.Println("Consumer started")

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println("Consumer error:", err)
			case msg := <-consumer.Messages():
				msgCnt++
				newUUID := uuid.New().String()
				res, err := client.Index(
					indexName,                            // Index name
					bytes.NewReader(msg.Value),           // Document body
					client.Index.WithDocumentID(newUUID), // Document ID
					client.Index.WithRefresh("true"),     // Refresh the index immediately
				)
				if err != nil {
					log.Fatalf("Error indexing document: %s", err)
				}
				defer res.Body.Close()

				// Print the response status
				if res.StatusCode == 201 {
					fmt.Println("Document indexed successfully! ", newUUID)
				} else {
					fmt.Printf("Failed to index document. Status: %s\n", res.Status())
				}

			case <-sigchan:
				fmt.Println("Interrupt is detected")
				close(doneCh)
				return
			}
		}
	}()

	// Wait for signal or completion
	<-doneCh
	fmt.Println("Processed", msgCnt, "messages")

}
