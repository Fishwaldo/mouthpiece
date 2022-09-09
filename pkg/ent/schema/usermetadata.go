package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)
// UserMetaData holds the schema definition for the UserMetaData entity.
type UserMetaData struct {
	ent.Schema
}

// Fields of the UserMetaData.
func (UserMetaData) Fields() []ent.Field {
	return []ent.Field {
		field.Text("Name").
			NotEmpty().
			StructTag(`doc:"Name of the Field"`),
		field.Text("Value").
			NotEmpty().
			StructTag(`doc:"Value of the Field"`),
		}
}

// Edges of the UserMetaData.
func (UserMetaData) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
		Ref("metadata").
		Required().
		Unique(),
	}
}

// Mixin of the User schema.
func (UserMetaData) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}