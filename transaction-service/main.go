package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/segmentio/kafka-go"
)

// As seen in the kotlin generator
type Transaction struct {
	TransactionID string  `json:"transaction_id"`
	UserID        int     `json:"user_id"`
	Amount        float64 `json:"amount"`
	Timestamp     int64   `json:"timestamp"`
}

// Database
var database *sql.DB

// Prometheus metrics
var transactionCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "transactions_processed_total",
		Help: "Total number of transactions processed",
	},
	[]string{"status"}, // did the request succeed or fail?
)

func main() {
	// Set up Prometheus
	prometheus.MustRegister(transactionCount)

	// Set up database connection
	var err error
	database, err = sql.Open("postgres", "user=admin password=password dbname=transactions sslmode=disable host=postgres port=5432")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// Start HTTP server for Prometheus metrics
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":2112", nil)

	// Concurrently run the Kafka consumer
	go consumeKafka()

	// Keep the service running, stop it from exiting while we run!
	select {}
}

func consumeKafka() {
	// Kafka variables
	topic := "transactions"
	brokerAddress := "kafka:9092"

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "go-consumer-group",
	})

	fmt.Println("Listening for messages on Kafka")

	// --- Simple loops, reads messages, saves to database ---
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Reader error with transactions: %v", err)
			continue
		}

		var transaction Transaction
		if err := json.Unmarshal(msg.Value, &transaction); err != nil {
			log.Printf("Failed to parse transaction json: %v", err)
			continue
		}

		saveToDatabase(transaction)

		// Increment prometheus counter
		transactionCount.WithLabelValues("success").Inc()
	}
}

func saveToDatabase(transaction Transaction) {
	// Explicitly mention the public schema
	query := `INSERT INTO public.transactions (transaction_id, user_id, amount, timestamp)
	          VALUES ($1, $2, $3, $4)`
	_, err := database.Exec(query, transaction.TransactionID, transaction.UserID, transaction.Amount, transaction.Timestamp)
	if err != nil {
		log.Printf("Failed to save transaction to database: %v", err)
		transactionCount.WithLabelValues("fail").Inc()
		return
	}

	log.Printf("Saved transaction to database: %+v", transaction)
}
