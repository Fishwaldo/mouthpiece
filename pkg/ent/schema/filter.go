package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Filter holds the schema definition for the Filter entity.
type Filter struct {
	ent.Schema
}

// Fields of the Filter.
func (Filter) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").
			NotEmpty().
			StructTag(`doc:"Name of the User`),
		field.String("Description").
			NotEmpty().
			StructTag(`doc:"Description of the Group`),
		field.Enum("Type").
			Values("InvalidFilter", "AppFilter", "UserFilter", "TransportFilter").
			Default("InvalidFilter").
			StructTag(`doc:"Type of Filter`),
		field.Bool("Enabled").
			Default(true).
			StructTag(`doc:"Is the Filter Enabled`),
		field.String("FilterImpl").
			NotEmpty().
			Immutable().
			StructTag(`doc:"Filter Implementation`),
	} 
}

// Edges of the Filter.
func (Filter) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("config", FilterConfig.Type).
		Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
		edge.To("groups", Group.Type),
		edge.From("app", App.Type).
			Ref("filters"),
		edge.From("user", User.Type).
			Ref("filters"),
	}
}

func (Filter) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}