package ekyc

import (
	"entgo.io/ent/dialect/sql"
	"time"
)

type Schema struct {
	ID        int          `db:"id"`
	EkycName  string       `db:"ekyc_name"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at" swaggertype:"string"`
}
