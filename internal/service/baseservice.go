package service

import (
	"github.com/doublehops/dhapi-example/internal/model"
)

//type App struct {
//	app *app.App
//}

func HasPermission(ID int32, record model.BaseModel) bool {
	return ID == record.UserID
}
