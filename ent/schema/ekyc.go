package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Ekyc struct {
	ent.Schema
}

// Fields of the Ekycs.
func (Ekyc) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id"),
		field.String("ekyc_name"),
		field.Time("created_at").Optional().StructTag(`json:"-"`),
		field.Time("updated_at").Optional().StructTag(`json:"-"`),
		field.Time("deleted_at").Optional().Nillable().StructTag(`json:"-"`),
	}
}

// Edges of the Ekyc.
func (Ekyc) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("accounts", Account.Type),
	}
}
