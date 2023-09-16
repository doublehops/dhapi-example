package main

import (
	"github.com/gin-gonic/gin"

	"github.com/doublehops/dhapi-example/internal/routes"
)

func main() {
	r := gin.New()
	//r.Use(logger.Logger()) // add middleware for all endpoints.
	r.Use(gin.Recovery())
	routes.GetRoutes(r)

	r.Run(":8080")
}
