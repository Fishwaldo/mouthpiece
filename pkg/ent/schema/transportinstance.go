package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TransportInstances holds the schema definition for the TransportInstances entity.
type TransportInstance struct {
	ent.Schema
}

// Fields of the TransportInstances.
func (TransportInstance) Fields() []ent.Field {
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

// Edges of the TransportInstances.
func (TransportInstance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("TransportRecipients", TransportRecipient.Type).
		Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
	}
}

func (TransportInstance) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}