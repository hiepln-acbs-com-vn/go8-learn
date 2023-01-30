package usecase

import (
	"context"
	"github.com/gmhafiz/go8/config"
	"github.com/gmhafiz/go8/internal/domain/account"
	"github.com/gmhafiz/go8/internal/domain/account/repository"
)

type AccountUseCase struct {
	repo repository.Account

	cfg config.Cache
}

type Account interface {
	Create(ctx context.Context, a *account.CreateRequest) (*account.Schema, error)
}

func New(c config.Cache, repo repository.Account) *AccountUseCase {
	return &AccountUseCase{
		cfg:  c,
		repo: repo,
	}
}

func (u *AccountUseCase) Create(ctx context.Context, r *account.CreateRequest) (*account.Schema, error) {
	return u.repo.Create(ctx, r)
}
