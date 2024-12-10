package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Room holds the schema definition for the Room entity.
type Room struct {
	ent.Schema
}

// Fields of the Room.
func (Room) Fields() []ent.Field {
	return nil
}

// Edges of the Room.
func (Room) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("doors", Door.Type).
			Ref("from"),
		edge.From("doors_in", Door.Type).
			Ref("to"),
		edge.To("players", Player.Type),
	}
}

// Mixin of the Room.
func (Room) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Entity{},
	}
}
