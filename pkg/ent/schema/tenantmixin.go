package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/hook"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/privacy"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/rules"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
)

// BaseMixin holds the schema definition for the BaseMixin entity.
type BaseMixin struct {
	mixin.Schema
}

//Policy defines the privacy policy of the BaseMixin.
func (BaseMixin) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			// Deny any operation in case there is no "viewer context".
			rules.DenyIfNoViewer(),
		},
		Mutation: privacy.MutationPolicy{
			// Deny any operation in case there is no "viewer context".
			rules.DenyIfNoViewer(),
		},
	}
}

// TenantMixin holds the schema definition for the TenantMixin entity.
type TenantMixin struct {
	mixin.Schema
}

// Fields of the TenantMixin.
func (TenantMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("tenant_id"),
		field.JSON("AppData", interfaces.AppData{}).
			Optional().
			Sensitive(),
	}
}

// Edges of the TenantMixin.
func (TenantMixin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tenant", Tenant.Type).
			Field("tenant_id").
			Required().
			Unique(),
	}
}

// Policy for all schemas that embed TenantMixin.
func (TenantMixin) Policy() ent.Policy {
	return rules.FilterTenantRule()
}

func (TenantMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			rules.TenantCreateHook,
			ent.OpCreate,
		),
		hook.On(
			rules.TenantMutateHook,
			ent.OpUpdateOne|ent.OpUpdate,
		),
	}
}
