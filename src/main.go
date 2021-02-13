package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NoahWTeng/go-clear-architecture/infrastructure/db"

	"github.com/NoahWTeng/go-clear-architecture/cases"

	"github.com/NoahWTeng/go-clear-architecture/interface/controllers"

	"github.com/NoahWTeng/go-clear-architecture/infrastructure/router"

	"github.com/NoahWTeng/go-clear-architecture/interface/repository"
)

var (
	httpRouter = router.NewMuxRouter()
	dbHandler  db.DBHandler
)

func getBookController() controllers.BookController {
	bookRepo := repository.NewBookRepo(dbHandler)
	bookInteractor := cases.NewBookInteractor(bookRepo)
	bookController := controllers.NewBookController(bookInteractor)
	return *bookController
}

func getAuthorController() controllers.AuthorController {
	authorRepo := repository.NewAuthorRepo(dbHandler)
	authorInteractor := cases.NewAuthorInteractor(authorRepo)
	authorController := controllers.NewAuthorController(authorInteractor)
	return *authorController
}

func main() {

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "App is up and running..")
	})
	var err error
	dbHandler, err = db.NewDBHandler("mongodb://localhost:27017", "bookstore")
	if err != nil {
		log.Println("Unable to connect to the DataBase")
		return
	}
	bookController := getBookController()
	authorController := getAuthorController()
	httpRouter.POST("/book/add", bookController.Add)
	httpRouter.GET("/book", bookController.FindAll)
	httpRouter.POST("/author/add", authorController.Add)
	httpRouter.SERVE(":8000")
}
