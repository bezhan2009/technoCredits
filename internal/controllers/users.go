package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"technoCredits/internal/app/models"
	"technoCredits/internal/app/service"
	"technoCredits/internal/controllers/middlewares"
	"technoCredits/pkg/errs"
)

func GetMyDataUser(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)

	user, err := service.GetUserByID(userID)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	user.ID = c.GetUint(middlewares.UserIDCtx)

	err = service.UpdateUser(user)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Worker updated successfully"})
}

func UpdateUsersPassword(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)

	var newPassword models.NewUsersPassword
	if err := c.ShouldBindJSON(&newPassword); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	if newPassword.NewPassword == "" {
		HandleError(c, errs.ErrPasswordIsEmpty)
		return
	}

	err := service.UpdateUserPassword(userID, newPassword.OldPassword, newPassword.NewPassword)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Password updated successfully",
	})
}
