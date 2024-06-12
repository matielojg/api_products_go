# API de Produtos em GO

Este projeto é uma API de Produtos desenvolvida em Go, utilizando o framework Gin, Docker e Postgres para persistência de dados.

## Requisitos

- Go 1.18+
- Docker
- Docker Compose

## Instalação

1. Inicie um novo módulo Go:

   ```bash
   go mod init api-products_go
   ```
2. Baixe as dependências:

   ```bash
   go mod tidy
   ```
3. Inicie os containers Docker:

   ```bash
   docker-compose up -d
   ```
4. Build a imagem Docker:

   ```bash
   docker build -t go-api-tutorial .
   ```
## Execução

Para executar a API localmente, utilize:

```bash
cd cmd/
go run main.go
```

### Testando a API

#### Adicionar um produto

```bash
curl -X POST http://localhost:8000/product -H "Content-Type: application/json" -d '{ "nome": "Product 1", "price":100}'
```

#### Listar todos os produtos

```bash
curl http://localhost:8000/products
```

#### Buscar um produto por ID

```bash
curl http://localhost:8080/product/{id}
```

### Fonte de estudo: [Go Lab Tutoriais](https://www.youtube.com/watch?v=3p4mpId_ZU8)

```
.
├── cmd
│   └── main.go
├── controller
│   └── product_controller.go
├── db
│   └── conn.go
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── model
│   ├── product.go
│   └── response.go
├── README.md
├── repository
│   └── product_repository.go
└── usecase
    └── product_usecase.go
```
