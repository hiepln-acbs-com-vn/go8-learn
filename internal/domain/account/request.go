package account

import (
	"encoding/json"
	"io"
)

type CreateRequest struct {
	AccountName string `json:"account_name" validate:"required"`
	Ekycs       []Ekyc `json:"ekycs"`
}

type Ekyc struct {
	EkycID   int    `json:"id"`
	EkycName string `json:"ekyc_name" validate:"required"`
}

func (r *CreateRequest) Bind(body io.ReadCloser) error {
	return json.NewDecoder(body).Decode(r)
}
