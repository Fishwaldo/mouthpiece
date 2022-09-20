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

package dbtransportinstances

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TenantID applies equality check predicate on the "tenant_id" field. It's identical to TenantIDEQ.
func TenantID(v int) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTenantID), v))
	})
}

// Name applies equality check predicate on the "Name" field. It's identical to NameEQ.
func Name(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Description applies equality check predicate on the "Description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// Config applies equality check predicate on the "Config" field. It's identical to ConfigEQ.
func Config(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldConfig), v))
	})
}

// TransportProvider applies equality check predicate on the "TransportProvider" field. It's identical to TransportProviderEQ.
func TransportProvider(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTransportProvider), v))
	})
}

// TenantIDEQ applies the EQ predicate on the "tenant_id" field.
func TenantIDEQ(v int) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTenantID), v))
	})
}

// TenantIDNEQ applies the NEQ predicate on the "tenant_id" field.
func TenantIDNEQ(v int) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTenantID), v))
	})
}

// TenantIDIn applies the In predicate on the "tenant_id" field.
func TenantIDIn(vs ...int) predicate.DbTransportInstances {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTenantID), v...))
	})
}

// TenantIDNotIn applies the NotIn predicate on the "tenant_id" field.
func TenantIDNotIn(vs ...int) predicate.DbTransportInstances {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTenantID), v...))
	})
}

// NameEQ applies the EQ predicate on the "Name" field.
func NameEQ(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "Name" field.
func NameNEQ(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "Name" field.
func NameIn(vs ...string) predicate.DbTransportInstances {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "Name" field.
func NameNotIn(vs ...string) predicate.DbTransportInstances {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "Name" field.
func NameGT(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "Name" field.
func NameGTE(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "Name" field.
func NameLT(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "Name" field.
func NameLTE(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "Name" field.
func NameContains(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "Name" field.
func NameHasPrefix(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "Name" field.
func NameHasSuffix(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "Name" field.
func NameEqualFold(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "Name" field.
func NameContainsFold(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "Description" field.
func DescriptionEQ(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "Description" field.
func DescriptionNEQ(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "Description" field.
func DescriptionIn(vs ...string) predicate.DbTransportInstances {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "Description" field.
func DescriptionNotIn(vs ...string) predicate.DbTransportInstances {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "Description" field.
func DescriptionGT(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "Description" field.
func DescriptionGTE(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "Description" field.
func DescriptionLT(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "Description" field.
func DescriptionLTE(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "Description" field.
func DescriptionContains(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "Description" field.
func DescriptionHasPrefix(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "Description" field.
func DescriptionHasSuffix(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionIsNil applies the IsNil predicate on the "Description" field.
func DescriptionIsNil() predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDescription)))
	})
}

// DescriptionNotNil applies the NotNil predicate on the "Description" field.
func DescriptionNotNil() predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDescription)))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "Description" field.
func DescriptionEqualFold(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "Description" field.
func DescriptionContainsFold(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// ConfigEQ applies the EQ predicate on the "Config" field.
func ConfigEQ(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldConfig), v))
	})
}

// ConfigNEQ applies the NEQ predicate on the "Config" field.
func ConfigNEQ(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldConfig), v))
	})
}

// ConfigIn applies the In predicate on the "Config" field.
func ConfigIn(vs ...string) predicate.DbTransportInstances {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldConfig), v...))
	})
}

// ConfigNotIn applies the NotIn predicate on the "Config" field.
func ConfigNotIn(vs ...string) predicate.DbTransportInstances {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldConfig), v...))
	})
}

// ConfigGT applies the GT predicate on the "Config" field.
func ConfigGT(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldConfig), v))
	})
}

// ConfigGTE applies the GTE predicate on the "Config" field.
func ConfigGTE(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldConfig), v))
	})
}

// ConfigLT applies the LT predicate on the "Config" field.
func ConfigLT(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldConfig), v))
	})
}

// ConfigLTE applies the LTE predicate on the "Config" field.
func ConfigLTE(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldConfig), v))
	})
}

// ConfigContains applies the Contains predicate on the "Config" field.
func ConfigContains(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldConfig), v))
	})
}

// ConfigHasPrefix applies the HasPrefix predicate on the "Config" field.
func ConfigHasPrefix(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldConfig), v))
	})
}

// ConfigHasSuffix applies the HasSuffix predicate on the "Config" field.
func ConfigHasSuffix(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldConfig), v))
	})
}

// ConfigEqualFold applies the EqualFold predicate on the "Config" field.
func ConfigEqualFold(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldConfig), v))
	})
}

// ConfigContainsFold applies the ContainsFold predicate on the "Config" field.
func ConfigContainsFold(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldConfig), v))
	})
}

// TransportProviderEQ applies the EQ predicate on the "TransportProvider" field.
func TransportProviderEQ(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTransportProvider), v))
	})
}

// TransportProviderNEQ applies the NEQ predicate on the "TransportProvider" field.
func TransportProviderNEQ(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTransportProvider), v))
	})
}

// TransportProviderIn applies the In predicate on the "TransportProvider" field.
func TransportProviderIn(vs ...string) predicate.DbTransportInstances {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTransportProvider), v...))
	})
}

// TransportProviderNotIn applies the NotIn predicate on the "TransportProvider" field.
func TransportProviderNotIn(vs ...string) predicate.DbTransportInstances {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTransportProvider), v...))
	})
}

// TransportProviderGT applies the GT predicate on the "TransportProvider" field.
func TransportProviderGT(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTransportProvider), v))
	})
}

// TransportProviderGTE applies the GTE predicate on the "TransportProvider" field.
func TransportProviderGTE(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTransportProvider), v))
	})
}

// TransportProviderLT applies the LT predicate on the "TransportProvider" field.
func TransportProviderLT(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTransportProvider), v))
	})
}

// TransportProviderLTE applies the LTE predicate on the "TransportProvider" field.
func TransportProviderLTE(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTransportProvider), v))
	})
}

// TransportProviderContains applies the Contains predicate on the "TransportProvider" field.
func TransportProviderContains(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTransportProvider), v))
	})
}

// TransportProviderHasPrefix applies the HasPrefix predicate on the "TransportProvider" field.
func TransportProviderHasPrefix(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTransportProvider), v))
	})
}

// TransportProviderHasSuffix applies the HasSuffix predicate on the "TransportProvider" field.
func TransportProviderHasSuffix(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTransportProvider), v))
	})
}

// TransportProviderEqualFold applies the EqualFold predicate on the "TransportProvider" field.
func TransportProviderEqualFold(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTransportProvider), v))
	})
}

// TransportProviderContainsFold applies the ContainsFold predicate on the "TransportProvider" field.
func TransportProviderContainsFold(v string) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTransportProvider), v))
	})
}

// HasTenant applies the HasEdge predicate on the "tenant" edge.
func HasTenant() predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TenantTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TenantTable, TenantColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTenantWith applies the HasEdge predicate on the "tenant" edge with a given conditions (other predicates).
func HasTenantWith(preds ...predicate.Tenant) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
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

// HasTransportRecipients applies the HasEdge predicate on the "TransportRecipients" edge.
func HasTransportRecipients() predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TransportRecipientsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TransportRecipientsTable, TransportRecipientsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTransportRecipientsWith applies the HasEdge predicate on the "TransportRecipients" edge with a given conditions (other predicates).
func HasTransportRecipientsWith(preds ...predicate.DbTransportRecipients) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TransportRecipientsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TransportRecipientsTable, TransportRecipientsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DbTransportInstances) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DbTransportInstances) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
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
func Not(p predicate.DbTransportInstances) predicate.DbTransportInstances {
	return predicate.DbTransportInstances(func(s *sql.Selector) {
		p(s.Not())
	})
}
