package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"technoCredits/internal/app/service"
	"technoCredits/pkg/errs"
)

func CheckUsersCardPermission(c *gin.Context) {
	cardID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errs.ErrInvalidID.Error()})
		return
	}

	userID := c.GetUint(UserIDCtx)

	_, err = service.GetCardExpenseByID(userID, uint(cardID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": errs.ErrUserNotFound.Error()})
		return
	}

	c.Next()
}
