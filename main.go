package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // Driver PostgreSQL
)

type Produto struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

func main() {
	// Conecta ao banco de dados PostgreSQL
	db, err := sql.Open("postgres", "host=localhost port=5440 user=postgres password=1234 dbname=go_db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verifica se a conexão está ativa
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	// Cria as tabelas no banco de dados (se necessário)
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS produtos (
        id SERIAL PRIMARY KEY,
        nome VARCHAR(255) NOT NULL
    )`)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	// Rotas da API
	router.GET("/produtos", getProdutos(db))
	router.POST("/produtos", createProduto(db))
	// Inicia o servidor HTTP na porta 8080
	log.Fatal(router.Run(":8080"))
}

func getProdutos(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT * FROM produtos")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()
		var produtos []Produto
		for rows.Next() {
			var produto Produto
			err := rows.Scan(&produto.ID, &produto.Nome)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			produtos = append(produtos, produto)
		}
		c.JSON(http.StatusOK, produtos)
	}
}
func createProduto(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var produto Produto
		err := c.ShouldBindJSON(&produto)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Insira o novo produto no banco de dados
		stmt, err := db.Prepare("INSERT INTO produtos (nome) VALUES ($1)")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer stmt.Close()
		_, err = stmt.Exec(produto.Nome)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "produto criado com sucesso"})
	}
}
