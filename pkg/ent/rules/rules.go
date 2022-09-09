package rules

import (
	"context"

	"entgo.io/ent/entql"
	"github.com/Fishwaldo/mouthpiece/pkg/ent"
//	"github.com/Fishwaldo/mouthpiece/pkg/ent/app"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/privacy"
)

// DenyIfNoViewer is a rule that returns deny decision if the viewer is missing in the context.
func DenyIfNoViewer() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		view := FromContext(ctx)
		if view == nil {
			return privacy.Denyf("viewer-context is missing")
		}
		// Skip to the next privacy rule (equivalent to return nil).
		return privacy.Skip
	})
}

// AllowIfGlobalAdmin is a rule that returns allow decision if the viewer is admin.
func AllowIfGlobalAdmin() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		//fmt.Printf("Allow If GAdmin\n")
		view := FromContext(ctx)
		if view.GlobalAdmin() {
			return privacy.Allow
		}
		// Skip to the next privacy rule (equivalent to return nil). 
		return privacy.Skip
	})
}

// AllowIfGlobalAdmin is a rule that returns allow decision if the viewer is admin.
func AllowIfAdmin() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		//fmt.Printf("Allow If Admin\n")
		view := FromContext(ctx)
		if view.Admin() {
			return privacy.Allow
		}
		// Skip to the next privacy rule (equivalent to return nil). 
		return privacy.Skip
	})
}


// FilterTenantRule is a query/mutation rule that filters out entities that are not in the tenant.
func FilterTenantRule() privacy.QueryMutationRule {
	// TenantsFilter is an interface to wrap WhereTenantID()
	// predicate that is used by both schemas.
	type TenantsFilter interface {
		WhereTenantID(entql.IntP)
	}
//	fmt.Printf("blah\n")
	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		view := FromContext(ctx)
		tid, ok := view.Tenant()
		if !ok {
			return privacy.Denyf("missing tenant information in viewer")
		}
		tf, ok := f.(TenantsFilter)
		if !ok {
			return privacy.Denyf("unexpected filter type %T", f)
		}
		// Make sure that a tenant reads only entities that have an edge to it.
		tf.WhereTenantID(entql.IntEQ(tid))
		// Skip to the next privacy rule (equivalent to return nil). 
		return privacy.Skip
	})
}
 
// DenyMismatchedTenants is a rule that runs only on create operations and returns a deny
// decision if the operation tries to add users to groups that are not in the same tenant.
func DenyMismatchedTenants() privacy.MutationRule {
	return privacy.GroupMutationRuleFunc(func(ctx context.Context, m *ent.GroupMutation) error {
		/// TODO: Fix this up
		// tid, exists := m.TenantID()
		// if !exists {
		// 	return privacy.Denyf("missing tenant information in mutation")
		// }
		// users := m.UsersIDs()
		// // If there are no users in the mutation, skip this rule-check.
		// if len(users) == 0 {
		// 	return privacy.Skip
		// }
		// // Query the tenant-ids of all attached users. Expect all users to be connected to the same tenant
		// // as the group. Note, we use privacy.DecisionContext to skip the FilterTenantRule defined above.
		// ids, err := m.Client().User.Query().Where(user.IDIn(users...)).Select(user.FieldTenantID).Ints(privacy.DecisionContext(ctx, privacy.Allow))
		// if err != nil {
		// 	return privacy.Denyf("querying the tenant-ids %v", err)
		// }
		// if len(ids) != len(users) {
		// 	return privacy.Denyf("one the attached users is not connected to a tenant %v", err)
		// }
		// for _, id := range ids {
		// 	if id != tid {
		// 		return privacy.Denyf("mismatch tenant-ids for group/users %d != %d", tid, id)
		// 	}
		// }
		// // Skip to the next privacy rule (equivalent to return nil).
		return privacy.Skip
	})
}

// Role for viewer actions.
type Role int

// List of roles.
const (
	_ Role = 1 << iota
	GlobalAdmin
	Admin
	View
)

// Viewer describes the query/mutation viewer-context.
type Viewer interface {
	Admin() bool         // If viewer is admin.
	GlobalAdmin() bool   // If viewer is global admin.
	Tenant() (int, bool) // Tenant identifier.
}

// UserViewer describes a user-viewer.
type UserViewer struct {
	T    *ent.Tenant
	Role Role // Attached roles.
}

func (v UserViewer) Admin() bool {
	if v.Role & GlobalAdmin == GlobalAdmin {
		return true
	}
	if v.Role & Admin == Admin {
		return true
	} 
	return false
}

func (v UserViewer) GlobalAdmin() bool {
	return v.Role & GlobalAdmin == GlobalAdmin
}

func (v UserViewer) Tenant() (int, bool) {
	 if v.T != nil {
	 	return v.T.ID, true
	 }
	return 0, false
}

type ctxKey struct{}

// FromContext returns the Viewer stored in a context.
func FromContext(ctx context.Context) Viewer {
	v, _ := ctx.Value(ctxKey{}).(Viewer)
	return v
}

// NewContext returns a copy of parent context with the given Viewer attached with it.
func NewContext(parent context.Context, v Viewer) context.Context {
	return context.WithValue(parent, ctxKey{}, v)
}

func FromContextGetTenant(ctx context.Context) (*ent.Tenant) {
	v := FromContext(ctx)
	_, ok := v.Tenant()
	if !ok {
		return nil
	}
	return v.(*UserViewer).T
}
