package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"

	"github.com/Fishwaldo/mouthpiece/pkg/validate"
)

// DbApp holds the schema definition for the DbApp entity.
type DbApp struct {
	ent.Schema
}

// Fields of the DbApp.
func (DbApp) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").
			NotEmpty().
			Unique().
			Validate(validate.EntStringValidator("required,max=255,alphanum")).
			StructTag(`doc:"Application Name"`),
		field.Enum("Status").
			GoType(interfaces.AppStatus(interfaces.Enabled)).
			StructTag(`doc:"Status of Application"`),
		field.String("Description").
			NotEmpty().
			Validate(validate.EntStringValidator("required,max=255")).
			StructTag(`doc:"Description of Application"`),
		field.String("icon").
			Optional().
			Validate(validate.EntStringValidator("url")).
			StructTag(`doc:"Icon of Application"`),
		field.String("url").
			Optional().
			Validate(validate.EntStringValidator("url")).
			StructTag(`doc:"URL of Application"`),
	}
}

// Edges of the Apps.
func (DbApp) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("messages", DbMessage.Type),
		edge.To("filters", DbFilter.Type),
		edge.To("groups", DbGroup.Type),
	}
}

func (DbApp) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}