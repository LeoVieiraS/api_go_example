package web

import (
	"encoding/json"

	"net/http"

	"log/slog"

	"github.com/LeoVieiraS/api_go_example/internal/infra/logging"
	"github.com/LeoVieiraS/api_go_example/internal/usecase"

	"github.com/go-chi/chi"
)

type OperationHandlers struct {
	logger                 *slog.Logger
	ListOperationsUseCase  *usecase.ListOperationsUseCase
	CreateOperationUseCase *usecase.CreateOperationUseCase
	DeleteOperationUseCase *usecase.DeleteOperationUseCase
}

func NewOperationHendlers(
	listOperationUseCase *usecase.ListOperationsUseCase,
	createOperationUseCase *usecase.CreateOperationUseCase,
	deleteOperationUseCase *usecase.DeleteOperationUseCase) *OperationHandlers {
	return &OperationHandlers{
		logger:                 logging.NewLogger("Operations"),
		ListOperationsUseCase:  listOperationUseCase,
		CreateOperationUseCase: createOperationUseCase,
		DeleteOperationUseCase: deleteOperationUseCase,
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(operation)
}

func (h *OperationHandlers) DeleteOperationHandler(w http.ResponseWriter, r *http.Request) {
	var response = usecase.DeleteOperationDtoResponse{}
	operationId := chi.URLParam(r, "operation_id")

	err := h.DeleteOperationUseCase.Execute(operationId)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {

		h.logger.Error(err.Error())

		response.Success = false
		response.Errors = []string{"An error occured. Operation was'nt deleted"}

		json.NewEncoder(w).Encode(response)
		return
	}
	response.Success = true
	response.Errors = []string{}
	json.NewEncoder(w).Encode(response)
}
