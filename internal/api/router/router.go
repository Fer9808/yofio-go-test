package router

import (
	"fmt"

	"github.com/Fer9808/yofio-go-test/internal/api/controllers"
	"github.com/Fer9808/yofio-go-test/internal/api/middlewares"
	"github.com/Fer9808/yofio-go-test/internal/api/services"
	"github.com/Fer9808/yofio-go-test/internal/pkg/persistence"

	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	app := gin.New()

	// Config Dependencies
	repo := persistence.GetAssignmentsRepository()
	creditService := services.NewCreditServiceImpl(repo)

	ctrl := controllers.NewController(creditService, repo)

	// Logging to a file.
	f, _ := os.Create("log/api.log")
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(f)

	// Middlewares
	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - - [%s] \"%s %s %s %d %s \" \" %s\" \" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format("02/Jan/2006:15:04:05 -0700"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	app.Use(gin.Recovery())
	app.Use(middlewares.CORS())
	app.NoRoute(middlewares.NoRouteHandler())

	// Routes
	app.POST("/api/credit-assignment", ctrl.CreateAssignments)
	app.POST("/api/statistics", ctrl.GetStatistics)

	return app
}
