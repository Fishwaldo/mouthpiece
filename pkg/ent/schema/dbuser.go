package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// DbUser holds the schema definition for the DbUser entity.
type DbUser struct {
	ent.Schema
}

// Mixin of the User schema.
func (DbUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}

// Fields of the User.
func (DbUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("Email").
			NotEmpty().
			StructTag(`doc:"Email Address of the User`),
		field.String("Name").
			NotEmpty().
			StructTag(`doc:"Name of the User`),
		field.String("Description").
			Optional().
			StructTag(`doc:"Description of the Group`),
	}
}

// Index of the User 
func (DbUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("Email", "tenant_id").
			Unique(),
	}
}


// Edges of the User.
func (DbUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("metadata", DbUserMetaData.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.To("filters", DbFilter.Type),
		edge.To("groups", DbGroup.Type),
		edge.To("TransportRecipients", DbTransportRecipients.Type),
	}
}
