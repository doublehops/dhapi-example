package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/doublehops/dhapi/resp"
	"github.com/doublehops/dhapi/validator"
)

type User struct {
	Username     string `json:"username"`
	EmailAddress string `json:"emailAddress"`
	Age          int    `json:"age"`
}

func GetUser(c *gin.Context) {
	user := User{
		Username:     c.MustGet("username").(string),
		EmailAddress: c.MustGet("emailAddress").(string),
	}

	c.JSON(http.StatusOK, resp.GetSingleItemResp(user))
}

func ListUser(c *gin.Context) {
	users := []User{
		{
			Username:     "Alice",
			EmailAddress: "alice@example.com",
		},
		{
			Username:     "Bob",
			EmailAddress: "bob@example.com",
		},
		{
			Username:     "Carol",
			EmailAddress: "carol@example.com",
		},
	}

	p := resp.Pagination{
		CurrentPage: 1,
		PerPage:     10,
		PageCount:   22,
		TotalCount:  229,
	}

	resp.GetListResp(users, p)
}

func getRules(user *User) []validator.Rule {
	return []validator.Rule{
		{"username", user.Username, true, []validator.ValidateFuncs{validator.LengthInRange(3, 8, "")}},
		{"emailAddress", user.EmailAddress, true, []validator.ValidateFuncs{validator.EmailAddress("")}},
		{"age", user.Age, true, []validator.ValidateFuncs{validator.MinValue(18, "")}},
	}
}

// UpdateUser - Validation error example.
// Example valid test request: curl -s -X PUT localhost:8080/v1/user -H "Content-Type: application/json" --data '{"username": "johns", "emailAddress": "john@example.com", "age": 30}'| jq; echo
// Example invalid test request: curl -s -X PUT localhost:8080/v1/user -H "Content-Type: application/json" --data '{"username": "j", "emailAddress": "john.smith", "age": 17}'| jq; echo
func UpdateUser(c *gin.Context) {
	var user User

	_ = c.ShouldBindJSON(&user)

	validationErrors := validator.RunValidation(getRules(&user))

	if len(validationErrors) > 0 {
		c.JSON(http.StatusBadRequest, resp.GetValidateErrResp(validationErrors))

		return
	}

	c.JSON(http.StatusOK, resp.GetSingleItemResp(user))
}
