package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/doublehops/dhapi/responses"
)

type User struct {
	Username     string `json:"username"`
	EmailAddress string `json:"emailAddress"`
}

func GetUser(c *gin.Context) {
	user := User{
		Username:     c.MustGet("username").(string),
		EmailAddress: c.MustGet("emailAddress").(string),
	}

	responses.SingleItemResponse(c, user)
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

	p := responses.PaginationType{
		CurrentPage: 1,
		PerPage:     10,
		PageCount:   22,
		TotalCount:  229,
	}

	responses.MultiItemResponse(c, users, p)
}

type UpdateUserObj struct {
	Username string `json:"username" binding:"required,min=3"`
	Country  string `json:"country" binding:"min=3,max=3"`
}

// GetErrorMessages will return the list of custom validation error messages to use in response.
func GetErrorMessages() responses.CustomerErrorMessages {
	return responses.CustomerErrorMessages{
		"username": {"required": "This field is required"},
		"country":  {"required": "This field is required", "min": "Must be at least 3 characters long"},
	}
}

// UpdateUser - Validation error example.
// Example test request: curl -s -X PUT localhost:8080/v1/user -H "Content-Type: application/json" --data '{"username": "aaa", "country": "AUA"}'| jq; echo
func UpdateUser(c *gin.Context) {
	var user UpdateUserObj

	_ = c.ShouldBindJSON(&user)

	validationErrors := responses.Validate(c, &user, GetErrorMessages())

	if len(validationErrors) > 0 {
		responses.ValidationErrorResponse(c, 422, validationErrors)

		return
	}

	responses.SingleItemResponse(c, user)
}
