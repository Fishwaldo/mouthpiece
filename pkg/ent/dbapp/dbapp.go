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

package dbapp

import (
	"fmt"

	"entgo.io/ent"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
)

const (
	// Label holds the string label denoting the dbapp type in the database.
	Label = "db_app"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTenantID holds the string denoting the tenant_id field in the database.
	FieldTenantID = "tenant_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldIcon holds the string denoting the icon field in the database.
	FieldIcon = "icon"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// EdgeTenant holds the string denoting the tenant edge name in mutations.
	EdgeTenant = "tenant"
	// EdgeMessages holds the string denoting the messages edge name in mutations.
	EdgeMessages = "messages"
	// EdgeFilters holds the string denoting the filters edge name in mutations.
	EdgeFilters = "filters"
	// EdgeGroups holds the string denoting the groups edge name in mutations.
	EdgeGroups = "groups"
	// Table holds the table name of the dbapp in the database.
	Table = "db_apps"
	// TenantTable is the table that holds the tenant relation/edge.
	TenantTable = "db_apps"
	// TenantInverseTable is the table name for the Tenant entity.
	// It exists in this package in order to avoid circular dependency with the "tenant" package.
	TenantInverseTable = "tenants"
	// TenantColumn is the table column denoting the tenant relation/edge.
	TenantColumn = "tenant_id"
	// MessagesTable is the table that holds the messages relation/edge.
	MessagesTable = "db_messages"
	// MessagesInverseTable is the table name for the DbMessage entity.
	// It exists in this package in order to avoid circular dependency with the "dbmessage" package.
	MessagesInverseTable = "db_messages"
	// MessagesColumn is the table column denoting the messages relation/edge.
	MessagesColumn = "db_app_messages"
	// FiltersTable is the table that holds the filters relation/edge. The primary key declared below.
	FiltersTable = "db_app_filters"
	// FiltersInverseTable is the table name for the DbFilter entity.
	// It exists in this package in order to avoid circular dependency with the "dbfilter" package.
	FiltersInverseTable = "db_filters"
	// GroupsTable is the table that holds the groups relation/edge. The primary key declared below.
	GroupsTable = "db_app_groups"
	// GroupsInverseTable is the table name for the DbGroup entity.
	// It exists in this package in order to avoid circular dependency with the "dbgroup" package.
	GroupsInverseTable = "db_groups"
)

// Columns holds all SQL columns for dbapp fields.
var Columns = []string{
	FieldID,
	FieldTenantID,
	FieldName,
	FieldStatus,
	FieldDescription,
	FieldIcon,
	FieldURL,
}

var (
	// FiltersPrimaryKey and FiltersColumn2 are the table columns denoting the
	// primary key for the filters relation (M2M).
	FiltersPrimaryKey = []string{"db_app_id", "db_filter_id"}
	// GroupsPrimaryKey and GroupsColumn2 are the table columns denoting the
	// primary key for the groups relation (M2M).
	GroupsPrimaryKey = []string{"db_app_id", "db_group_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/Fishwaldo/mouthpiece/pkg/ent/runtime"
//
var (
	Hooks  [3]ent.Hook
	Policy ent.Policy
	// NameValidator is a validator for the "Name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DescriptionValidator is a validator for the "Description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// IconValidator is a validator for the "icon" field. It is called by the builders before save.
	IconValidator func(string) error
	// URLValidator is a validator for the "url" field. It is called by the builders before save.
	URLValidator func(string) error
)

// StatusValidator is a validator for the "Status" field enum values. It is called by the builders before save.
func StatusValidator(_status interfaces.AppStatus) error {
	switch _status.String() {
	case "AppEnabled", "AppDisabled":
		return nil
	default:
		return fmt.Errorf("dbapp: invalid enum value for Status field: %q", _status)
	}
}
