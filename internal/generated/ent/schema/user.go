package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("uuid"),
		field.String("email").
			Unique(),
		field.String("login").
			Unique(),
		field.String("password"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("metrics", Metric.Type),
		edge.To("given_accesses", Access.Type),
		edge.To("recieved_accesses", Access.Type),
	}
}

// Indexes if the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email").
			Unique(),
		index.Fields("login").
			Unique(),
	}
}
