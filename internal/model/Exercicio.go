package model

import "time"

type Exercicio struct {
	ID         string    `json:"id"`
	Nome       string    `json:"nome"`
	Series     int       `json:"series"`
	Repeticoes int       `json:"repeticoes"`
	Descanso   int       `json:"descanso"` // em segundos
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

