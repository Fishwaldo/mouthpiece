package schema

import (
//	"github.com/Fishwaldo/mouthpiece/pkg/ent/privacy"
//	"github.com/Fishwaldo/mouthpiece/pkg/ent/rules"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
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

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("TransportRecipients", TransportRecipient.Type),
		edge.From("users", User.Type).
			Ref("groups"),
		edge.From("filters", Filter.Type).
			Ref("groups"),
		edge.From("apps", App.Type).
			Ref("groups"),

	}
}

func (Group) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}

/*
func (Group) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			// Limit DenyMismatchedTenants only for
			// Create operations
			privacy.OnMutationOperation(
				rules.DenyMismatchedTenants(),
				ent.OpCreate,
			),
		},
	}
}
*/