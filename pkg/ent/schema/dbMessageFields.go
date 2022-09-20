package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DbMessageFields holds the schema definition for the DbMessageFields entity.
type DbMessageFields struct {
	ent.Schema
}

// Fields of the DbMessageFields.
func (DbMessageFields) Fields() []ent.Field {
	return []ent.Field {
		field.Text("Name").
			NotEmpty().
			Unique().
			StructTag(`doc:"Name of the Field"`),
		field.Text("Value").
			NotEmpty().
			StructTag(`doc:"Value of the Field"`),
		}
}

// Edges of the DbMessageFields.
func (DbMessageFields) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", DbMessage.Type).
		Ref("fields").
		Required().
		Unique(),
	}
}

func (DbMessageFields) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}