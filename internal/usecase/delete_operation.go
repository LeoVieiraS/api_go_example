package usecase

import "github.com/LeoVieiraS/api_go_example/internal/entity"

type DeleteOperationDtoResponse struct {
	Success bool
	Errors  []string
}

type DeleteOperationUseCase struct {
	OperationRepository entity.OperationRepository
}

func NewDeleteOperationUseCase(operationRepository entity.OperationRepository) *DeleteOperationUseCase {
	return &DeleteOperationUseCase{
		OperationRepository: operationRepository,
	}
}

func (d *DeleteOperationUseCase) Execute(operationId string) (err error) {
	err = d.OperationRepository.Delete(operationId)
	return
}
