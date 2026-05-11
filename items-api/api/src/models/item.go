package models

import (
	"errors"
	"strings"
	"time"
)

type Item struct {
	ID        uint64    `json:"id",omitempty"`
	Product   string    `json:"Product,omitempty"`
	Name      string    `json:"name,omitempty"`
	Category string    `json:"Category,omitempty"`
	Token     string    `json:"token,omitempty"`
	CreatedAt  time.Time `json:"CreatedAt,omitempty"`
}

// Prepare vai call methods to validate and format the user received
func (item *Item) Prepare(etapa string) error {
	if err := item.validar(etapa); err != nil {
		return err
	}

	item.formatar()
	return nil
}

func (item *Item) validar(etapa string) error {
	if item.Product == "" {
		return errors.New("O Product is required and cannot be blank")
	}
	if item.Name == "" {
		return errors.New("O Name is required and cannot be blank")
	}
	if item.Category == "" {
		return errors.New("A Category é obrigatória e não pode estar em branco")
	}
	if etapa == "registration" && item.Token == "" {
		return errors.New("O token is required and cannot be blank")
	}
	return nil
}
func (item *Item) formatar() {
	item.Product = strings.TrimSpace(item.Product)
	item.Name = strings.TrimSpace(item.Name)
	item.Category = strings.TrimSpace(item.Category)
}
