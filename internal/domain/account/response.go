package account

import (
	"github.com/gmhafiz/go8/internal/domain/ekyc"
)

type GetResponse struct {
	ID          uint           `json:"id"`
	AccountName string         `json:"account_name"`
	Ekycs       []*ekyc.Schema `json:"ekycs"`
}

func Resource(a *Schema) *GetResponse {
	if a == nil {
		return &GetResponse{}
	}

	return &GetResponse{
		ID:          a.ID,
		AccountName: a.AccountName,
		Ekycs:       a.Ekycs,
	}
}

func Resources(accounts []*Schema) []*GetResponse {
	if len(accounts) == 0 {
		return make([]*GetResponse, 0)
	}

	var resources []*GetResponse
	for _, a := range accounts {
		res := Resource(a)
		resources = append(resources, res)
	}
	return resources
}
