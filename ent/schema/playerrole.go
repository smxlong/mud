package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PlayerRole holds the schema definition for the PlayerRole entity.
type PlayerRole struct {
	ent.Schema
}

// Fields of the PlayerRole.
func (PlayerRole) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Unique(),
	}
}

// Edges of the PlayerRole.
func (PlayerRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("players", Player.Type).
			Ref("player_roles"),
	}
}
