package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Door holds the schema definition for the Door entity.
type Door struct {
	ent.Schema
}

// Fields of the Door.
func (Door) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("direction").
			Values("north", "south", "east", "west", "up", "down"),
	}
}

// Edges of the Door.
func (Door) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("from", Room.Type).
			Unique(),
		edge.To("to", Room.Type).
			Unique(),
	}
}

// Mixin of the Door.
func (Door) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Entity{},
	}
}
