package schema

import (
	
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"


)

// DbFilter holds the schema definition for the DbFilter entity.
type DbFilter struct {
	ent.Schema
}

// Fields of the Filter.
func (DbFilter) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").
			NotEmpty().
			StructTag(`doc:"Name of the User"`),
		field.String("Description").
			Optional().
			StructTag(`doc:"Description of the Filter"`),
		field.Enum("Type").
			GoType(interfaces.FilterType(interfaces.InvalidFilter)).	
			StructTag(`doc:"Type of Filter"`),
		field.Bool("Enabled").
			Default(true).
			StructTag(`doc:"Is the Filter Enabled"`),
		field.String("FilterImpl").
			NotEmpty().
			Immutable().
			StructTag(`doc:"Filter Implementation"`),
		field.String("Config").
			StructTag(`doc:"Filter Configuration"`),
	} 
}

// Edges of the Filter.
func (DbFilter) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("groups", DbGroup.Type),
		edge.From("app", DbApp.Type).
			Ref("filters"),
		edge.From("user", DbUser.Type).
			Ref("filters"),
	}
}

func (DbFilter) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}