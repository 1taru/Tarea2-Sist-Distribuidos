// productor.go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	// Configuración del productor
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	// Crear un nuevo productor
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Error al crear el productor: %v", err)
	}
	defer producer.Close()

	// Enviar mensajes en un bucle
	for {
		message := &sarama.ProducerMessage{
			Topic: "test",
			Key:   sarama.StringEncoder("key"),
			Value: sarama.StringEncoder(fmt.Sprintf("Mensaje enviado a las %s", time.Now().String())),
		}

		partition, offset, err := producer.SendMessage(message)
		if err != nil {
			log.Printf("Error al enviar el mensaje: %v", err)
		} else {
			log.Printf("Mensaje enviado a la partición %d con el offset %d", partition, offset)
		}

		time.Sleep(2 * time.Second) // Espera 2 segundos antes de enviar el siguiente mensaje
	}
}

/*
package producer

import (
	"github.com/Shopify/sarama"
)

type Producer struct {
	BrokerList []string
	Producer   sarama.SyncProducer
}

func (p *Producer) Close() error {
	return p.Producer.Close()
}

func NewProducer(brokerList []string) (prod *Producer, err error) {
	prod = &Producer{
		BrokerList: brokerList,
	}

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewManualPartitioner

	prod.Producer, err = sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		return nil, err
	}

	return
}

func (p *Producer) SendMessage(topic string, part int32, message []byte) (partition int32, offset int64, err error) {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: part,
		Value:     sarama.StringEncoder(message),
	}

	partition, offset, err = p.Producer.SendMessage(msg)
	return
}
*/
