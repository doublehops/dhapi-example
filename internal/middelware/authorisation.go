package middelware

import (
	"context"
	"github.com/doublehops/dhapi-example/internal/app"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// AuthMiddleware will authenticate user by the bearer token passed in through the authorization header.
// todo - needs implementation.
func AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var userID int32 = 4

		r = r.WithContext(context.WithValue(r.Context(), app.UserIDKey, userID))
		log.Println(">>>>> middelware")
		next(w, r, ps)
	}
}
