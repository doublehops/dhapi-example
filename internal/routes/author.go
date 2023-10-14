package routes

import (
	"github.com/doublehops/dhapi-example/internal/app"
	"github.com/gin-gonic/gin"

	"github.com/doublehops/dhapi-example/internal/handlers/author"
)

func authorHandle(rg *gin.RouterGroup, app *app.App) {
	// *****  USER  *****
	ag := rg.Group("/author")

	authorHandle := author.New(app)

	//User.GET("", authorHandle.ListUser)
	//User.GET("/bobby", customauth.Auth(), authorHandle.GetUser)
	ag.POST("", authorHandle.Create)
	//User.PUT("", authorHandle.UpdateUser)
}
