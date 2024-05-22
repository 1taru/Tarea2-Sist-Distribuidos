package main

//https://go-mail.dev/getting-started/introduction/
import (
	"log"

	"github.com/wneessen/go-mail"
)

func main() {
	m := mail.NewMsg()
	if err := m.From("toni.sender@example.com"); err != nil {
		log.Fatalf("failed to set From address: %s", err)
	}
	if err := m.To("tina.recipient@example.com"); err != nil {
		log.Fatalf("failed to set To address: %s", err)
	}
}

/*
func main() {
	// Definir la URL de Kafka y el tópico al que consultar
	kafkaURL := "http://localhost:9092"
	topic := "mi_topico"

	// Número de consultas a realizar
	numQueries := 100

	// Inicializar variables para el promedio de tiempo de respuesta
	totalDuration := time.Duration(0)
	count := 0

	// Loop infinito para realizar consultas
	for {
		// Tiempo de inicio de la iteración
		start := time.Now()

		// Realizar la consulta a Kafka utilizando net/http
		_, err := http.Post(kafkaURL+"/topics/"+topic, "application/vnd.kafka.json.v2+json", bytes.NewBuffer([]byte(`{"records":[{"value":"Mensaje de prueba"}]}`)))
		if err != nil {
			log.Fatalf("Error al realizar la consulta: %v", err)
		}

		// Tiempo de finalización de la iteración
		end := time.Now()

		// Calcular el tiempo de respuesta
		duration := end.Sub(start)

		// Sumar el tiempo de respuesta al total
		totalDuration += duration

		// Incrementar el contador de consultas
		count++

		// Imprimir el promedio de tiempo de respuesta
		averageDuration := totalDuration / time.Duration(count)
		fmt.Printf("Promedio de tiempo de respuesta: %s\n", averageDuration)

		// Esperar un segundo antes de la siguiente consulta
		time.Sleep(time.Second)
	}
}
*/
