package brokers

import (
	"log"
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

var Brokers []string

func init() {
	getBrokerList()
}

func getBrokerList() {
	broker := os.Getenv("BROKER_NET")

	if broker == "" {
		broker = "localhost:9092"
	}

	Brokers = strings.Split(broker, ",")

	if len(Brokers) == 0 {
		log.Panic("Brokers not found")
	}
}

/*package brokers

import (
	"log"

	"github.com/IBM/sarama"
)

func main() {
	// Configuración del productor
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	// Lista de brokers de Kafka
	brokers := []string{"localhost:9092"}

	// Crear el productor de Kafka
	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Error al crear el productor de Kafka: %v", err)
	}
	defer producer.Close()

	// Crear un mensaje
	message := &sarama.ProducerMessage{
		Topic: "mi-topico",
		Value: sarama.StringEncoder("Hola, Kafka!"),
	}

	// Enviar el mensaje al productor
	producer.Input() <- message

	// Escuchar por errores y eventos del productor
	go func() {
		for err := range producer.Errors() {
			log.Printf("Error al enviar mensaje a Kafka: %v", err)
		}
	}()

	// Manejar eventos de éxito
	go func() {
		for success := range producer.Successes() {
			log.Printf("Mensaje enviado a la partición %d con el offset %d", success.Partition, success.Offset)
		}
	}()

	// Esperar hasta que el productor haya enviado todos los mensajes pendientes
	producer.AsyncClose()
}
*/
