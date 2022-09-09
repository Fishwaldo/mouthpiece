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

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AppsColumns holds the columns for the "apps" table.
	AppsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"enabled", "disabled"}},
		{Name: "description", Type: field.TypeString},
		{Name: "icon", Type: field.TypeString},
		{Name: "url", Type: field.TypeString},
		{Name: "tenant_id", Type: field.TypeInt},
	}
	// AppsTable holds the schema information for the "apps" table.
	AppsTable = &schema.Table{
		Name:       "apps",
		Columns:    AppsColumns,
		PrimaryKey: []*schema.Column{AppsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "apps_tenants_tenant",
				Columns:    []*schema.Column{AppsColumns[6]},
				RefColumns: []*schema.Column{TenantsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// FiltersColumns holds the columns for the "filters" table.
	FiltersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"InvalidFilter", "AppFilter", "UserFilter", "TransportFilter"}, Default: "InvalidFilter"},
		{Name: "enabled", Type: field.TypeBool, Default: true},
		{Name: "filter_impl", Type: field.TypeString},
		{Name: "tenant_id", Type: field.TypeInt},
	}
	// FiltersTable holds the schema information for the "filters" table.
	FiltersTable = &schema.Table{
		Name:       "filters",
		Columns:    FiltersColumns,
		PrimaryKey: []*schema.Column{FiltersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "filters_tenants_tenant",
				Columns:    []*schema.Column{FiltersColumns[6]},
				RefColumns: []*schema.Column{TenantsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// FilterConfigsColumns holds the columns for the "filter_configs" table.
	FilterConfigsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Size: 2147483647},
		{Name: "value", Type: field.TypeString, Size: 2147483647},
		{Name: "filter_config", Type: field.TypeInt},
		{Name: "tenant_id", Type: field.TypeInt},
	}
	// FilterConfigsTable holds the schema information for the "filter_configs" table.
	FilterConfigsTable = &schema.Table{
		Name:       "filter_configs",
		Columns:    FilterConfigsColumns,
		PrimaryKey: []*schema.Column{FilterConfigsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "filter_configs_filters_config",
				Columns:    []*schema.Column{FilterConfigsColumns[3]},
				RefColumns: []*schema.Column{FiltersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "filter_configs_tenants_tenant",
				Columns:    []*schema.Column{FilterConfigsColumns[4]},
				RefColumns: []*schema.Column{TenantsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString},
		{Name: "tenant_id", Type: field.TypeInt},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "groups_tenants_tenant",
				Columns:    []*schema.Column{GroupsColumns[3]},
				RefColumns: []*schema.Column{TenantsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// MessagesColumns holds the columns for the "messages" table.
	MessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "message", Type: field.TypeString, Size: 2147483647},
		{Name: "short_msg", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "topic", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "severity", Type: field.TypeInt, Nullable: true, Default: 3},
		{Name: "timestamp", Type: field.TypeTime},
		{Name: "app_messages", Type: field.TypeInt},
		{Name: "tenant_id", Type: field.TypeInt},
	}
	// MessagesTable holds the schema information for the "messages" table.
	MessagesTable = &schema.Table{
		Name:       "messages",
		Columns:    MessagesColumns,
		PrimaryKey: []*schema.Column{MessagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "messages_apps_messages",
				Columns:    []*schema.Column{MessagesColumns[6]},
				RefColumns: []*schema.Column{AppsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "messages_tenants_tenant",
				Columns:    []*schema.Column{MessagesColumns[7]},
				RefColumns: []*schema.Column{TenantsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// MsgVarsColumns holds the columns for the "msg_vars" table.
	MsgVarsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 2147483647},
		{Name: "value", Type: field.TypeString, Size: 2147483647},
		{Name: "message_vars", Type: field.TypeUUID},
		{Name: "tenant_id", Type: field.TypeInt},
	}
	// MsgVarsTable holds the schema information for the "msg_vars" table.
	MsgVarsTable = &schema.Table{
		Name:       "msg_vars",
		Columns:    MsgVarsColumns,
		PrimaryKey: []*schema.Column{MsgVarsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "msg_vars_messages_vars",
				Columns:    []*schema.Column{MsgVarsColumns[3]},
				RefColumns: []*schema.Column{MessagesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "msg_vars_tenants_tenant",
				Columns:    []*schema.Column{MsgVarsColumns[4]},
				RefColumns: []*schema.Column{TenantsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TenantsColumns holds the columns for the "tenants" table.
	TenantsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// TenantsTable holds the schema information for the "tenants" table.
	TenantsTable = &schema.Table{
		Name:       "tenants",
		Columns:    TenantsColumns,
		PrimaryKey: []*schema.Column{TenantsColumns[0]},
	}
	// TransportInstancesColumns holds the columns for the "transport_instances" table.
	TransportInstancesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString},
		{Name: "tenant_id", Type: field.TypeInt},
	}
	// TransportInstancesTable holds the schema information for the "transport_instances" table.
	TransportInstancesTable = &schema.Table{
		Name:       "transport_instances",
		Columns:    TransportInstancesColumns,
		PrimaryKey: []*schema.Column{TransportInstancesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "transport_instances_tenants_tenant",
				Columns:    []*schema.Column{TransportInstancesColumns[3]},
				RefColumns: []*schema.Column{TenantsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TransportRecipientsColumns holds the columns for the "transport_recipients" table.
	TransportRecipientsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString},
		{Name: "transport_instance_transport_recipients", Type: field.TypeInt},
		{Name: "tenant_id", Type: field.TypeInt},
	}
	// TransportRecipientsTable holds the schema information for the "transport_recipients" table.
	TransportRecipientsTable = &schema.Table{
		Name:       "transport_recipients",
		Columns:    TransportRecipientsColumns,
		PrimaryKey: []*schema.Column{TransportRecipientsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "transport_recipients_transport_instances_TransportRecipients",
				Columns:    []*schema.Column{TransportRecipientsColumns[3]},
				RefColumns: []*schema.Column{TransportInstancesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "transport_recipients_tenants_tenant",
				Columns:    []*schema.Column{TransportRecipientsColumns[4]},
				RefColumns: []*schema.Column{TenantsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "tenant_id", Type: field.TypeInt},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_tenants_tenant",
				Columns:    []*schema.Column{UsersColumns[4]},
				RefColumns: []*schema.Column{TenantsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "user_email_tenant_id",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[1], UsersColumns[4]},
			},
		},
	}
	// UserMetaDataColumns holds the columns for the "user_meta_data" table.
	UserMetaDataColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Size: 2147483647},
		{Name: "value", Type: field.TypeString, Size: 2147483647},
		{Name: "user_metadata", Type: field.TypeInt},
		{Name: "tenant_id", Type: field.TypeInt},
	}
	// UserMetaDataTable holds the schema information for the "user_meta_data" table.
	UserMetaDataTable = &schema.Table{
		Name:       "user_meta_data",
		Columns:    UserMetaDataColumns,
		PrimaryKey: []*schema.Column{UserMetaDataColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_meta_data_users_metadata",
				Columns:    []*schema.Column{UserMetaDataColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_meta_data_tenants_tenant",
				Columns:    []*schema.Column{UserMetaDataColumns[4]},
				RefColumns: []*schema.Column{TenantsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// AppFiltersColumns holds the columns for the "app_filters" table.
	AppFiltersColumns = []*schema.Column{
		{Name: "app_id", Type: field.TypeInt},
		{Name: "filter_id", Type: field.TypeInt},
	}
	// AppFiltersTable holds the schema information for the "app_filters" table.
	AppFiltersTable = &schema.Table{
		Name:       "app_filters",
		Columns:    AppFiltersColumns,
		PrimaryKey: []*schema.Column{AppFiltersColumns[0], AppFiltersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "app_filters_app_id",
				Columns:    []*schema.Column{AppFiltersColumns[0]},
				RefColumns: []*schema.Column{AppsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "app_filters_filter_id",
				Columns:    []*schema.Column{AppFiltersColumns[1]},
				RefColumns: []*schema.Column{FiltersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// AppGroupsColumns holds the columns for the "app_groups" table.
	AppGroupsColumns = []*schema.Column{
		{Name: "app_id", Type: field.TypeInt},
		{Name: "group_id", Type: field.TypeInt},
	}
	// AppGroupsTable holds the schema information for the "app_groups" table.
	AppGroupsTable = &schema.Table{
		Name:       "app_groups",
		Columns:    AppGroupsColumns,
		PrimaryKey: []*schema.Column{AppGroupsColumns[0], AppGroupsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "app_groups_app_id",
				Columns:    []*schema.Column{AppGroupsColumns[0]},
				RefColumns: []*schema.Column{AppsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "app_groups_group_id",
				Columns:    []*schema.Column{AppGroupsColumns[1]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// AppTransportRecipientsColumns holds the columns for the "app_TransportRecipients" table.
	AppTransportRecipientsColumns = []*schema.Column{
		{Name: "app_id", Type: field.TypeInt},
		{Name: "transport_recipient_id", Type: field.TypeInt},
	}
	// AppTransportRecipientsTable holds the schema information for the "app_TransportRecipients" table.
	AppTransportRecipientsTable = &schema.Table{
		Name:       "app_TransportRecipients",
		Columns:    AppTransportRecipientsColumns,
		PrimaryKey: []*schema.Column{AppTransportRecipientsColumns[0], AppTransportRecipientsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "app_TransportRecipients_app_id",
				Columns:    []*schema.Column{AppTransportRecipientsColumns[0]},
				RefColumns: []*schema.Column{AppsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "app_TransportRecipients_transport_recipient_id",
				Columns:    []*schema.Column{AppTransportRecipientsColumns[1]},
				RefColumns: []*schema.Column{TransportRecipientsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// FilterGroupsColumns holds the columns for the "filter_groups" table.
	FilterGroupsColumns = []*schema.Column{
		{Name: "filter_id", Type: field.TypeInt},
		{Name: "group_id", Type: field.TypeInt},
	}
	// FilterGroupsTable holds the schema information for the "filter_groups" table.
	FilterGroupsTable = &schema.Table{
		Name:       "filter_groups",
		Columns:    FilterGroupsColumns,
		PrimaryKey: []*schema.Column{FilterGroupsColumns[0], FilterGroupsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "filter_groups_filter_id",
				Columns:    []*schema.Column{FilterGroupsColumns[0]},
				RefColumns: []*schema.Column{FiltersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "filter_groups_group_id",
				Columns:    []*schema.Column{FilterGroupsColumns[1]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// GroupTransportRecipientsColumns holds the columns for the "group_TransportRecipients" table.
	GroupTransportRecipientsColumns = []*schema.Column{
		{Name: "group_id", Type: field.TypeInt},
		{Name: "transport_recipient_id", Type: field.TypeInt},
	}
	// GroupTransportRecipientsTable holds the schema information for the "group_TransportRecipients" table.
	GroupTransportRecipientsTable = &schema.Table{
		Name:       "group_TransportRecipients",
		Columns:    GroupTransportRecipientsColumns,
		PrimaryKey: []*schema.Column{GroupTransportRecipientsColumns[0], GroupTransportRecipientsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "group_TransportRecipients_group_id",
				Columns:    []*schema.Column{GroupTransportRecipientsColumns[0]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "group_TransportRecipients_transport_recipient_id",
				Columns:    []*schema.Column{GroupTransportRecipientsColumns[1]},
				RefColumns: []*schema.Column{TransportRecipientsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserFiltersColumns holds the columns for the "user_filters" table.
	UserFiltersColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "filter_id", Type: field.TypeInt},
	}
	// UserFiltersTable holds the schema information for the "user_filters" table.
	UserFiltersTable = &schema.Table{
		Name:       "user_filters",
		Columns:    UserFiltersColumns,
		PrimaryKey: []*schema.Column{UserFiltersColumns[0], UserFiltersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_filters_user_id",
				Columns:    []*schema.Column{UserFiltersColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_filters_filter_id",
				Columns:    []*schema.Column{UserFiltersColumns[1]},
				RefColumns: []*schema.Column{FiltersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserGroupsColumns holds the columns for the "user_groups" table.
	UserGroupsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "group_id", Type: field.TypeInt},
	}
	// UserGroupsTable holds the schema information for the "user_groups" table.
	UserGroupsTable = &schema.Table{
		Name:       "user_groups",
		Columns:    UserGroupsColumns,
		PrimaryKey: []*schema.Column{UserGroupsColumns[0], UserGroupsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_groups_user_id",
				Columns:    []*schema.Column{UserGroupsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_groups_group_id",
				Columns:    []*schema.Column{UserGroupsColumns[1]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserTransportRecipientsColumns holds the columns for the "user_TransportRecipients" table.
	UserTransportRecipientsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "transport_recipient_id", Type: field.TypeInt},
	}
	// UserTransportRecipientsTable holds the schema information for the "user_TransportRecipients" table.
	UserTransportRecipientsTable = &schema.Table{
		Name:       "user_TransportRecipients",
		Columns:    UserTransportRecipientsColumns,
		PrimaryKey: []*schema.Column{UserTransportRecipientsColumns[0], UserTransportRecipientsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_TransportRecipients_user_id",
				Columns:    []*schema.Column{UserTransportRecipientsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_TransportRecipients_transport_recipient_id",
				Columns:    []*schema.Column{UserTransportRecipientsColumns[1]},
				RefColumns: []*schema.Column{TransportRecipientsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AppsTable,
		FiltersTable,
		FilterConfigsTable,
		GroupsTable,
		MessagesTable,
		MsgVarsTable,
		TenantsTable,
		TransportInstancesTable,
		TransportRecipientsTable,
		UsersTable,
		UserMetaDataTable,
		AppFiltersTable,
		AppGroupsTable,
		AppTransportRecipientsTable,
		FilterGroupsTable,
		GroupTransportRecipientsTable,
		UserFiltersTable,
		UserGroupsTable,
		UserTransportRecipientsTable,
	}
)

func init() {
	AppsTable.ForeignKeys[0].RefTable = TenantsTable
	FiltersTable.ForeignKeys[0].RefTable = TenantsTable
	FilterConfigsTable.ForeignKeys[0].RefTable = FiltersTable
	FilterConfigsTable.ForeignKeys[1].RefTable = TenantsTable
	GroupsTable.ForeignKeys[0].RefTable = TenantsTable
	MessagesTable.ForeignKeys[0].RefTable = AppsTable
	MessagesTable.ForeignKeys[1].RefTable = TenantsTable
	MsgVarsTable.ForeignKeys[0].RefTable = MessagesTable
	MsgVarsTable.ForeignKeys[1].RefTable = TenantsTable
	TransportInstancesTable.ForeignKeys[0].RefTable = TenantsTable
	TransportRecipientsTable.ForeignKeys[0].RefTable = TransportInstancesTable
	TransportRecipientsTable.ForeignKeys[1].RefTable = TenantsTable
	UsersTable.ForeignKeys[0].RefTable = TenantsTable
	UserMetaDataTable.ForeignKeys[0].RefTable = UsersTable
	UserMetaDataTable.ForeignKeys[1].RefTable = TenantsTable
	AppFiltersTable.ForeignKeys[0].RefTable = AppsTable
	AppFiltersTable.ForeignKeys[1].RefTable = FiltersTable
	AppGroupsTable.ForeignKeys[0].RefTable = AppsTable
	AppGroupsTable.ForeignKeys[1].RefTable = GroupsTable
	AppTransportRecipientsTable.ForeignKeys[0].RefTable = AppsTable
	AppTransportRecipientsTable.ForeignKeys[1].RefTable = TransportRecipientsTable
	FilterGroupsTable.ForeignKeys[0].RefTable = FiltersTable
	FilterGroupsTable.ForeignKeys[1].RefTable = GroupsTable
	GroupTransportRecipientsTable.ForeignKeys[0].RefTable = GroupsTable
	GroupTransportRecipientsTable.ForeignKeys[1].RefTable = TransportRecipientsTable
	UserFiltersTable.ForeignKeys[0].RefTable = UsersTable
	UserFiltersTable.ForeignKeys[1].RefTable = FiltersTable
	UserGroupsTable.ForeignKeys[0].RefTable = UsersTable
	UserGroupsTable.ForeignKeys[1].RefTable = GroupsTable
	UserTransportRecipientsTable.ForeignKeys[0].RefTable = UsersTable
	UserTransportRecipientsTable.ForeignKeys[1].RefTable = TransportRecipientsTable
}