package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Fishwaldo/mouthpiece/pkg/validate"
	"github.com/google/uuid"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field {
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.Text("Message").
			NotEmpty().
			Validate(validate.EntStringValidator("required")).
			StructTag(`doc:"Message to be Sent"`),
		field.Text("ShortMsg").
			Optional().
			Nillable().
			StructTag(`doc:"Short Message to be Sent" `),
		field.Text("Topic").
			Optional().
			Nillable().
			Validate(validate.EntStringValidator("alphanum")).
			StructTag(`doc:"Topic of Message"`),
		field.Int("Severity").
			Optional().
			Default(3).
			Min(1).
			Max(5).
			StructTag(`doc:"Severity of Message" minimum:"1" maximum:"5" default:"3" validate:"min=1,max=5"`),
		field.Time("Timestamp").
			Default(time.Now()).
			StructTag(`doc:"Timestamp of Message"`),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("vars", MsgVar.Type).
		Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
		edge.From("app", App.Type).
			Ref("messages").
			Required().
			Unique(),
	}
}

func (Message) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}