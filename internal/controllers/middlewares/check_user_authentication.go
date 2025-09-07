package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"technoCredits/pkg/utils"
)

const (
	authorizationHeader = "Authorization"
	UserIDCtx           = "userID"
	UserRoleIDCtx       = "roleID"
	tokenQueryParam     = "token"
	UserIDCtx1          = "userID"
	UserRoleIDCtx1      = "roleID"
)

func CheckUserAuthentication(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)

	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "empty auth header",
		})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid auth header",
		})
		return
	}

	if len(headerParts[1]) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "token is empty",
		})
		return
	}

	accessToken := headerParts[1]

	claims, err := utils.ParseToken(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.Set(UserIDCtx, claims.UserID)
	c.Set(UserRoleIDCtx, claims.RoleID)

	c.Next()
}

func CheckUserAuthenticationQuery(c *gin.Context) {
	header := c.Query(tokenQueryParam)

	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "empty auth header",
		})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid auth header",
		})
		return
	}

	if len(headerParts[1]) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "token is empty",
		})
		return
	}

	accessToken := headerParts[1]

	claims, err := utils.ParseToken(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.Set(UserIDCtx, claims.UserID)
	c.Set(UserRoleIDCtx, claims.RoleID)

	c.Next()
}
