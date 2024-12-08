package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/LeoVieiraS/api_go_example/internal/infra/repository"
	"github.com/LeoVieiraS/api_go_example/internal/infra/web"
	"github.com/LeoVieiraS/api_go_example/internal/usecase"

	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./transactions.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	operationRepository := repository.NewOperationRepository(db)
	listOperationUsecase := usecase.NewListOperationsUseCase(operationRepository)
	createOperationUsecase := usecase.NewCreateOperationUseCase(operationRepository)
	deleteOperationUseCase := usecase.NewDeleteOperationUseCase(operationRepository)
	operationHendlers := web.NewOperationHendlers(listOperationUsecase, createOperationUsecase, deleteOperationUseCase)

	listOperationUsecase.Execute()
	if err != nil {
		fmt.Println(err)
	}

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	r.Get("/operations", operationHendlers.ListOperationsHandler)
	r.Post("/operations", operationHendlers.CreateOperationHandler)
	r.Delete("/operations/{operation_id}", operationHendlers.DeleteOperationHandler)

	http.ListenAndServe(":8000", r)

}
