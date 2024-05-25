package controller

import (
	e "go-ddd-practice/pkg/errors"

	"go-ddd-practice/pkg/controller/response"
	"go-ddd-practice/pkg/model"
	"go-ddd-practice/pkg/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func appControllers(r *gin.RouterGroup) {

	service := usecase.NewAppService()

	// アカウントにログインします。
	{
		r.POST("/login", func(c *gin.Context) {
			param := new(model.VerLogin)
			if err := c.Bind(&param); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"Error": e.LOGIN_BAD_REQUEST.Error(),
				})
				return
			}

			result := service.Login(param)
			response.Res(c, result)
		})
	}

}
