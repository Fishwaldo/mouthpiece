package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MsgVar holds the schema definition for the MsgVar entity.
type MsgVar struct {
	ent.Schema
}

// Fields of the MsgVar.
func (MsgVar) Fields() []ent.Field {
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

// Edges of the MsgVar.
func (MsgVar) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Message.Type).
		Ref("vars").
		Required().
		Unique(),
	}
}

func (MsgVar) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}