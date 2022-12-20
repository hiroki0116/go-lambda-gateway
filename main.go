package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	log.Println("gin cold start")
	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "CICD Success!",
		})
	})
	server.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Modified Hello World!",
		})
	})

	server.POST("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Test success with modified messsage",
		})
	})

	ginLambda = ginadapter.New(server)
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(handler)
}
