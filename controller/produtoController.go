package controller

import (
	"database/sql"
	"net/http"

	"github.com/belzebells/desafio-desenvolvedora-laiob/database"
	"github.com/belzebells/desafio-desenvolvedora-laiob/model"
	"github.com/gin-gonic/gin"
)

func ListarProdutos(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, nome, preco, quantidade, descricao FROM produtos")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar produtos"})
		return
	}
	defer rows.Close()

	var produtos []model.Produto
	for rows.Next() {
		var p model.Produto
		if err := rows.Scan(&p.ID, &p.Nome, &p.Preco, &p.Quantidade, &p.Descricao); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar produtos"})
			return
		}
		produtos = append(produtos, p)
	}

	if produtos == nil {
		produtos = []model.Produto{}
	}

	c.JSON(http.StatusOK, produtos)
}

func ObterProduto(c *gin.Context) {
	id := c.Param("id")

	var p model.Produto
	err := database.DB.QueryRow("SELECT id, nome, preco, quantidade, descricao FROM produtos WHERE id = ?", id).
		Scan(&p.ID, &p.Nome, &p.Preco, &p.Quantidade, &p.Descricao)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar produto"})
		}
		return
	}

	c.JSON(http.StatusOK, p)
}

func CriarProduto(c *gin.Context) {
	var novo model.Produto
	if err := c.ShouldBindJSON(&novo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	res, err := database.DB.Exec("INSERT INTO produtos (nome, preco, quantidade, descricao) VALUES (?, ?, ?, ?)",
		novo.Nome, novo.Preco, novo.Quantidade, novo.Descricao)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar produto"})
		return
	}

	id, _ := res.LastInsertId()
	novo.ID = int(id)

	c.JSON(http.StatusCreated, novo)
}

func AtualizarProduto(c *gin.Context) {
	id := c.Param("id")
	var p model.Produto

	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	res, err := database.DB.Exec("UPDATE produtos SET nome=?, preco=?, quantidade=?, descricao=? WHERE id=?",
		p.Nome, p.Preco, p.Quantidade, p.Descricao, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar produto"})
		return
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produto atualizado"})
}

func DeletarProduto(c *gin.Context) {
	id := c.Param("id")

	res, err := database.DB.Exec("DELETE FROM produtos WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar produto"})
		return
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produto deletado"})
}
