package middelware

import (
	"context"
	"github.com/doublehops/dhapi-example/internal/app"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var userID int32 = 11

		r = r.WithContext(context.WithValue(r.Context(), app.UserIDKey, userID))
		log.Println(">>>>> middelware")
		next(w, r, ps)
	}
}
