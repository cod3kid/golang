package main

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"log"
	"net/http"
)

type Order struct {
	CustomerName string `json:"customer_name"`
	CoffeeType   string `json:"coffee_type"`
}

var producer sarama.SyncProducer

func main() {
	// Initialize the producer
	var err error
	brokers := []string{"localhost:9092"}
	producer, err = ConnectProducer(brokers)
	if err != nil {
		log.Fatalf("Failed to connect producer: %v", err)
	}
	defer producer.Close()

	http.HandleFunc("/order", HandleOrder)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func HandleOrder(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
		return
	}

	// 1. Parse request body into order
	order := new(Order)
	if err := json.NewDecoder(req.Body).Decode(order); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 2. Convert body into bytes
	orderInBytes, err := json.Marshal(order)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 3. Send the bytes to kafka topic
	err = PushOrderToQueue("coffee_orders", orderInBytes)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 4. Respond back to the user
	response := map[string]interface{}{
		"success": true,
		"msg":     "Order for " + order.CustomerName + " placed successfully!",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println(err)
		http.Error(w, "Error placing order", http.StatusInternalServerError)
		return
	}
}

func ConnectProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	return sarama.NewSyncProducer(brokers, config)
}

func PushOrderToQueue(topic string, message []byte) error {
	// Create a message
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}

	log.Printf("Order is stored in topic(%s)/partition(%d)/offset(%d)\n",
		topic,
		partition,
		offset)

	return nil
}
