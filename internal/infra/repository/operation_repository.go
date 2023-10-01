package repository

import (
	"database/sql"

	"github.com/LeoVieiraS/api_go_example/internal/entity"
)

type OperationRepository struct {
	DB *sql.DB
}

func NewOperationRepository(db *sql.DB) *OperationRepository {
	return &OperationRepository{
		DB: db,
	}
}

func (r *OperationRepository) Create(operation *entity.Operations) (*entity.Operations, error) {
	_, err := r.DB.Exec("INSERT INTO transactions (id, tipo_movimentacao, mercado, nome_instituicao, codigo_negociacao, quantidade, preco, valor, data) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", operation.Id, operation.TipoMovimentacao, operation.Mercado, operation.NomeInstituicao, operation.CodigoNegociacao, operation.Quantidade, operation.Preco, operation.Valor, operation.Data)
	if err != nil {
		return nil, err
	}
	return operation, nil

}

func (r *OperationRepository) FindAll() ([]*entity.Operations, error) {
	rown, err := r.DB.Query("SELECT * FROM transactions")
	if err != nil {
		return nil, err
	}
	defer rown.Close()

	var operations []*entity.Operations
	for rown.Next() {
		var operation entity.Operations
		if err := rown.Scan(
			&operation.Id,
			&operation.TipoMovimentacao,
			&operation.Mercado,
			&operation.NomeInstituicao,
			&operation.CodigoNegociacao,
			&operation.Quantidade,
			&operation.Preco,
			&operation.Valor,
			&operation.Data); err != nil {
			return nil, err
		}
		operations = append(operations, &operation)
	}

	return operations, nil

}
