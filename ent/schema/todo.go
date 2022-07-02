package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").SchemaType(map[string]string{
			dialect.MySQL: "varchar(30)",
		}),
		field.String("content").SchemaType(map[string]string{
			dialect.MySQL: "text(300)",
		}),
		field.Int("user_id"),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("todos").
			Unique().
			Field("user_id").Required(),
	}
}
