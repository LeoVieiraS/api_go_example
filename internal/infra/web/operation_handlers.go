package web

import (
	"encoding/json"
	"net/http"

	"github.com/LeoVieiraS/api_go_example/internal/usecase"
)

type OperationHandlers struct {
	ListOperationsUseCase  *usecase.ListOperationsUseCase
	CreateOperationUseCase *usecase.CreateOperationUseCase
}

func NewOperationHendlers(
	listOperationUseCase *usecase.ListOperationsUseCase,
	createOperationUseCase *usecase.CreateOperationUseCase) *OperationHandlers {
	return &OperationHandlers{
		ListOperationsUseCase:  listOperationUseCase,
		CreateOperationUseCase: createOperationUseCase,
	}
}

func (h *OperationHandlers) ListOperationsHandler(w http.ResponseWriter, r *http.Request) {
	operations, err := h.ListOperationsUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(operations)
}

func (h *OperationHandlers) CreateOperationHandler(w http.ResponseWriter, r *http.Request) {
	var inputDto usecase.CreateOperationInputDto
	if err := json.NewDecoder(r.Body).Decode(&inputDto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	operation, err := h.CreateOperationUseCase.Execute(inputDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode(operation)
}
