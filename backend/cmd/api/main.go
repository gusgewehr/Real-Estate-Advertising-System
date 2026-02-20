package main

import (
	"net/http"
	"real-state-api/internal/infrastructure"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	logger := infrastructure.NewLogger()

	env := infrastructure.NewEnv(".env", logger)

	engine := gin.Default()
	engine.Use(cors.Default())

	engine.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	engine.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"code": "METHOD_NOT_FOUND", "error": "Method not found"})
	})

	port := ":8080"
	if env.Port != 0 {
		port = ":" + strconv.Itoa(env.Port)
	}

	err := engine.Run(port)

	if err != nil {
		logger.Fatalw("erro ao iniciar servidor HTTP", "erro", err)
	}

}
