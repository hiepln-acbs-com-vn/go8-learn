package account

import (
	"github.com/gmhafiz/go8/internal/domain/ekyc"
	"time"
)

type Schema struct {
	ID          uint
	AccountName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	Ekycs       []*ekyc.Schema
}
