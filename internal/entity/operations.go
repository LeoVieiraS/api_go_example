package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type OperationRepository interface {
	Create(operation *Operations) (*Operations, error)
	FindAll() ([]*Operations, error)
}

type Operations struct {
	Id               string
	TipoMovimentacao string
	Mercado          string
	NomeInstituicao  string
	CodigoNegociacao string
	Preco            float64
	Quantidade       float64
	Valor            float64
	Data             time.Time
}

func NewOperations(
	tipoMovimentacao string,
	mercado string,
	nomeInstituicao string,
	codigoNegociacao string,
	preco float64,
	quantidade float64,
	valor float64,
	data time.Time) *Operations {
	return &Operations{
		Id:               uuid.NewV4().String(),
		TipoMovimentacao: tipoMovimentacao,
		Mercado:          mercado,
		NomeInstituicao:  nomeInstituicao,
		CodigoNegociacao: codigoNegociacao,
		Preco:            preco,
		Quantidade:       quantidade,
		Valor:            valor,
		Data:             data,
	}
}
