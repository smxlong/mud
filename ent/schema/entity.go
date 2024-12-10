package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Entity holds the schema definition for the Entity entity.
type Entity struct {
	ent.Schema
}

// newID returns a new unique identifier.
func newID() string {
	return uuid.New().String()
}

// Fields of the Entity.
func (Entity) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(newID).
			Unique().
			Immutable(),
		field.String("name").
			NotEmpty(),
		field.String("description").
			Optional(),
	}
}
