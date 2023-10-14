package routes

import (
	"fmt"
	"github.com/doublehops/dhapi-example/internal/app"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/doublehops/dhapi-example/internal/handlers/user"
	"github.com/doublehops/dhapi-example/internal/middleware/customauth"
)

func GetRoutes(router *gin.Engine, app *app.App) {
	v1 := router.Group("/v1")
	v1routes(v1, app)
}

func v1routes(rg *gin.RouterGroup, app *app.App) {
	// *****  USER  *****
	User := rg.Group("/user")

	userHandle := user.New(app)

	User.GET("", userHandle.ListUser)
	User.GET("/bobby", customauth.Auth(), userHandle.GetUser)
	User.PUT("", userHandle.UpdateUser)

	User.GET("/middleware-test", func(c *gin.Context) {
		example, _ := c.MustGet("example").(string)

		log.Println(example)
		c.JSON(http.StatusOK, fmt.Sprintf("User: %s", example))
	})

	User.GET("/by-id/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, fmt.Sprintf("User: %s", c.Param("id")))
	})

	authorHandle(rg, app)
}
