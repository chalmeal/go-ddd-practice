package controller

import (
	e "go-ddd-practice/pkg/errors"

	"go-ddd-practice/pkg/controller/response"
	"go-ddd-practice/pkg/model"
	"go-ddd-practice/pkg/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func masterControllers(r *gin.RouterGroup) {

	service := usecase.NewMasterService()

	// アカウントを1件取得します。
	{
		r.GET("/:id", func(c *gin.Context) {
			id := c.Param("id")
			result := service.GetAccount(id)

			response.Res(c, result)
		})
	}

	// アカウントを全件取得します。
	{
		r.GET("", func(c *gin.Context) {
			result := service.GetAccountAll()

			response.Res(c, result)
		})
	}

	// アカウントを1件登録します。
	{
		r.POST("", func(c *gin.Context) {
			param := new(model.RegAccount)
			if err := c.Bind(&param); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"Error": e.REGISTER_ACCOUNT_BAD_REQUEST.Error(),
				})
				return
			}

			result := service.RegisterAccount(param)
			if result.Error != nil {
				response.Res(c, result)
				return
			}

			response.Res(c, result)
		})
	}

	// アカウントを編集します。
	{
		r.PUT("/edit", func(c *gin.Context) {
			param := new(model.EdtAccount)
			if err := c.Bind(&param); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"Error": e.EDIT_ACCOUNT_BAD_REQUEST.Error(),
				})
				return
			}

			result := service.EditAccount(param)
			if result.Error != nil {
				response.Res(c, result)
				return
			}

			response.Res(c, result)
		})
	}

	// アカウントを削除します。
	{
		type specifyDeleteAccount struct {
			Id string `json:"id" binding:"required"`
		}

		r.DELETE("/delete", func(c *gin.Context) {
			param := new(specifyDeleteAccount)
			if err := c.Bind(&param); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"Error": e.DELETE_ACCOUNT_BAD_REQUEST.Error(),
				})
				return
			}

			result := service.DeleteAccount(param.Id)
			if result.Error != nil {
				response.Res(c, result)
				return
			}

			response.Res(c, result)
		})
	}

}
