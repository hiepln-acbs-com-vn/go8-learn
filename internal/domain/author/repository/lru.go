package repository

import (
	"context"
	"strings"

	lru "github.com/hashicorp/golang-lru"

	"github.com/gmhafiz/go8/internal/domain/author"
	"github.com/gmhafiz/go8/internal/middleware"
)

type AuthorLRU struct {
	service Author
	lru     *lru.Cache
}

type AuthorLRUService interface {
	Read(ctx context.Context, id uint) (*author.Schema, error)
	Update(ctx context.Context, toAuthor *author.UpdateRequest) (*author.Schema, error)
	Delete(ctx context.Context, id uint) error
}

func NewLRUCache(service Author) *AuthorLRU {
	// Creates a cache with a size.
	// https://en.wikipedia.org/wiki/Cache_replacement_policies#Least_recently_used_(LRU)
	// Once cache is filled, the least recently used key is discarded to make
	// way for a new key.
	c, _ := lru.New(8)
	return &AuthorLRU{
		service: service,
		lru:     c,
	}
}

func (c *AuthorLRU) Read(ctx context.Context, id uint) (*author.Schema, error) {
	// (1) Picks up the key from context which is added in the handler layer.
	url, ok := ctx.Value(middleware.CacheURL).(string)
	if !ok {
		return c.service.Read(ctx, id)
	}

	if val, ok := c.lru.Get(url); ok {
		// (3) Simply cast the returned value.
		return val.(*author.Schema), nil
	}

	// (2) If key doesn't exist or pushed off the LRU queue, we call the
	//     repository layer.
	res, err := c.service.Read(ctx, id)
	if err != nil {
		return nil, err
	} else if res != nil {
		c.lru.Add(url, res)
	}

	return res, nil
}

func (c *AuthorLRU) Update(ctx context.Context, toAuthor *author.UpdateRequest) (*author.Schema, error) {
	c.invalidate(ctx)

	return c.service.Update(ctx, toAuthor)
}

func (c *AuthorLRU) Delete(ctx context.Context, id uint) error {
	c.invalidate(ctx)

	return c.service.Delete(ctx, id)
}

func (c *AuthorLRU) invalidate(ctx context.Context) {
	url, ok := ctx.Value(middleware.CacheURL).(string)
	if !ok {
		return
	}

	keys := c.lru.Keys()
	for _, key := range keys {
		if strings.HasPrefix(key.(string), url) {
			c.lru.Remove(key)
		}
	}
}
