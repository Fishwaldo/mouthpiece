package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DbTransportRecipients holds the schema definition for the DbTransportRecipients entity.
type DbTransportRecipients struct {
	ent.Schema
}

// Fields of the DbTransportRecipients.
func (DbTransportRecipients) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").
			NotEmpty().
			Unique().
			StructTag(`doc:"Name of the Transport Recipient"`),
		field.String("Description").
			Optional().
			StructTag(`doc:"Description of the Transport Recipient"`),
		field.String("config").
			StructTag(`doc:"Config of the Transport Recipient"`),
	}
}

// Edges of the DbTransportRecipients.
func (DbTransportRecipients) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("TransportInstance", DbTransportInstances.Type).
			Ref("TransportRecipients").
			Required().
			Unique(),
		edge.From("GroupRecipient", DbGroup.Type).
			Ref("TransportRecipients").
			Unique(),
		edge.From("UserRecipient", DbUser.Type).
			Ref("TransportRecipients").
			Unique(),	
	}
}

func (DbTransportRecipients) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}