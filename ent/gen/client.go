// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/gmhafiz/go8/ent/gen/migrate"

	"github.com/gmhafiz/go8/ent/gen/account"
	"github.com/gmhafiz/go8/ent/gen/author"
	"github.com/gmhafiz/go8/ent/gen/book"
	"github.com/gmhafiz/go8/ent/gen/ekyc"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Account is the client for interacting with the Account builders.
	Account *AccountClient
	// Author is the client for interacting with the Author builders.
	Author *AuthorClient
	// Book is the client for interacting with the Book builders.
	Book *BookClient
	// Ekyc is the client for interacting with the Ekyc builders.
	Ekyc *EkycClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Account = NewAccountClient(c.config)
	c.Author = NewAuthorClient(c.config)
	c.Book = NewBookClient(c.config)
	c.Ekyc = NewEkycClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("gen: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("gen: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Account: NewAccountClient(cfg),
		Author:  NewAuthorClient(cfg),
		Book:    NewBookClient(cfg),
		Ekyc:    NewEkycClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Account: NewAccountClient(cfg),
		Author:  NewAuthorClient(cfg),
		Book:    NewBookClient(cfg),
		Ekyc:    NewEkycClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Account.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Account.Use(hooks...)
	c.Author.Use(hooks...)
	c.Book.Use(hooks...)
	c.Ekyc.Use(hooks...)
}

// AccountClient is a client for the Account schema.
type AccountClient struct {
	config
}

// NewAccountClient returns a client for the Account from the given config.
func NewAccountClient(c config) *AccountClient {
	return &AccountClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `account.Hooks(f(g(h())))`.
func (c *AccountClient) Use(hooks ...Hook) {
	c.hooks.Account = append(c.hooks.Account, hooks...)
}

// Create returns a builder for creating a Account entity.
func (c *AccountClient) Create() *AccountCreate {
	mutation := newAccountMutation(c.config, OpCreate)
	return &AccountCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Account entities.
func (c *AccountClient) CreateBulk(builders ...*AccountCreate) *AccountCreateBulk {
	return &AccountCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Account.
func (c *AccountClient) Update() *AccountUpdate {
	mutation := newAccountMutation(c.config, OpUpdate)
	return &AccountUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AccountClient) UpdateOne(a *Account) *AccountUpdateOne {
	mutation := newAccountMutation(c.config, OpUpdateOne, withAccount(a))
	return &AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AccountClient) UpdateOneID(id uint) *AccountUpdateOne {
	mutation := newAccountMutation(c.config, OpUpdateOne, withAccountID(id))
	return &AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Account.
func (c *AccountClient) Delete() *AccountDelete {
	mutation := newAccountMutation(c.config, OpDelete)
	return &AccountDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AccountClient) DeleteOne(a *Account) *AccountDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AccountClient) DeleteOneID(id uint) *AccountDeleteOne {
	builder := c.Delete().Where(account.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AccountDeleteOne{builder}
}

// Query returns a query builder for Account.
func (c *AccountClient) Query() *AccountQuery {
	return &AccountQuery{
		config: c.config,
	}
}

// Get returns a Account entity by its id.
func (c *AccountClient) Get(ctx context.Context, id uint) (*Account, error) {
	return c.Query().Where(account.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AccountClient) GetX(ctx context.Context, id uint) *Account {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryEkycs queries the ekycs edge of a Account.
func (c *AccountClient) QueryEkycs(a *Account) *EkycQuery {
	query := &EkycQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, id),
			sqlgraph.To(ekyc.Table, ekyc.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, account.EkycsTable, account.EkycsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AccountClient) Hooks() []Hook {
	return c.hooks.Account
}

// AuthorClient is a client for the Author schema.
type AuthorClient struct {
	config
}

// NewAuthorClient returns a client for the Author from the given config.
func NewAuthorClient(c config) *AuthorClient {
	return &AuthorClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `author.Hooks(f(g(h())))`.
func (c *AuthorClient) Use(hooks ...Hook) {
	c.hooks.Author = append(c.hooks.Author, hooks...)
}

// Create returns a builder for creating a Author entity.
func (c *AuthorClient) Create() *AuthorCreate {
	mutation := newAuthorMutation(c.config, OpCreate)
	return &AuthorCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Author entities.
func (c *AuthorClient) CreateBulk(builders ...*AuthorCreate) *AuthorCreateBulk {
	return &AuthorCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Author.
func (c *AuthorClient) Update() *AuthorUpdate {
	mutation := newAuthorMutation(c.config, OpUpdate)
	return &AuthorUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AuthorClient) UpdateOne(a *Author) *AuthorUpdateOne {
	mutation := newAuthorMutation(c.config, OpUpdateOne, withAuthor(a))
	return &AuthorUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AuthorClient) UpdateOneID(id uint) *AuthorUpdateOne {
	mutation := newAuthorMutation(c.config, OpUpdateOne, withAuthorID(id))
	return &AuthorUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Author.
func (c *AuthorClient) Delete() *AuthorDelete {
	mutation := newAuthorMutation(c.config, OpDelete)
	return &AuthorDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AuthorClient) DeleteOne(a *Author) *AuthorDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AuthorClient) DeleteOneID(id uint) *AuthorDeleteOne {
	builder := c.Delete().Where(author.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AuthorDeleteOne{builder}
}

// Query returns a query builder for Author.
func (c *AuthorClient) Query() *AuthorQuery {
	return &AuthorQuery{
		config: c.config,
	}
}

// Get returns a Author entity by its id.
func (c *AuthorClient) Get(ctx context.Context, id uint) (*Author, error) {
	return c.Query().Where(author.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AuthorClient) GetX(ctx context.Context, id uint) *Author {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryBooks queries the books edge of a Author.
func (c *AuthorClient) QueryBooks(a *Author) *BookQuery {
	query := &BookQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(author.Table, author.FieldID, id),
			sqlgraph.To(book.Table, book.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, author.BooksTable, author.BooksPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AuthorClient) Hooks() []Hook {
	return c.hooks.Author
}

// BookClient is a client for the Book schema.
type BookClient struct {
	config
}

// NewBookClient returns a client for the Book from the given config.
func NewBookClient(c config) *BookClient {
	return &BookClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `book.Hooks(f(g(h())))`.
func (c *BookClient) Use(hooks ...Hook) {
	c.hooks.Book = append(c.hooks.Book, hooks...)
}

// Create returns a builder for creating a Book entity.
func (c *BookClient) Create() *BookCreate {
	mutation := newBookMutation(c.config, OpCreate)
	return &BookCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Book entities.
func (c *BookClient) CreateBulk(builders ...*BookCreate) *BookCreateBulk {
	return &BookCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Book.
func (c *BookClient) Update() *BookUpdate {
	mutation := newBookMutation(c.config, OpUpdate)
	return &BookUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BookClient) UpdateOne(b *Book) *BookUpdateOne {
	mutation := newBookMutation(c.config, OpUpdateOne, withBook(b))
	return &BookUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BookClient) UpdateOneID(id uint) *BookUpdateOne {
	mutation := newBookMutation(c.config, OpUpdateOne, withBookID(id))
	return &BookUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Book.
func (c *BookClient) Delete() *BookDelete {
	mutation := newBookMutation(c.config, OpDelete)
	return &BookDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BookClient) DeleteOne(b *Book) *BookDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *BookClient) DeleteOneID(id uint) *BookDeleteOne {
	builder := c.Delete().Where(book.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BookDeleteOne{builder}
}

// Query returns a query builder for Book.
func (c *BookClient) Query() *BookQuery {
	return &BookQuery{
		config: c.config,
	}
}

// Get returns a Book entity by its id.
func (c *BookClient) Get(ctx context.Context, id uint) (*Book, error) {
	return c.Query().Where(book.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BookClient) GetX(ctx context.Context, id uint) *Book {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAuthors queries the authors edge of a Book.
func (c *BookClient) QueryAuthors(b *Book) *AuthorQuery {
	query := &AuthorQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(book.Table, book.FieldID, id),
			sqlgraph.To(author.Table, author.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, book.AuthorsTable, book.AuthorsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BookClient) Hooks() []Hook {
	return c.hooks.Book
}

// EkycClient is a client for the Ekyc schema.
type EkycClient struct {
	config
}

// NewEkycClient returns a client for the Ekyc from the given config.
func NewEkycClient(c config) *EkycClient {
	return &EkycClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `ekyc.Hooks(f(g(h())))`.
func (c *EkycClient) Use(hooks ...Hook) {
	c.hooks.Ekyc = append(c.hooks.Ekyc, hooks...)
}

// Create returns a builder for creating a Ekyc entity.
func (c *EkycClient) Create() *EkycCreate {
	mutation := newEkycMutation(c.config, OpCreate)
	return &EkycCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Ekyc entities.
func (c *EkycClient) CreateBulk(builders ...*EkycCreate) *EkycCreateBulk {
	return &EkycCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Ekyc.
func (c *EkycClient) Update() *EkycUpdate {
	mutation := newEkycMutation(c.config, OpUpdate)
	return &EkycUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EkycClient) UpdateOne(e *Ekyc) *EkycUpdateOne {
	mutation := newEkycMutation(c.config, OpUpdateOne, withEkyc(e))
	return &EkycUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EkycClient) UpdateOneID(id uint) *EkycUpdateOne {
	mutation := newEkycMutation(c.config, OpUpdateOne, withEkycID(id))
	return &EkycUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Ekyc.
func (c *EkycClient) Delete() *EkycDelete {
	mutation := newEkycMutation(c.config, OpDelete)
	return &EkycDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *EkycClient) DeleteOne(e *Ekyc) *EkycDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *EkycClient) DeleteOneID(id uint) *EkycDeleteOne {
	builder := c.Delete().Where(ekyc.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EkycDeleteOne{builder}
}

// Query returns a query builder for Ekyc.
func (c *EkycClient) Query() *EkycQuery {
	return &EkycQuery{
		config: c.config,
	}
}

// Get returns a Ekyc entity by its id.
func (c *EkycClient) Get(ctx context.Context, id uint) (*Ekyc, error) {
	return c.Query().Where(ekyc.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EkycClient) GetX(ctx context.Context, id uint) *Ekyc {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAccounts queries the accounts edge of a Ekyc.
func (c *EkycClient) QueryAccounts(e *Ekyc) *AccountQuery {
	query := &AccountQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(ekyc.Table, ekyc.FieldID, id),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ekyc.AccountsTable, ekyc.AccountsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *EkycClient) Hooks() []Hook {
	return c.hooks.Ekyc
}
