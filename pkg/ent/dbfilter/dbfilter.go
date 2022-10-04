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

package dbfilter

import (
	"fmt"

	"entgo.io/ent"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
)

const (
	// Label holds the string label denoting the dbfilter type in the database.
	Label = "db_filter"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTenantID holds the string denoting the tenant_id field in the database.
	FieldTenantID = "tenant_id"
	// FieldAppData holds the string denoting the appdata field in the database.
	FieldAppData = "app_data"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldEnabled holds the string denoting the enabled field in the database.
	FieldEnabled = "enabled"
	// FieldFilterImpl holds the string denoting the filterimpl field in the database.
	FieldFilterImpl = "filter_impl"
	// FieldConfig holds the string denoting the config field in the database.
	FieldConfig = "config"
	// EdgeTenant holds the string denoting the tenant edge name in mutations.
	EdgeTenant = "tenant"
	// EdgeGroups holds the string denoting the groups edge name in mutations.
	EdgeGroups = "groups"
	// EdgeApp holds the string denoting the app edge name in mutations.
	EdgeApp = "app"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the dbfilter in the database.
	Table = "db_filters"
	// TenantTable is the table that holds the tenant relation/edge.
	TenantTable = "db_filters"
	// TenantInverseTable is the table name for the Tenant entity.
	// It exists in this package in order to avoid circular dependency with the "tenant" package.
	TenantInverseTable = "tenants"
	// TenantColumn is the table column denoting the tenant relation/edge.
	TenantColumn = "tenant_id"
	// GroupsTable is the table that holds the groups relation/edge. The primary key declared below.
	GroupsTable = "db_filter_groups"
	// GroupsInverseTable is the table name for the DbGroup entity.
	// It exists in this package in order to avoid circular dependency with the "dbgroup" package.
	GroupsInverseTable = "db_groups"
	// AppTable is the table that holds the app relation/edge. The primary key declared below.
	AppTable = "db_app_filters"
	// AppInverseTable is the table name for the DbApp entity.
	// It exists in this package in order to avoid circular dependency with the "dbapp" package.
	AppInverseTable = "db_apps"
	// UserTable is the table that holds the user relation/edge. The primary key declared below.
	UserTable = "db_user_filters"
	// UserInverseTable is the table name for the DbUser entity.
	// It exists in this package in order to avoid circular dependency with the "dbuser" package.
	UserInverseTable = "db_users"
)

// Columns holds all SQL columns for dbfilter fields.
var Columns = []string{
	FieldID,
	FieldTenantID,
	FieldAppData,
	FieldName,
	FieldDescription,
	FieldType,
	FieldEnabled,
	FieldFilterImpl,
	FieldConfig,
}

var (
	// GroupsPrimaryKey and GroupsColumn2 are the table columns denoting the
	// primary key for the groups relation (M2M).
	GroupsPrimaryKey = []string{"db_filter_id", "db_group_id"}
	// AppPrimaryKey and AppColumn2 are the table columns denoting the
	// primary key for the app relation (M2M).
	AppPrimaryKey = []string{"db_app_id", "db_filter_id"}
	// UserPrimaryKey and UserColumn2 are the table columns denoting the
	// primary key for the user relation (M2M).
	UserPrimaryKey = []string{"db_user_id", "db_filter_id"}
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
	// DefaultEnabled holds the default value on creation for the "Enabled" field.
	DefaultEnabled bool
	// FilterImplValidator is a validator for the "FilterImpl" field. It is called by the builders before save.
	FilterImplValidator func(string) error
)

// TypeValidator is a validator for the "Type" field enum values. It is called by the builders before save.
func TypeValidator(_type interfaces.FilterType) error {
	switch _type.String() {
	case "InvalidFilter", "AppFilter", "UserFilter", "TransportFilter":
		return nil
	default:
		return fmt.Errorf("dbfilter: invalid enum value for Type field: %q", _type)
	}
}
