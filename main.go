package main

import (
    "github.com/belzebells/desafio-desenvolvedora-laiob/controller"
    "github.com/belzebells/desafio-desenvolvedora-laiob/database"
    "github.com/gin-gonic/gin"
)

// Middleware CORS manual
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

func main() {
    database.Conectar()

    r := gin.Default()

    // Usar CORS manual
    r.Use(CORSMiddleware())
    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    // Rotas
    r.GET("/produtos", controller.ListarProdutos)
    r.GET("/produtos/:id", controller.ObterProduto)
    r.POST("/produtos", controller.CriarProduto)
    r.PUT("/produtos/:id", controller.AtualizarProduto)
    r.DELETE("/produtos/:id", controller.DeletarProduto)

    r.Run(":8080")
}
