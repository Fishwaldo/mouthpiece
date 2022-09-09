package rules

import (
	"context"
	"fmt"

	"github.com/Fishwaldo/mouthpiece/pkg/ent"
//	"github.com/Fishwaldo/mouthpiece/pkg/ent/hook"
)


func TenantCreateHook(next ent.Mutator) ent.Mutator{
	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		if tenant := FromContextGetTenant(ctx); tenant != nil {
			if err := m.SetField("tenant_id", tenant.ID); err != nil {
				return nil, fmt.Errorf("%s: failed setting tenant_id: %w", m.Type(), err)
			}		
		} 
		return next.Mutate(ctx, m)
	})
}
func TenantMutateHook(next ent.Mutator) ent.Mutator{
	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		if user := FromContext(ctx); user != nil {
			if user.GlobalAdmin() {
				return next.Mutate(ctx, m)
			}
		} else {
			return nil, fmt.Errorf("No user in context")
		}
		if m.FieldCleared("tenant_id") {
			return nil, fmt.Errorf("%s: cannot clear tenant_id", m.Type())
		}
		if _, changed := m.Field("tenant_id"); changed {
			return nil, fmt.Errorf("%s: cannot change tenant_id", m.Type())
		}
		return next.Mutate(ctx, m)
	})
}
