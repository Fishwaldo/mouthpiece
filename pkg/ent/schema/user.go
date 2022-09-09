package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixin of the User schema.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("Email").
			NotEmpty().
			StructTag(`doc:"Email Address of the User`),
		field.String("Name").
			NotEmpty().
			StructTag(`doc:"Name of the User`),
		field.String("Description").
			NotEmpty().
			StructTag(`doc:"Description of the Group`),
	}
}

// Index of the User 
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("Email", "tenant_id").
			Unique(),
	}
}


// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("metadata", UserMetaData.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.To("filters", Filter.Type),
		edge.To("groups", Group.Type),
		edge.To("TransportRecipients", TransportRecipient.Type),
	}
}
