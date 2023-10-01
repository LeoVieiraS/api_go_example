package usecase

import (
	"github.com/LeoVieiraS/api_go_example/internal/entity"
)

type ListOperationsOutputDto struct {
	Id               string  `json:"id"`
	TipoMovimentacao string  `json:"tipo_movimentacao"`
	Mercado          string  `json:"mercado"`
	NomeInstituicao  string  `json:"nome_instituicao"`
	CodigoNegociacao string  `json:"codigo_negociacao"`
	Preco            float64 `json:"preco"`
	Quantidade       float64 `json:"quantidade"`
	Valor            float64 `json:"valor"`
	Data             string  `json:"data"`
}

type ListOperationsUseCase struct {
	OperationRepository entity.OperationRepository
}

func NewListOperationsUseCase(operationRepository entity.OperationRepository) *ListOperationsUseCase {
	return &ListOperationsUseCase{
		OperationRepository: operationRepository,
	}
}

func (o *ListOperationsUseCase) Execute() ([]*ListOperationsOutputDto, error) {
	operations, err := o.OperationRepository.FindAll()
	if err != nil {
		return nil, err
	}
	layout := "2006-01-02 15:04:05"

	var operationsOutputDto []*ListOperationsOutputDto
	var formated_date string

	for _, operation := range operations {
		formated_date = operation.Data.Format(layout)

		operationsOutputDto = append(operationsOutputDto, &ListOperationsOutputDto{
			Id:               operation.Id,
			TipoMovimentacao: operation.TipoMovimentacao,
			Mercado:          operation.Mercado,
			NomeInstituicao:  operation.NomeInstituicao,
			CodigoNegociacao: operation.CodigoNegociacao,
			Preco:            operation.Preco,
			Quantidade:       operation.Quantidade,
			Valor:            operation.Valor,
			Data:             formated_date,
		})

	}
	return operationsOutputDto, nil
}
