package authorization

import (
	"github.com/pets-shelters/backend-svc/internal/controller/helpers"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pkg/errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/pets-shelters/backend-svc/internal/entity"
)

type registrationRequest struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Logo        string `json:"logo" binding:"required,url"`
	City        string `json:"city" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required,len=12"`
	Instagram   string `json:"instagram" binding:"url"`
	Facebook    string `json:"facebook" binding:"url"`
}

func (r *routes) registration(c *gin.Context) {
	var request registrationRequest
	err := c.BindJSON(&request)
	if err != nil {
		r.log.Error(err, "failed to bind json - registration")
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.FormBadRequestError(err.Error()))
		return
	}

	err = r.useCase.Registration(
		c.Request.Context(),
		entity.Shelter{
			Name:        request.Name,
			Logo:        request.Logo,
			PhoneNumber: request.PhoneNumber,
			City:        request.City,
			Instagram:   request.Instagram,
			Facebook:    request.Facebook,
			CreatedAt:   time.Now(),
		},
		request.Email,
	)
	if err != nil {
		switch errors.Cause(err).(type) {
		case exceptions.UserExistsException:
			c.AbortWithStatusJSON(http.StatusConflict, helpers.FormCustomError(helpers.UserAlreadyExists, err.Error()))
		default:
			r.log.Error(err, "failed to process usecase - registration")
			c.AbortWithStatusJSON(http.StatusInternalServerError, helpers.FormInternalError(err.Error()))
		}
		return
	}

	c.Status(http.StatusOK)
}
