package main

import (
	"github.com/doublehops/dhapi-example/internal/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.New()
	r.ForwardedByClientIP = true
	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatalf("error setting trusted proxy. %s", err)
	}

	//r.Use(logger.Logger()) // add middleware for all endpoints.

	r.Use(gin.Recovery())
	routes.GetRoutes(r)

	r.Run(":8080")
}
