package infrastructure

import (
	"net/http"
	"real-state-api/docs"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

func NewEngine(env *Env) *gin.Engine {
	engine := gin.Default()
	engine.Use(cors.Default())

	engine.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	engine.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"code": "METHOD_NOT_FOUND", "error": "Method not found"})
	})

	engine.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	docs.SwaggerInfo.Title = "Real Estate Advertising API"
	docs.SwaggerInfo.Description = "Application that manages real estate and exchange rates"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = env.Host
	docs.SwaggerInfo.BasePath = env.BasePath

	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return engine
}

func StartHttpServer(engine *gin.Engine, logger *zap.SugaredLogger, env *Env) {

	port := ":8080"
	if env.Port != 0 {
		port = ":" + strconv.Itoa(env.Port)
	}

	err := engine.Run(port)

	if err != nil {
		logger.Fatalw("erro ao iniciar servidor HTTP", "erro", err)
	}
}
