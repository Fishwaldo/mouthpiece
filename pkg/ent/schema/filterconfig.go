package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)


// FilterConfig holds the schema definition for the FilterConfig entity.
type FilterConfig struct {
	ent.Schema
}

// Fields of the FilterConfig.
func (FilterConfig) Fields() []ent.Field {
	return []ent.Field {
		field.Text("Name").
			NotEmpty().
			StructTag(`doc:"Name of the Field"`),
		field.Text("Value").
			NotEmpty().
			StructTag(`doc:"Value of the Field"`),
		}
}

// Edges of the FilterConfig.
func (FilterConfig) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("filter", Filter.Type).
		Ref("config").
		Required().
		Unique(),
	}
}

func (FilterConfig) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}