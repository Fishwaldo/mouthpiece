package schema

import (
	"github.com/Fishwaldo/mouthpiece/pkg/ent/privacy"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/rules"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Tenant holds the schema definition for the Tenant entity.
type Tenant struct {
	ent.Schema
}

func (Tenant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Tenant.
func (Tenant) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
	}
}

// Edges of the Tenant.
func (Tenant) Edges() []ent.Edge {
	return nil
}

// Policy defines the privacy policy of the User.
func (Tenant) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			//For Tenant type, we only allow admin users to mutate
			//the tenant information and deny otherwise.
			rules.AllowIfGlobalAdmin(),
			privacy.AlwaysDenyRule(),
		},
	}
}
