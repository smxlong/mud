package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Player holds the schema definition for the Player entity.
type Player struct {
	ent.Schema
}

// Fields of the Player.
func (Player) Fields() []ent.Field {
	return []ent.Field{
		field.String("password").
			Sensitive(),
		field.String("email").
			NotEmpty().
			Unique(),
	}
}

// Edges of the Player.
func (Player) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("room", Room.Type).
			Ref("players").
			Unique(),
		edge.To("player_roles", PlayerRole.Type),
	}
}

// Mixin of the Player.
func (Player) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Entity{},
	}
}
