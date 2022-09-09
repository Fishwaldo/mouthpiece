package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TransportRecipients holds the schema definition for the TransportRecipients entity.
type TransportRecipient struct {
	ent.Schema
}

// Fields of the TransportRecipients.
func (TransportRecipient) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").
			NotEmpty().
			Unique().
			StructTag(`doc:"Name of the Group`),
		field.String("Description").
			NotEmpty().
			StructTag(`doc:"Description of the Group`),
	}
}

// Edges of the TransportRecipients.
func (TransportRecipient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("TransportInstance", TransportInstance.Type).
			Ref("TransportRecipients").
			Required().
			Unique(),
		edge.From("AppRecipient", App.Type).
			Ref("TransportRecipients"),
		edge.From("GroupRecipient", Group.Type).
			Ref("TransportRecipients"),
		edge.From("UserRecipient", User.Type).
			Ref("TransportRecipients"),			
	}
}

func (TransportRecipient) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}