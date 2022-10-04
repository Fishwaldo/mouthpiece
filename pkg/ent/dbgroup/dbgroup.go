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

package dbgroup

import (
	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the dbgroup type in the database.
	Label = "db_group"
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
	// EdgeTenant holds the string denoting the tenant edge name in mutations.
	EdgeTenant = "tenant"
	// EdgeTransportRecipients holds the string denoting the transportrecipients edge name in mutations.
	EdgeTransportRecipients = "TransportRecipients"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeFilters holds the string denoting the filters edge name in mutations.
	EdgeFilters = "filters"
	// EdgeApps holds the string denoting the apps edge name in mutations.
	EdgeApps = "apps"
	// Table holds the table name of the dbgroup in the database.
	Table = "db_groups"
	// TenantTable is the table that holds the tenant relation/edge.
	TenantTable = "db_groups"
	// TenantInverseTable is the table name for the Tenant entity.
	// It exists in this package in order to avoid circular dependency with the "tenant" package.
	TenantInverseTable = "tenants"
	// TenantColumn is the table column denoting the tenant relation/edge.
	TenantColumn = "tenant_id"
	// TransportRecipientsTable is the table that holds the TransportRecipients relation/edge.
	TransportRecipientsTable = "db_transport_recipients"
	// TransportRecipientsInverseTable is the table name for the DbTransportRecipients entity.
	// It exists in this package in order to avoid circular dependency with the "dbtransportrecipients" package.
	TransportRecipientsInverseTable = "db_transport_recipients"
	// TransportRecipientsColumn is the table column denoting the TransportRecipients relation/edge.
	TransportRecipientsColumn = "db_group_transport_recipients"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "db_user_groups"
	// UsersInverseTable is the table name for the DbUser entity.
	// It exists in this package in order to avoid circular dependency with the "dbuser" package.
	UsersInverseTable = "db_users"
	// FiltersTable is the table that holds the filters relation/edge. The primary key declared below.
	FiltersTable = "db_filter_groups"
	// FiltersInverseTable is the table name for the DbFilter entity.
	// It exists in this package in order to avoid circular dependency with the "dbfilter" package.
	FiltersInverseTable = "db_filters"
	// AppsTable is the table that holds the apps relation/edge. The primary key declared below.
	AppsTable = "db_app_groups"
	// AppsInverseTable is the table name for the DbApp entity.
	// It exists in this package in order to avoid circular dependency with the "dbapp" package.
	AppsInverseTable = "db_apps"
)

// Columns holds all SQL columns for dbgroup fields.
var Columns = []string{
	FieldID,
	FieldTenantID,
	FieldAppData,
	FieldName,
	FieldDescription,
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"db_user_id", "db_group_id"}
	// FiltersPrimaryKey and FiltersColumn2 are the table columns denoting the
	// primary key for the filters relation (M2M).
	FiltersPrimaryKey = []string{"db_filter_id", "db_group_id"}
	// AppsPrimaryKey and AppsColumn2 are the table columns denoting the
	// primary key for the apps relation (M2M).
	AppsPrimaryKey = []string{"db_app_id", "db_group_id"}
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
)
