package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DbTransportInstances holds the schema definition for the DbTransportInstances entity.
type DbTransportInstances struct {
	ent.Schema
}

// Fields of the TransportInstances.
func (DbTransportInstances) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").
			NotEmpty().
			Unique().
			StructTag(`doc:"Name of the Transport Instance`),
		field.String("Description").
			Optional().
			StructTag(`doc:"Description of the Transport Instance`),
		field.String("Config").
			NotEmpty().
			StructTag(`doc:"Config of the Transport Instance`),
		field.String("TransportProvider").
			NotEmpty().
			StructTag(`doc:"The Transport Provider`),
	} 
}

// Edges of the TransportInstances.
func (DbTransportInstances) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("TransportRecipients", DbTransportRecipients.Type).
		Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
	}
}

func (DbTransportInstances) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}