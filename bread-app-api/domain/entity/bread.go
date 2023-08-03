package entity

// Bread Firestore Entity
type Bread struct {
	ID        string `firestore:"id,omitempty"`
	Name      string `firestore:"name,omitempty"`
	CreatedAt string `firestore:"createdAt,omitempty"`
}

func NewBread(id string, name string, createdAt string) Bread {
	return Bread{
		ID:        id,
		Name:      name,
		CreatedAt: createdAt,
	}
}
