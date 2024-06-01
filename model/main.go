package model

import (
	"time"
)

// Solicitud defines the structure of a request
type Solicitud struct {
	ID     string    `json:"id"`
	Nombre string    `json:"nombre"`
	Precio float64   `json:"precio"`
	Correo string    `json:"correo"`
	Estado string    `json:"estado"`
	Time   time.Time `json:"time"`
}
