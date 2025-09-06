package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"technoCredits/pkg/errs"
)

func CheckUserGroupPermissions(c *gin.Context) {
	groupID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errs.ErrInvalidID.Error()})

		return
	}

}
