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
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbmessage"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbmessagefields"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/tenant"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/validate"
	"github.com/google/uuid"
)

// DbMessageFields is the model entity for the DbMessageFields schema.
type DbMessageFields struct {
	config `doc:"-" json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// TenantID holds the value of the "tenant_id" field.
	TenantID int `json:"tenant_id,omitempty"`
	// AppData holds the value of the "AppData" field.
	AppData interfaces.AppData `doc:"-" json:"-"`
	// Name holds the value of the "Name" field.
	Name string `json:"Name,omitempty" doc:"Name of the Field"`
	// Value holds the value of the "Value" field.
	Value string `json:"Value,omitempty" doc:"Value of the Field"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DbMessageFieldsQuery when eager-loading is set.
	Edges             DbMessageFieldsEdges `json:"edges"`
	db_message_fields *uuid.UUID
}

// DbMessageFieldsEdges holds the relations/edges for other nodes in the graph.
type DbMessageFieldsEdges struct {
	// Tenant holds the value of the tenant edge.
	Tenant *Tenant `json:"tenant,omitempty"`
	// Owner holds the value of the owner edge.
	Owner *DbMessage `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TenantOrErr returns the Tenant value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DbMessageFieldsEdges) TenantOrErr() (*Tenant, error) {
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
func (e DbMessageFieldsEdges) OwnerOrErr() (*DbMessage, error) {
	if e.loadedTypes[1] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: dbmessage.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DbMessageFields) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case dbmessagefields.FieldAppData:
			values[i] = new([]byte)
		case dbmessagefields.FieldID, dbmessagefields.FieldTenantID:
			values[i] = new(sql.NullInt64)
		case dbmessagefields.FieldName, dbmessagefields.FieldValue:
			values[i] = new(sql.NullString)
		case dbmessagefields.ForeignKeys[0]: // db_message_fields
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type DbMessageFields", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DbMessageFields fields.
func (dmf *DbMessageFields) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case dbmessagefields.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			dmf.ID = int(value.Int64)
		case dbmessagefields.FieldTenantID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_id", values[i])
			} else if value.Valid {
				dmf.TenantID = int(value.Int64)
			}
		case dbmessagefields.FieldAppData:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field AppData", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &dmf.AppData); err != nil {
					return fmt.Errorf("unmarshal field AppData: %w", err)
				}
			}
		case dbmessagefields.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Name", values[i])
			} else if value.Valid {
				dmf.Name = value.String
			}
		case dbmessagefields.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Value", values[i])
			} else if value.Valid {
				dmf.Value = value.String
			}
		case dbmessagefields.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field db_message_fields", values[i])
			} else if value.Valid {
				dmf.db_message_fields = new(uuid.UUID)
				*dmf.db_message_fields = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryTenant queries the "tenant" edge of the DbMessageFields entity.
func (dmf *DbMessageFields) QueryTenant() *TenantQuery {
	return (&DbMessageFieldsClient{config: dmf.config}).QueryTenant(dmf)
}

// QueryOwner queries the "owner" edge of the DbMessageFields entity.
func (dmf *DbMessageFields) QueryOwner() *DbMessageQuery {
	return (&DbMessageFieldsClient{config: dmf.config}).QueryOwner(dmf)
}

// Update returns a builder for updating this DbMessageFields.
// Note that you need to call DbMessageFields.Unwrap() before calling this method if this DbMessageFields
// was returned from a transaction, and the transaction was committed or rolled back.
func (dmf *DbMessageFields) Update() *DbMessageFieldsUpdateOne {
	return (&DbMessageFieldsClient{config: dmf.config}).UpdateOne(dmf)
}

// Unwrap unwraps the DbMessageFields entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dmf *DbMessageFields) Unwrap() *DbMessageFields {
	_tx, ok := dmf.config.driver.(*txDriver)
	if !ok {
		panic("ent: DbMessageFields is not a transactional entity")
	}
	dmf.config.driver = _tx.drv
	return dmf
}

// String implements the fmt.Stringer.
func (dmf *DbMessageFields) String() string {
	var builder strings.Builder
	builder.WriteString("DbMessageFields(")
	builder.WriteString(fmt.Sprintf("id=%v, ", dmf.ID))
	builder.WriteString("tenant_id=")
	builder.WriteString(fmt.Sprintf("%v", dmf.TenantID))
	builder.WriteString(", ")
	builder.WriteString("AppData=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("Name=")
	builder.WriteString(dmf.Name)
	builder.WriteString(", ")
	builder.WriteString("Value=")
	builder.WriteString(dmf.Value)
	builder.WriteByte(')')
	return builder.String()
}

func (dmf *DbMessageFields) ValidateDbMessageFields() error {
	if err := validate.Get().Struct(dmf); err != nil {
		return err
	}
	return nil
}

// DbMessageFieldsSlice is a parsable slice of DbMessageFields.
type DbMessageFieldsSlice []*DbMessageFields

func (dmf DbMessageFieldsSlice) config(cfg config) {
	for _i := range dmf {
		dmf[_i].config = cfg
	}
}