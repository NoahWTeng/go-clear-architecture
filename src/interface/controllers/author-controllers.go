package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/NoahWTeng/go-clear-architecture/common"

	"github.com/NoahWTeng/go-clear-architecture/cases"
)

type AuthorController struct {
	authorInteracts cases.AuthorInteracts
}

func NewAuthorController(authorInteracts cases.AuthorInteracts) *AuthorController {
	return &AuthorController{authorInteracts}
}

func (controller *AuthorController) Add(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var author common.Author
	err := json.NewDecoder(req.Body).Decode(&author)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid Payload"})
		return
	}
	err2 := controller.authorInteracts.CreateAuthor(author)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
}
