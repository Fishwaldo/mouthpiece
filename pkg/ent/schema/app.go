package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Fishwaldo/mouthpiece/pkg/validate"
)

// Apps holds the schema definition for the Apps entity.
type App struct {
	ent.Schema
}

// Fields of the Apps.
func (App) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").
			NotEmpty().
			Unique().
			Validate(validate.EntStringValidator("required,max=255,alphanum")).
			StructTag(`doc:"Application Name" pattern:"^[a-z0-9]+$" gorm:"unique;uniqueIndex"`),
		field.Enum("Status").
			Values("enabled", "disabled").
			StructTag(`doc:"Status of Application" enum:"Enabled,Disabled" default:"Enabled"`),
		field.String("Description").
			NotEmpty().
			Validate(validate.EntStringValidator("required,max=255")).
			StructTag(`doc:"Description of Application"`),
		field.String("icon").
			Validate(validate.EntStringValidator("url")).
			StructTag(`doc:"Icon of Application"`),
		field.String("url").
			Validate(validate.EntStringValidator("url")).
			StructTag(`doc:"URL of Application"`),
	}
}

// Edges of the Apps.
func (App) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("messages", Message.Type),
		edge.To("filters", Filter.Type),
		edge.To("groups", Group.Type),
		edge.To("TransportRecipients", TransportRecipient.Type),
	}
}

func (App) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}