package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"kafka-processing-system/model"
	"log"
	"net/http"
)

func main() {
	url := "http://localhost:8080/send"

	for i := 1; i <= 10; i++ {
		solicitud := model.Solicitud{
			ID:     fmt.Sprintf("solicitud-%d", i),
			Nombre: fmt.Sprintf("Producto %d", i),
			Precio: 759.000,
			Correo: "distritest681@gmail.com",
		}

		jsonData, err := json.Marshal(solicitud)
		if err != nil {
			log.Printf("Failed to marshal message: %s", err)
			continue
		}

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Printf("Failed to send request: %s", err)
			continue
		}

		if resp.StatusCode == http.StatusOK {
			log.Printf("Solicitud %d enviada exitosamente", i)
		} else {
			log.Printf("Failed to send request: received status code %d", resp.StatusCode)
		}
	}
}
