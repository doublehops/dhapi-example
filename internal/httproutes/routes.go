package httproutes

import (
	"github.com/julienschmidt/httprouter"
	group "github.com/mythrnr/httprouter-group"

	"github.com/doublehops/dhapi-example/internal/app"
	"github.com/doublehops/dhapi-example/internal/handlers/author"
	"github.com/doublehops/dhapi-example/internal/handlers/user"
)

func GetV1Routes(app *app.App) *group.RouteGroup {
	userHandle := user.New(app)
	authorHandle := author.New(app)

	router := httprouter.New()
	router.GET("/user", userHandle.Hello)

	authorGroup := group.New("/user")
	authorGroup.GET(authorHandle.GetAll)
	authorGroup.Children(
		group.New("/:id").GET(authorHandle.GetByID),
		group.New("").POST(authorHandle.Create),
		group.New("/:id").PUT(authorHandle.Update),
	)

	g := group.New("/v1").Children(
		authorGroup,
	)

	return g
}
