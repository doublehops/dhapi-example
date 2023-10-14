package author

import (
	"github.com/doublehops/dhapi-example/internal/app"
	"github.com/doublehops/dhapi-example/internal/model"
	"github.com/doublehops/dhapi-example/internal/repository/repositoryauthor"
	"github.com/doublehops/dhapi-example/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/doublehops/dhapi/resp"
)

type Handle struct {
	app *app.App
	ar  *repositoryauthor.RepositoryAuthor
	as  *service.AuthorService
}

func New(app *app.App) *Handle {
	ar := repositoryauthor.New(app.DB, app.Log)
	return &Handle{
		app: app,
		ar:  ar,
		as:  service.New(app, ar),
	}
}

func (h *Handle) Create(c *gin.Context) {

	//jsonData, esrr := io.ReadAll(c.Request.Body)
	//if esrr != nil {
	//	// Handle error
	//}
	//fmt.Printf(">>>>>> %s", jsonData)

	h.app.Log.Info(c, "Request made to CreateAuthor")

	var author *model.Author
	//fmt.Printf("body: %s", c.Request.Body)
	err := c.BindJSON(&author)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Unable to parse request")
	}

	a, err := h.as.Create(c, author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Unable to process request")

		return
	}

	c.JSON(http.StatusOK, resp.GetSingleItemResp(a))
}

//func (h *Handle) GetUser(c *gin.Context) {
//	c.Set("traceID", "AB19-B891-CA8D")
//	c.Set("userID", 123)
//
//	h.app.Log.Info(c, "**** call to /v1/user", "custom", "hello")
//	h.app.Log.Error(c, "**** ERROR /v1/user")
//
//	user := User{
//		Username:     c.MustGet("username").(string),
//		EmailAddress: c.MustGet("emailAddress").(string),
//	}
//
//	c.JSON(http.StatusOK, resp.GetSingleItemResp(user))
//}
//
//func (h *Handle) ListUser(c *gin.Context) {
//	users := []User{
//		{
//			Username:     "Alice",
//			EmailAddress: "alice@example.com",
//		},
//		{
//			Username:     "Bob",
//			EmailAddress: "bob@example.com",
//		},
//		{
//			Username:     "Carol",
//			EmailAddress: "carol@example.com",
//		},
//	}
//
//	p := resp.Pagination{
//		CurrentPage: 1,
//		PerPage:     10,
//		PageCount:   22,
//		TotalCount:  229,
//	}
//
//	resp.GetListResp(users, p)
//}
//
//// nolint:govet
//func (u *User) getRules() []validator.Rule {
//	return []validator.Rule{
//		{"username", u.Username, true, []validator.ValidationFuncs{validator.LengthInRange(3, 8, "")}},
//		{"emailAddress", u.EmailAddress, true, []validator.ValidationFuncs{validator.EmailAddress("")}},
//		{"age", u.Age, true, []validator.ValidationFuncs{validator.MinValue(18, "")}},
//	}
//}
//
//// UpdateUser - Validation error example.
//// Example valid test request: curl -s -X PUT localhost:8080/v1/user -H "Content-Type: application/json" --data '{"username": "johns", "emailAddress": "john@example.com", "age": 30}'| jq; echo
//// Example invalid test request: curl -s -X PUT localhost:8080/v1/user -H "Content-Type: application/json" --data '{"username": "j", "emailAddress": "john.smith", "age": 17}'| jq; echo
//func (h *Handle) UpdateUser(c *gin.Context) {
//	var user User
//
//	h.app.Log.Info(c, fmt.Sprintf("RequestMade: %s %s", c.Request.Method, c.Request.RequestURI))
//	_ = c.ShouldBindJSON(&user)
//
//	validationErrors := validator.RunValidation(user.getRules())
//
//	if len(validationErrors) > 0 {
//		c.JSON(http.StatusBadRequest, resp.GetValidateErrResp(validationErrors))
//
//		return
//	}
//
//	c.JSON(http.StatusOK, resp.GetSingleItemResp(user))
//}
