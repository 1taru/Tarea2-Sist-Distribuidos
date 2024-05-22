package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/1taru/Tarea2-Sist-Distribuidos/internal/consumer"
	"github.com/1taru/Tarea2-Sist-Distribuidos/internal/database"
	"github.com/1taru/Tarea2-Sist-Distribuidos/pkg/brokers"
	"github.com/1taru/Tarea2-Sist-Distribuidos/pkg/miembro"
	"github.com/IBM/sarama"
)

var (
	sig chan bool
	db  *sql.DB
)

func processMembers() {

	cg, err := consumer.NewConsumerGroup(brokers.Brokers, "membresias", sarama.OffsetOldest)

	if err != nil {
		log.Panic(err)
	}

	defer cg.Close()

	ch := consumer.ConsumerHandler{
		Ready: make(chan bool),
		F: func(msg *sarama.ConsumerMessage) {
			var m miembro.Miembro

			err := json.Unmarshal(msg.Value, &m)

			if err != nil {
				log.Panic(err)
			}

			stmt, err := db.Prepare("INSERT INTO Miembros (nombre, apellido, rut, email, patente, premium) VALUES ($1, $2, $3, $4, $5, $6);")

			if err != nil {
				log.Panic(err)
			}

			defer stmt.Close()

			_, err = stmt.Query(
				m.Nombre,
				m.Apellido,
				m.Rut,
				m.Email,
				m.Patente,
				m.Premium,
			)

			if err != nil {
				return
			}
			fmt.Println("El miembro ha sido añadido 🙌")

		},
	}

	ctx, cancel := context.WithCancel(context.Background())

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {

			if err := cg.Consume(ctx, []string{"Membresias"}, &ch); err != nil {
				log.Panic(err)
			}

			if ctx.Err() != nil {
				return
			}

			ch.Ready = make(chan bool)

		}

	}()

	<-ch.Ready

	for {

		<-ctx.Done()
		break

	}

	cancel()
	wg.Wait()

	sig <- true

}

func main() {
	sig = make(chan bool)

	db = database.New()
	defer db.Close()

	go processMembers()

	<-sig
}
