/*
	MIT License

	Copyright (c) 2021 Justin Hammond

	Permission is hereby granted, free of charge, to any person obtaining a copy
	of this software and associated documentation files (the "Software"), to deal
	in the Software without restriction, including without limitation the rights
	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
	copies of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:

	The above copyright notice and this permission notice shall be included in all
	copies or substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
	SOFTWARE.
*/

// Code generated by entc, DO NOT EDIT.

package privacy

import (
	"context"
	"fmt"

	"github.com/Fishwaldo/mouthpiece/pkg/ent"

	"entgo.io/ent/entql"
	"entgo.io/ent/privacy"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with allow decision.
	Allow = privacy.Allow

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with deny decision.
	Deny = privacy.Deny

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = privacy.Skip
)

// Allowf returns an formatted wrapped Allow decision.
func Allowf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Allow)...)
}

// Denyf returns an formatted wrapped Deny decision.
func Denyf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Deny)...)
}

// Skipf returns an formatted wrapped Skip decision.
func Skipf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Skip)...)
}

// DecisionContext creates a new context from the given parent context with
// a policy decision attach to it.
func DecisionContext(parent context.Context, decision error) context.Context {
	return privacy.DecisionContext(parent, decision)
}

// DecisionFromContext retrieves the policy decision from the context.
func DecisionFromContext(ctx context.Context) (error, bool) {
	return privacy.DecisionFromContext(ctx)
}

type (
	// Policy groups query and mutation policies.
	Policy = privacy.Policy

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule = privacy.QueryRule
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy = privacy.QueryPolicy

	// MutationRule defines the interface which decides whether a
	// mutation is allowed and optionally modifies it.
	MutationRule = privacy.MutationRule
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy = privacy.MutationPolicy
)

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

// MutationRuleFunc type is an adapter which allows the use of
// ordinary functions as mutation rules.
type MutationRuleFunc func(context.Context, ent.Mutation) error

// EvalMutation returns f(ctx, m).
func (f MutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return f(ctx, m)
}

// QueryMutationRule is an interface which groups query and mutation rules.
type QueryMutationRule interface {
	QueryRule
	MutationRule
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return fixedDecision{Allow}
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return fixedDecision{Deny}
}

type fixedDecision struct {
	decision error
}

func (f fixedDecision) EvalQuery(context.Context, ent.Query) error {
	return f.decision
}

func (f fixedDecision) EvalMutation(context.Context, ent.Mutation) error {
	return f.decision
}

type contextDecision struct {
	eval func(context.Context) error
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return contextDecision{eval}
}

func (c contextDecision) EvalQuery(ctx context.Context, _ ent.Query) error {
	return c.eval(ctx)
}

func (c contextDecision) EvalMutation(ctx context.Context, _ ent.Mutation) error {
	return c.eval(ctx)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op ent.Op) MutationRule {
	return MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		if m.Op().Is(op) {
			return rule.EvalMutation(ctx, m)
		}
		return Skip
	})
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op ent.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		return Denyf("ent/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The AppQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AppQueryRuleFunc func(context.Context, *ent.AppQuery) error

// EvalQuery return f(ctx, q).
func (f AppQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AppQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AppQuery", q)
}

// The AppMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AppMutationRuleFunc func(context.Context, *ent.AppMutation) error

// EvalMutation calls f(ctx, m).
func (f AppMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AppMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AppMutation", m)
}

// The FilterQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type FilterQueryRuleFunc func(context.Context, *ent.FilterQuery) error

// EvalQuery return f(ctx, q).
func (f FilterQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.FilterQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.FilterQuery", q)
}

// The FilterMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type FilterMutationRuleFunc func(context.Context, *ent.FilterMutation) error

// EvalMutation calls f(ctx, m).
func (f FilterMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.FilterMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.FilterMutation", m)
}

// The FilterConfigQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type FilterConfigQueryRuleFunc func(context.Context, *ent.FilterConfigQuery) error

// EvalQuery return f(ctx, q).
func (f FilterConfigQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.FilterConfigQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.FilterConfigQuery", q)
}

// The FilterConfigMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type FilterConfigMutationRuleFunc func(context.Context, *ent.FilterConfigMutation) error

// EvalMutation calls f(ctx, m).
func (f FilterConfigMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.FilterConfigMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.FilterConfigMutation", m)
}

// The GroupQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type GroupQueryRuleFunc func(context.Context, *ent.GroupQuery) error

// EvalQuery return f(ctx, q).
func (f GroupQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.GroupQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.GroupQuery", q)
}

// The GroupMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type GroupMutationRuleFunc func(context.Context, *ent.GroupMutation) error

// EvalMutation calls f(ctx, m).
func (f GroupMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.GroupMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.GroupMutation", m)
}

// The MessageQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type MessageQueryRuleFunc func(context.Context, *ent.MessageQuery) error

// EvalQuery return f(ctx, q).
func (f MessageQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.MessageQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.MessageQuery", q)
}

// The MessageMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type MessageMutationRuleFunc func(context.Context, *ent.MessageMutation) error

// EvalMutation calls f(ctx, m).
func (f MessageMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.MessageMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.MessageMutation", m)
}

// The MsgVarQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type MsgVarQueryRuleFunc func(context.Context, *ent.MsgVarQuery) error

// EvalQuery return f(ctx, q).
func (f MsgVarQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.MsgVarQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.MsgVarQuery", q)
}

// The MsgVarMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type MsgVarMutationRuleFunc func(context.Context, *ent.MsgVarMutation) error

// EvalMutation calls f(ctx, m).
func (f MsgVarMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.MsgVarMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.MsgVarMutation", m)
}

// The TenantQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type TenantQueryRuleFunc func(context.Context, *ent.TenantQuery) error

// EvalQuery return f(ctx, q).
func (f TenantQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TenantQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.TenantQuery", q)
}

// The TenantMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type TenantMutationRuleFunc func(context.Context, *ent.TenantMutation) error

// EvalMutation calls f(ctx, m).
func (f TenantMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.TenantMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.TenantMutation", m)
}

// The TransportInstanceQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type TransportInstanceQueryRuleFunc func(context.Context, *ent.TransportInstanceQuery) error

// EvalQuery return f(ctx, q).
func (f TransportInstanceQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TransportInstanceQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.TransportInstanceQuery", q)
}

// The TransportInstanceMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type TransportInstanceMutationRuleFunc func(context.Context, *ent.TransportInstanceMutation) error

// EvalMutation calls f(ctx, m).
func (f TransportInstanceMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.TransportInstanceMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.TransportInstanceMutation", m)
}

// The TransportRecipientQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type TransportRecipientQueryRuleFunc func(context.Context, *ent.TransportRecipientQuery) error

// EvalQuery return f(ctx, q).
func (f TransportRecipientQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TransportRecipientQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.TransportRecipientQuery", q)
}

// The TransportRecipientMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type TransportRecipientMutationRuleFunc func(context.Context, *ent.TransportRecipientMutation) error

// EvalMutation calls f(ctx, m).
func (f TransportRecipientMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.TransportRecipientMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.TransportRecipientMutation", m)
}

// The UserQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type UserQueryRuleFunc func(context.Context, *ent.UserQuery) error

// EvalQuery return f(ctx, q).
func (f UserQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.UserQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.UserQuery", q)
}

// The UserMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type UserMutationRuleFunc func(context.Context, *ent.UserMutation) error

// EvalMutation calls f(ctx, m).
func (f UserMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.UserMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.UserMutation", m)
}

// The UserMetaDataQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type UserMetaDataQueryRuleFunc func(context.Context, *ent.UserMetaDataQuery) error

// EvalQuery return f(ctx, q).
func (f UserMetaDataQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.UserMetaDataQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.UserMetaDataQuery", q)
}

// The UserMetaDataMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type UserMetaDataMutationRuleFunc func(context.Context, *ent.UserMetaDataMutation) error

// EvalMutation calls f(ctx, m).
func (f UserMetaDataMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.UserMetaDataMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.UserMetaDataMutation", m)
}

type (
	// Filter is the interface that wraps the Where function
	// for filtering nodes in queries and mutations.
	Filter interface {
		// Where applies a filter on the executed query/mutation.
		Where(entql.P)
	}

	// The FilterFunc type is an adapter that allows the use of ordinary
	// functions as filters for query and mutation types.
	FilterFunc func(context.Context, Filter) error
)

// EvalQuery calls f(ctx, q) if the query implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	fr, err := queryFilter(q)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

// EvalMutation calls f(ctx, q) if the mutation implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	fr, err := mutationFilter(m)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

var _ QueryMutationRule = FilterFunc(nil)

func queryFilter(q ent.Query) (Filter, error) {
	switch q := q.(type) {
	case *ent.AppQuery:
		return q.Filter(), nil
	case *ent.FilterQuery:
		return q.Filter(), nil
	case *ent.FilterConfigQuery:
		return q.Filter(), nil
	case *ent.GroupQuery:
		return q.Filter(), nil
	case *ent.MessageQuery:
		return q.Filter(), nil
	case *ent.MsgVarQuery:
		return q.Filter(), nil
	case *ent.TenantQuery:
		return q.Filter(), nil
	case *ent.TransportInstanceQuery:
		return q.Filter(), nil
	case *ent.TransportRecipientQuery:
		return q.Filter(), nil
	case *ent.UserQuery:
		return q.Filter(), nil
	case *ent.UserMetaDataQuery:
		return q.Filter(), nil
	default:
		return nil, Denyf("ent/privacy: unexpected query type %T for query filter", q)
	}
}

func mutationFilter(m ent.Mutation) (Filter, error) {
	switch m := m.(type) {
	case *ent.AppMutation:
		return m.Filter(), nil
	case *ent.FilterMutation:
		return m.Filter(), nil
	case *ent.FilterConfigMutation:
		return m.Filter(), nil
	case *ent.GroupMutation:
		return m.Filter(), nil
	case *ent.MessageMutation:
		return m.Filter(), nil
	case *ent.MsgVarMutation:
		return m.Filter(), nil
	case *ent.TenantMutation:
		return m.Filter(), nil
	case *ent.TransportInstanceMutation:
		return m.Filter(), nil
	case *ent.TransportRecipientMutation:
		return m.Filter(), nil
	case *ent.UserMutation:
		return m.Filter(), nil
	case *ent.UserMetaDataMutation:
		return m.Filter(), nil
	default:
		return nil, Denyf("ent/privacy: unexpected mutation type %T for mutation filter", m)
	}
}
