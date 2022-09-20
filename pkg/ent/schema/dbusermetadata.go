package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)
// DbUserMetaData holds the schema definition for the DbUserMetaData entity.
type DbUserMetaData struct {
	ent.Schema
}

// Fields of the DbUserMetaData.
func (DbUserMetaData) Fields() []ent.Field {
	return []ent.Field {
		field.Text("Name").
			NotEmpty().
			StructTag(`doc:"Name of the Field"`),
		field.Text("Value").
			NotEmpty().
			StructTag(`doc:"Value of the Field"`),
		}
}

// Edges of the DbUserMetaData.
func (DbUserMetaData) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", DbUser.Type).
		Ref("metadata").
		Required().
		Unique(),
	}
}

// Mixin of the User schema.
func (DbUserMetaData) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}