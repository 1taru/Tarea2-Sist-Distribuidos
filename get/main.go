package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Solicitud struct {
	ID     string    `json:"id"`
	Correo string    `json:"correo"`
	Estado string    `json:"estado"`
	Time   time.Time `json:"time"`
}

var (
	lastSolicitud Solicitud
	mu            sync.RWMutex
)

func main() {
	go consumeKafkaMessages()

	http.HandleFunc("/last", getLastSolicitudHandler)
	fmt.Println("Server is listening on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func consumeKafkaMessages() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "last-value-consumer",
		"auto.offset.reset": "latest",
	})
	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}
	defer c.Close()

	err = c.SubscribeTopics([]string{
		"solicitudes-recibido",
		"solicitudes-preparando",
		"solicitudes-entregando",
		"solicitudes-finalizado",
	}, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topics: %s", err)
	}

	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			continue
		}

		var solicitud Solicitud
		err = json.Unmarshal(msg.Value, &solicitud)
		if err != nil {
			log.Printf("Failed to unmarshal message: %s", err)
			continue
		}

		mu.Lock()
		lastSolicitud = solicitud
		mu.Unlock()
	}
}

func getLastSolicitudHandler(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(lastSolicitud)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
