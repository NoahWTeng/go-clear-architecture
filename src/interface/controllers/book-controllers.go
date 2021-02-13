package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/NoahWTeng/go-clear-architecture/common"

	"github.com/NoahWTeng/go-clear-architecture/cases"
)

type BookController struct {
	bookInteracts cases.BookInteractor
}

func NewBookController(bookInteracts cases.BookInteractor) *BookController {
	return &BookController{bookInteracts}
}

func (controller *BookController) Add(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var book common.Book
	err := json.NewDecoder(req.Body).Decode(&book)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid Payload"})
		return
	}
	err2 := controller.bookInteracts.CreateBook(book)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
}

func (controller *BookController) FindAll(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	results, err2 := controller.bookInteracts.FindAll()
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(results)
}
