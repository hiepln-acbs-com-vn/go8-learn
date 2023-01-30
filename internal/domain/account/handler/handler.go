package handler

import (
	"database/sql"
	"github.com/gmhafiz/go8/internal/domain/account"
	"github.com/gmhafiz/go8/internal/domain/account/usecase"
	"github.com/gmhafiz/go8/internal/utility/message"
	"github.com/gmhafiz/go8/internal/utility/respond"
	"github.com/gmhafiz/go8/internal/utility/validate"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

type Handler struct {
	useCase  usecase.Account
	validate *validator.Validate
}

func NewHandler(useCase usecase.Account, v *validator.Validate) *Handler {
	return &Handler{
		useCase:  useCase,
		validate: v,
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req account.CreateRequest
	err := req.Bind(r.Body)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	errs := validate.Validate(h.validate, req)
	if errs != nil {
		respond.Errors(w, http.StatusBadRequest, errs)
		return
	}

	create, err := h.useCase.Create(r.Context(), &req)
	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			respond.Error(w, http.StatusBadRequest, message.ErrBadRequest)
			return
		}
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	respond.Json(w, http.StatusCreated, account.Resource(create))
}
