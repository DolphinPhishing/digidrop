package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// FileMiddleware holds the schema definition for the FileMiddleware entity.
type FileMiddleware struct {
	ent.Schema
}

// Fields of the FileMiddleware.
func (FileMiddleware) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("url_id"),
		field.String("file_path"),
		field.Bool("accessed").
			Default(false),
	}
}

// Edges of the FileMiddleware.
func (FileMiddleware) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("FileMiddlewareToUser", User.Type).
			Unique(),
	}
}
