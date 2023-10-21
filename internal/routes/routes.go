package routes

import (
	group "github.com/mythrnr/httprouter-group"

	"github.com/doublehops/dhapi-example/internal/app"
	"github.com/doublehops/dhapi-example/internal/handlers/author"
)

func GetV1Routes(app *app.App) *group.RouteGroup {
	authorHandle := author.New(app)

	authorGroup := group.New("/author")
	authorGroup.GET(authorHandle.GetAll)
	authorGroup.Children(
		group.New("/:id").GET(authorHandle.GetByID),
		group.New("").POST(authorHandle.Create),
		group.New("/:id").PUT(authorHandle.UpdateByID),
		group.New("/:id").DELETE(authorHandle.DeleteByID),
	)

	g := group.New("/v1").Children(
		authorGroup,
	)

	return g
}
