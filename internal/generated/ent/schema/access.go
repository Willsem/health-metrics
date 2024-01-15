package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Access holds the schema definition for the Access entity.
type Access struct {
	ent.Schema
}

// Fields of the Access.
func (Access) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("uuid"),
	}
}

// Edges of the Access.
func (Access) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("given_accesses").
			Unique(),
		edge.From("user", User.Type).
			Ref("recieved_accesses").
			Unique(),
	}
}
