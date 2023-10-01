CREATE TABLE IF NOT EXISTS transactions (
    id TEXT PRIMARY KEY,
    tipo_movimentacao TEXT,
    mercado TEXT,
    nome_instituicao TEXT,
    codigo_negociacao TEXT,
    quantidade INTEGER,
    preco REAL,
    valor REAL,
    data date
)


--drop table transactions;

SELECT * FROM transactions;

