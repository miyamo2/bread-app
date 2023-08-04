package entity

import (
	"time"
)

// Bread Firestore Entity
type Bread struct {
	ID        string `firestore:"id,omitempty"`
	Name      string `firestore:"name,omitempty"`
	CreatedAt time.Time `firestore:"createdAt,omitempty"`
}

func NewBread(id string, name string, createdAt time.Time) Bread {
	return Bread{
		ID:        id,
		Name:      name,
		CreatedAt: createdAt,
	}
}
