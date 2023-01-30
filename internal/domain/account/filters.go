package account

import (
	"github.com/gmhafiz/go8/internal/utility/filter"
	"net/url"
)

type Filter struct {
	Base filter.Filter

	AccountName string `json:"account_name"`
}

func Filters(queries url.Values) *Filter {
	f := filter.New(queries)
	if queries.Has("first_name") || queries.Has("last_name") {
		f.Search = true
	}
	return &Filter{
		Base: *f,

		AccountName: queries.Get("first_name"),
	}
}
