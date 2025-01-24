package main

import (
	"bufio"
	"log"
	"net/http"
	"strings"

	"github.com/IBM/sarama"
)

var producer sarama.SyncProducer

func ConnectProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	// Safe Producer Options
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Idempotent=true
	config.Net.MaxOpenRequests=1

	// High Throughput Options
	config.Producer.Compression= sarama.CompressionSnappy
	config.Producer.Flush.Frequency = 20  // 20 ms
	config.Producer.Flush.Bytes = 32 * 1024   // 32 KB
	return sarama.NewSyncProducer(brokers, config)
}

func main() {
	var err error
	brokers := []string{"localhost:9092"}
	producer, err = ConnectProducer(brokers)
	if err != nil {
		log.Fatalf("Failed to connect producer: %v", err)
	}
	defer producer.Close()


	url := "https://stream.wikimedia.org/v2/stream/recentchange"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error fetching stream:", err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	topicName:="wikimedia.recentchange"

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "data: ") {
			data := strings.TrimPrefix(line, "data: ")
			msg := &sarama.ProducerMessage{
				Topic: topicName,
				Value: sarama.StringEncoder(data),
			}

			_, _, err := producer.SendMessage(msg)
			if err != nil {
				log.Fatal("Error sending message to a topic:", err)
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading from stream:", err)
	}
}
