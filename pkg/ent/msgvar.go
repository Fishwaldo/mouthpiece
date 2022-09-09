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

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/message"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/msgvar"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/tenant"
	"github.com/Fishwaldo/mouthpiece/pkg/validate"
	"github.com/google/uuid"
)

// MsgVar is the model entity for the MsgVar schema.
type MsgVar struct {
	config `doc:"-" json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// TenantID holds the value of the "tenant_id" field.
	TenantID int `json:"tenant_id,omitempty"`
	// Name holds the value of the "Name" field.
	Name string `json:"Name,omitempty" doc:"Name of the Field"`
	// Value holds the value of the "Value" field.
	Value string `json:"Value,omitempty" doc:"Value of the Field"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MsgVarQuery when eager-loading is set.
	Edges        MsgVarEdges `json:"edges"`
	message_vars *uuid.UUID
}

// MsgVarEdges holds the relations/edges for other nodes in the graph.
type MsgVarEdges struct {
	// Tenant holds the value of the tenant edge.
	Tenant *Tenant `json:"tenant,omitempty"`
	// Owner holds the value of the owner edge.
	Owner *Message `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TenantOrErr returns the Tenant value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MsgVarEdges) TenantOrErr() (*Tenant, error) {
	if e.loadedTypes[0] {
		if e.Tenant == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: tenant.Label}
		}
		return e.Tenant, nil
	}
	return nil, &NotLoadedError{edge: "tenant"}
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MsgVarEdges) OwnerOrErr() (*Message, error) {
	if e.loadedTypes[1] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: message.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MsgVar) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case msgvar.FieldID, msgvar.FieldTenantID:
			values[i] = new(sql.NullInt64)
		case msgvar.FieldName, msgvar.FieldValue:
			values[i] = new(sql.NullString)
		case msgvar.ForeignKeys[0]: // message_vars
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type MsgVar", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MsgVar fields.
func (mv *MsgVar) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case msgvar.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mv.ID = int(value.Int64)
		case msgvar.FieldTenantID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_id", values[i])
			} else if value.Valid {
				mv.TenantID = int(value.Int64)
			}
		case msgvar.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Name", values[i])
			} else if value.Valid {
				mv.Name = value.String
			}
		case msgvar.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Value", values[i])
			} else if value.Valid {
				mv.Value = value.String
			}
		case msgvar.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field message_vars", values[i])
			} else if value.Valid {
				mv.message_vars = new(uuid.UUID)
				*mv.message_vars = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryTenant queries the "tenant" edge of the MsgVar entity.
func (mv *MsgVar) QueryTenant() *TenantQuery {
	return (&MsgVarClient{config: mv.config}).QueryTenant(mv)
}

// QueryOwner queries the "owner" edge of the MsgVar entity.
func (mv *MsgVar) QueryOwner() *MessageQuery {
	return (&MsgVarClient{config: mv.config}).QueryOwner(mv)
}

// Update returns a builder for updating this MsgVar.
// Note that you need to call MsgVar.Unwrap() before calling this method if this MsgVar
// was returned from a transaction, and the transaction was committed or rolled back.
func (mv *MsgVar) Update() *MsgVarUpdateOne {
	return (&MsgVarClient{config: mv.config}).UpdateOne(mv)
}

// Unwrap unwraps the MsgVar entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mv *MsgVar) Unwrap() *MsgVar {
	_tx, ok := mv.config.driver.(*txDriver)
	if !ok {
		panic("ent: MsgVar is not a transactional entity")
	}
	mv.config.driver = _tx.drv
	return mv
}

// String implements the fmt.Stringer.
func (mv *MsgVar) String() string {
	var builder strings.Builder
	builder.WriteString("MsgVar(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mv.ID))
	builder.WriteString("tenant_id=")
	builder.WriteString(fmt.Sprintf("%v", mv.TenantID))
	builder.WriteString(", ")
	builder.WriteString("Name=")
	builder.WriteString(mv.Name)
	builder.WriteString(", ")
	builder.WriteString("Value=")
	builder.WriteString(mv.Value)
	builder.WriteByte(')')
	return builder.String()
}

func (mv *MsgVar) ValidateMsgVar() error {
	if err := validate.Get().Struct(mv); err != nil {
		return err
	}
	return nil
}

// MsgVars is a parsable slice of MsgVar.
type MsgVars []*MsgVar

func (mv MsgVars) config(cfg config) {
	for _i := range mv {
		mv[_i].config = cfg
	}
}