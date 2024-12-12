package main

import (
	"os"
	"{{module_name}}/internal/di"

	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var ginLambda *ginadapter.GinLambdaV2
var router *gin.Engine

func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	app, err := di.InitializeApp()

	if err != nil {
		log.Fatal().AnErr("Failed to initialize app: %v", err)
	}

	router = gin.Default()

	app.TransactionTypeController.RegisterRoutes(router)

	ginLambda = ginadapter.NewV2(router)
}

func main() {
	if inLambda() {
		lambda.Start(ginLambda.ProxyWithContext)
	} else {
		_ = router.Run()
	}
}

func inLambda() bool {
	return os.Getenv("LAMBDA_TASK_ROOT") != ""
}
