package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Account struct {
	ent.Schema
}

// Fields of the Accounts.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id"),
		field.String("account_name"),
		field.Time("created_at").Optional().StructTag(`json:"-"`),
		field.Time("updated_at").Optional().StructTag(`json:"-"`),
		field.Time("deleted_at").Optional().Nillable().StructTag(`json:"-"`),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("ekycs", Ekyc.Type).Ref("accounts"),
	}
}
