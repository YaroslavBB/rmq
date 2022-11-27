package rest

import (
	"net/http"
	"rabbitmq/rest_service/internal/entity/global"
	"rabbitmq/rest_service/internal/repository"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	server     *gin.Engine
	numberRepo repository.PhoneRepository
}

func NewRest(server *gin.Engine, numberRepo repository.PhoneRepository) *Rest {
	rest := &Rest{
		server:     server,
		numberRepo: numberRepo,
	}

	server.POST("/get_number", rest.FindPhoneNumber)

	return rest
}

func (r *Rest) Run() {
	r.server.Run("localhost:8080")
}

func (r *Rest) FindPhoneNumber(c *gin.Context) {
	type Param struct {
		RequestID int `json:"request_id"`
	}

	var p Param

	if err := c.BindJSON(&p); err != nil {
		c.JSON(http.StatusUnprocessableEntity, global.ErrIncorrectParam.Error())
		return
	}

	data, err := r.numberRepo.FindPhoneNumber(p.RequestID)
	if err != nil {
		if err == global.ErrNoData {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}
		c.AbortWithError(http.StatusInternalServerError, global.ErrInternalError)

		return
	}
	c.JSON(http.StatusOK, data)
}
