package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/gmhafiz/go8/ent/gen"
	"github.com/gmhafiz/go8/internal/domain/account"
	"github.com/gmhafiz/go8/internal/domain/ekyc"
)

type repository struct {
	ent *gen.Client
}

type Account interface {
	Create(ctx context.Context, a *account.CreateRequest) (*account.Schema, error)
}

func New(ent *gen.Client) *repository {
	return &repository{
		ent: ent,
	}
}

func (r *repository) Create(ctx context.Context, request *account.CreateRequest) (*account.Schema, error) {
	if request == nil {
		return nil, errors.New("request cannot be nil")
	}
	bulk := make([]*gen.EkycCreate, len(request.Ekycs))
	for i, b := range request.Ekycs {
		bulk[i] = r.ent.Ekyc.Create().
			SetEkycName(b.EkycName)
	}
	ekycs, err := r.ent.Ekyc.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("author.repository.Create bulk books: %w", err)
	}

	create, err := r.ent.Account.Create().
		SetAccountName(request.AccountName).
		AddEkycs(ekycs...).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("author.repository.Create: %w", err)
	}

	// Both created_at and updated_at are created database-side instead of ent.
	// So ent does not return both.
	created, err := r.ent.Account.Get(ctx, create.ID)
	if err != nil {
		return nil, fmt.Errorf("author not found: %w", err)
	}

	var b []*ekyc.Schema
	for _, i := range ekycs {

		b = append(b, &ekyc.Schema{
			ID:       int(i.ID),
			EkycName: i.EkycName,
		})
	}

	resp := &account.Schema{
		ID:          created.ID,
		AccountName: created.AccountName,
		CreatedAt:   created.CreatedAt,
		UpdatedAt:   created.UpdatedAt,
		DeletedAt:   created.DeletedAt,
		Ekycs:       b,
	}

	return resp, nil
}
