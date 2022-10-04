package schema

import (
	//	"github.com/Fishwaldo/mouthpiece/pkg/ent/privacy"
	//	"github.com/Fishwaldo/mouthpiece/pkg/ent/rules"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DbGroup holds the schema definition for the DbGroup entity.
type DbGroup struct {
	ent.Schema
}

// Fields of the DbGroup.
func (DbGroup) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").
			NotEmpty().
			Unique().
			StructTag(`doc:"Name of the Group`),
		field.String("Description").
			Optional().
			StructTag(`doc:"Description of the Group`),
	}
}

// Edges of the Group.
func (DbGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("TransportRecipients", DbTransportRecipients.Type),
		edge.From("users", DbUser.Type).
			Ref("groups"),
		edge.From("filters", DbFilter.Type).
			Ref("groups"),
		edge.From("apps", DbApp.Type).
			Ref("groups"),
	}
}

func (DbGroup) Mixin() []ent.Mixin {
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
