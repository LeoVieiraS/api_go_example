package usecase

import (
	"time"

	"github.com/LeoVieiraS/api_go_example/internal/entity"
)

type CreateOperationInputDto struct {
	TipoMovimentacao string  `json:"tipo_movimentacao"`
	Mercado          string  `json:"mercado"`
	NomeInstituicao  string  `json:"nome_instituicao"`
	CodigoNegociacao string  `json:"codigo_negociacao"`
	Preco            float64 `json:"preco"`
	Quantidade       float64 `json:"quantidade"`
	Valor            float64 `json:"valor"`
	Data             string  `json:"data"`
}

type CreateOperationUseCase struct {
	OperationRepository entity.OperationRepository
}

func NewCreateOperationUseCase(operationRepository entity.OperationRepository) *CreateOperationUseCase {
	return &CreateOperationUseCase{
		OperationRepository: operationRepository,
	}
}

func (u *CreateOperationUseCase) Execute(inputDto CreateOperationInputDto) (*entity.Operations, error) {
	layout := "2006-01-02 15:04:05"
	parsedDate, _ := time.Parse(layout, inputDto.Data)
	operation := entity.NewOperations(
		inputDto.TipoMovimentacao,
		inputDto.Mercado,
		inputDto.NomeInstituicao,
		inputDto.CodigoNegociacao,
		inputDto.Preco,
		inputDto.Quantidade,
		inputDto.Valor,
		parsedDate,
	)

	return u.OperationRepository.Create(operation)
}
