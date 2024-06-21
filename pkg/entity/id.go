package entity

import (
	"github.com/google/uuid"
)

type ID = uuid.UUID

// retorna um id no formato uuid
func NewID() ID {
	return ID(uuid.New())
}

// verificando se id é válido
func ParseID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}
