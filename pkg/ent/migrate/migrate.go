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
	"context"
	"fmt"
	"io"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
)

var (
	// WithGlobalUniqueID sets the universal ids options to the migration.
	// If this option is enabled, ent migration will allocate a 1<<32 range
	// for the ids of each entity (table).
	// Note that this option cannot be applied on tables that already exist.
	WithGlobalUniqueID = schema.WithGlobalUniqueID
	// WithDropColumn sets the drop column option to the migration.
	// If this option is enabled, ent migration will drop old columns
	// that were used for both fields and edges. This defaults to false.
	WithDropColumn = schema.WithDropColumn
	// WithDropIndex sets the drop index option to the migration.
	// If this option is enabled, ent migration will drop old indexes
	// that were defined in the schema. This defaults to false.
	// Note that unique constraints are defined using `UNIQUE INDEX`,
	// and therefore, it's recommended to enable this option to get more
	// flexibility in the schema changes.
	WithDropIndex = schema.WithDropIndex
	// WithForeignKeys enables creating foreign-key in schema DDL. This defaults to true.
	WithForeignKeys = schema.WithForeignKeys
)

// Schema is the API for creating, migrating and dropping a schema.
type Schema struct {
	drv dialect.Driver
}

// NewSchema creates a new schema client.
func NewSchema(drv dialect.Driver) *Schema { return &Schema{drv: drv} }

// Create creates all schema resources.
func (s *Schema) Create(ctx context.Context, opts ...schema.MigrateOption) error {
	return Create(ctx, s, Tables, opts...)
}

// Create creates all table resources using the given schema driver.
func Create(ctx context.Context, s *Schema, tables []*schema.Table, opts ...schema.MigrateOption) error {
	migrate, err := schema.NewMigrate(s.drv, opts...)
	if err != nil {
		return fmt.Errorf("ent/migrate: %w", err)
	}
	return migrate.Create(ctx, tables...)
}

// Diff compares the state read from a database connection or migration directory with
// the state defined by the Ent schema. Changes will be written to new migration files.
func Diff(ctx context.Context, url string, opts ...schema.MigrateOption) error {
	return NamedDiff(ctx, url, "changes", opts...)
}

// NamedDiff compares the state read from a database connection or migration directory with
// the state defined by the Ent schema. Changes will be written to new named migration files.
func NamedDiff(ctx context.Context, url, name string, opts ...schema.MigrateOption) error {
	return schema.Diff(ctx, url, name, Tables, opts...)
}

// Diff creates a migration file containing the statements to resolve the diff
// between the Ent schema and the connected database.
func (s *Schema) Diff(ctx context.Context, opts ...schema.MigrateOption) error {
	migrate, err := schema.NewMigrate(s.drv, opts...)
	if err != nil {
		return fmt.Errorf("ent/migrate: %w", err)
	}
	return migrate.Diff(ctx, Tables...)
}

// NamedDiff creates a named migration file containing the statements to resolve the diff
// between the Ent schema and the connected database.
func (s *Schema) NamedDiff(ctx context.Context, name string, opts ...schema.MigrateOption) error {
	migrate, err := schema.NewMigrate(s.drv, opts...)
	if err != nil {
		return fmt.Errorf("ent/migrate: %w", err)
	}
	return migrate.NamedDiff(ctx, name, Tables...)
}

// WriteTo writes the schema changes to w instead of running them against the database.
//
// 	if err := client.Schema.WriteTo(context.Background(), os.Stdout); err != nil {
//		log.Fatal(err)
// 	}
//
func (s *Schema) WriteTo(ctx context.Context, w io.Writer, opts ...schema.MigrateOption) error {
	return Create(ctx, &Schema{drv: &schema.WriteDriver{Writer: w, Driver: s.drv}}, Tables, opts...)
}