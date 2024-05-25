package response

import (
	e "go-ddd-practice/pkg/errors"
	"go-ddd-practice/pkg/usecase/dto"
	"net/http"

	"golang.org/x/exp/slices"

	"github.com/gin-gonic/gin"
)

func Res(c *gin.Context, result *dto.Dto) {
	status := http.StatusOK

	if slices.Contains(e.BadRequest, result.Error) {
		status = http.StatusBadRequest
	} else if slices.Contains(e.UnAuthorized, result.Error) {
		status = http.StatusUnauthorized
	} else if status == http.StatusInternalServerError {
		status = http.StatusInternalServerError
	}

	if status != http.StatusOK {
		c.JSON(status, gin.H{
			"Error": result.Error.Error(),
		})
	} else {
		c.JSON(status, gin.H{
			"Response": result,
		})
	}

}
