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

package transportrecipient

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TenantID applies equality check predicate on the "tenant_id" field. It's identical to TenantIDEQ.
func TenantID(v int) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTenantID), v))
	})
}

// Name applies equality check predicate on the "Name" field. It's identical to NameEQ.
func Name(v string) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Description applies equality check predicate on the "Description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// TenantIDEQ applies the EQ predicate on the "tenant_id" field.
func TenantIDEQ(v int) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTenantID), v))
	})
}

// TenantIDNEQ applies the NEQ predicate on the "tenant_id" field.
func TenantIDNEQ(v int) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTenantID), v))
	})
}

// TenantIDIn applies the In predicate on the "tenant_id" field.
func TenantIDIn(vs ...int) predicate.TransportRecipient {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTenantID), v...))
	})
}

// TenantIDNotIn applies the NotIn predicate on the "tenant_id" field.
func TenantIDNotIn(vs ...int) predicate.TransportRecipient {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTenantID), v...))
	})
}

// NameEQ applies the EQ predicate on the "Name" field.
func NameEQ(v string) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "Name" field.
func NameNEQ(v string) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "Name" field.
func NameIn(vs ...string) predicate.TransportRecipient {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "Name" field.
func NameNotIn(vs ...string) predicate.TransportRecipient {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "Name" field.
func NameGT(v string) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "Name" field.
func NameGTE(v string) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "Name" field.
func NameLT(v string) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "Name" field.
func NameLTE(v string) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "Name" field.
func NameContains(v string) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "Name" field.
func NameHasPrefix(v string) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "Name" field.
func NameHasSuffix(v string) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "Name" field.
func NameEqualFold(v string) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "Name" field.
func NameContainsFold(v string) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "Description" field.
func DescriptionEQ(v string) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "Description" field.
func DescriptionNEQ(v string) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "Description" field.
func DescriptionIn(vs ...string) predicate.TransportRecipient {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "Description" field.
func DescriptionNotIn(vs ...string) predicate.TransportRecipient {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "Description" field.
func DescriptionGT(v string) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "Description" field.
func DescriptionGTE(v string) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "Description" field.
func DescriptionLT(v string) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "Description" field.
func DescriptionLTE(v string) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "Description" field.
func DescriptionContains(v string) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "Description" field.
func DescriptionHasPrefix(v string) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "Description" field.
func DescriptionHasSuffix(v string) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "Description" field.
func DescriptionEqualFold(v string) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "Description" field.
func DescriptionContainsFold(v string) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// HasTenant applies the HasEdge predicate on the "tenant" edge.
func HasTenant() predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TenantTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TenantTable, TenantColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTenantWith applies the HasEdge predicate on the "tenant" edge with a given conditions (other predicates).
func HasTenantWith(preds ...predicate.Tenant) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TenantInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TenantTable, TenantColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTransportInstance applies the HasEdge predicate on the "TransportInstance" edge.
func HasTransportInstance() predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TransportInstanceTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TransportInstanceTable, TransportInstanceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTransportInstanceWith applies the HasEdge predicate on the "TransportInstance" edge with a given conditions (other predicates).
func HasTransportInstanceWith(preds ...predicate.TransportInstance) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TransportInstanceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TransportInstanceTable, TransportInstanceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAppRecipient applies the HasEdge predicate on the "AppRecipient" edge.
func HasAppRecipient() predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AppRecipientTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, AppRecipientTable, AppRecipientPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAppRecipientWith applies the HasEdge predicate on the "AppRecipient" edge with a given conditions (other predicates).
func HasAppRecipientWith(preds ...predicate.App) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AppRecipientInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, AppRecipientTable, AppRecipientPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGroupRecipient applies the HasEdge predicate on the "GroupRecipient" edge.
func HasGroupRecipient() predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GroupRecipientTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, GroupRecipientTable, GroupRecipientPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupRecipientWith applies the HasEdge predicate on the "GroupRecipient" edge with a given conditions (other predicates).
func HasGroupRecipientWith(preds ...predicate.Group) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GroupRecipientInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, GroupRecipientTable, GroupRecipientPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserRecipient applies the HasEdge predicate on the "UserRecipient" edge.
func HasUserRecipient() predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserRecipientTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UserRecipientTable, UserRecipientPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserRecipientWith applies the HasEdge predicate on the "UserRecipient" edge with a given conditions (other predicates).
func HasUserRecipientWith(preds ...predicate.User) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserRecipientInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UserRecipientTable, UserRecipientPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TransportRecipient) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TransportRecipient) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TransportRecipient) predicate.TransportRecipient {
	return predicate.TransportRecipient(func(s *sql.Selector) {
		p(s.Not())
	})
}
