// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint, Increment: true},
		{Name: "account_name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
	}
	// AuthorsColumns holds the columns for the "authors" table.
	AuthorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint, Increment: true},
		{Name: "first_name", Type: field.TypeString},
		{Name: "middle_name", Type: field.TypeString, Nullable: true},
		{Name: "last_name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// AuthorsTable holds the schema information for the "authors" table.
	AuthorsTable = &schema.Table{
		Name:       "authors",
		Columns:    AuthorsColumns,
		PrimaryKey: []*schema.Column{AuthorsColumns[0]},
	}
	// BooksColumns holds the columns for the "books" table.
	BooksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "published_date", Type: field.TypeTime},
		{Name: "image_url", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// BooksTable holds the schema information for the "books" table.
	BooksTable = &schema.Table{
		Name:       "books",
		Columns:    BooksColumns,
		PrimaryKey: []*schema.Column{BooksColumns[0]},
	}
	// EkycsColumns holds the columns for the "ekycs" table.
	EkycsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint, Increment: true},
		{Name: "ekyc_name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// EkycsTable holds the schema information for the "ekycs" table.
	EkycsTable = &schema.Table{
		Name:       "ekycs",
		Columns:    EkycsColumns,
		PrimaryKey: []*schema.Column{EkycsColumns[0]},
	}
	// BookAuthorsColumns holds the columns for the "book_authors" table.
	BookAuthorsColumns = []*schema.Column{
		{Name: "book_id", Type: field.TypeUint},
		{Name: "author_id", Type: field.TypeUint},
	}
	// BookAuthorsTable holds the schema information for the "book_authors" table.
	BookAuthorsTable = &schema.Table{
		Name:       "book_authors",
		Columns:    BookAuthorsColumns,
		PrimaryKey: []*schema.Column{BookAuthorsColumns[0], BookAuthorsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "book_authors_book_id",
				Columns:    []*schema.Column{BookAuthorsColumns[0]},
				RefColumns: []*schema.Column{BooksColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "book_authors_author_id",
				Columns:    []*schema.Column{BookAuthorsColumns[1]},
				RefColumns: []*schema.Column{AuthorsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// EkycAccountsColumns holds the columns for the "ekyc_accounts" table.
	EkycAccountsColumns = []*schema.Column{
		{Name: "ekyc_id", Type: field.TypeUint},
		{Name: "account_id", Type: field.TypeUint},
	}
	// EkycAccountsTable holds the schema information for the "ekyc_accounts" table.
	EkycAccountsTable = &schema.Table{
		Name:       "ekyc_accounts",
		Columns:    EkycAccountsColumns,
		PrimaryKey: []*schema.Column{EkycAccountsColumns[0], EkycAccountsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "ekyc_accounts_ekyc_id",
				Columns:    []*schema.Column{EkycAccountsColumns[0]},
				RefColumns: []*schema.Column{EkycsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "ekyc_accounts_account_id",
				Columns:    []*schema.Column{EkycAccountsColumns[1]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountsTable,
		AuthorsTable,
		BooksTable,
		EkycsTable,
		BookAuthorsTable,
		EkycAccountsTable,
	}
)

func init() {
	BookAuthorsTable.ForeignKeys[0].RefTable = BooksTable
	BookAuthorsTable.ForeignKeys[1].RefTable = AuthorsTable
	EkycAccountsTable.ForeignKeys[0].RefTable = EkycsTable
	EkycAccountsTable.ForeignKeys[1].RefTable = AccountsTable
}
