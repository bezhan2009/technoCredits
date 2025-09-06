package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"technoCredits/internal/app/service"
	"technoCredits/pkg/errs"
	"technoCredits/pkg/logger"
)

func CheckUserGroupPermissions(c *gin.Context) {
	userID := c.GetUint(UserIDCtx)

	groupID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errs.ErrInvalidID.Error()})

		return
	}

	_, err = service.GetAllUserGroupsByID(userID, uint(groupID))
	if err != nil {
		logger.Error.Printf("CheckUserGroupPermissions err: %v", err)
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"error": errs.ErrPermissionDenied.Error()})

		return
	}

	c.Next()
}
