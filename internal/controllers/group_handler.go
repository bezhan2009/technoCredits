package controllers

//var groupService *service.GroupService
//
//func InitGroupService(s *service.GroupService) {
//	groupService = s
//}
//
//// CreateGroup godoc
//// @Summary Создать группу
//// @Description Создаёт новую группу
//// @Tags groups
//// @Accept json
//// @Produce json
//// @Param group body models.Group true "Группа"
//// @Success 201 {object} models.Group
//// @Failure 400 {object} map[string]string
//// @Security ApiKeyAuth
//// @Router /groups [post]
//func CreateGroup(c *gin.Context) {
//	userID := c.GetUint(middlewares.UserIDCtx)
//
//	var group models.Group
//	if err := c.ShouldBindJSON(&group); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	group.OwnerID = userID
//
//	if err := groupService.Create(&group); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusCreated, group)
//}
//
//// GetGroupByID godoc
//// @Summary Получить группу по ID
//// @Description Возвращает группу по идентификатору
//// @Tags groups
//// @Produce json
//// @Param id path int true "ID группы"
//// @Success 200 {object} models.Group
//// @Failure 400 {object} map[string]string
//// @Failure 404 {object} map[string]string
//// @Security ApiKeyAuth
//// @Router /groups/{id} [get]
//func GetGroupByID(c *gin.Context) {
//	idStr := c.Param("id")
//	id, err := strconv.ParseUint(idStr, 10, 64)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
//		return
//	}
//
//	group, err := groupService.GetByID(uint(id))
//	if err != nil {
//		c.JSON(http.StatusNotFound, gin.H{"error": "Группа не найдена"})
//		return
//	}
//
//	c.JSON(http.StatusOK, group)
//}
//
//// UpdateGroup godoc
//// @Summary Обновить группу
//// @Description Обновляет данные группы
//// @Tags groups
//// @Accept json
//// @Produce json
//// @Param id path int true "ID группы"
//// @Param group body models.Group true "Группа"
//// @Success 200 {object} models.Group
//// @Failure 400 {object} map[string]string
//// @Failure 500 {object} map[string]string
//// @Security ApiKeyAuth
//// @Router /groups/{id} [put]
//func UpdateGroup(c *gin.Context) {
//	idStr := c.Param("id")
//	id, err := strconv.ParseUint(idStr, 10, 64)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
//		return
//	}
//
//	var group models.Group
//	if err := c.ShouldBindJSON(&group); err != nil {
//		HandleError(c, errs.ErrValidationFailed)
//		return
//	}
//
//	group.ID = uint(id)
//
//	if err := groupService.Update(&group); err != nil {
//		HandleError(c, err)
//		return
//	}
//
//	c.JSON(http.StatusOK, group)
//}
//
//// DeleteGroup godoc
//// @Summary Удалить группу
//// @Description Удаляет группу по ID
//// @Tags groups
//// @Param id path int true "ID группы"
//// @Success 204
//// @Failure 400 {object} map[string]string
//// @Failure 500 {object} map[string]string
//// @Security ApiKeyAuth
//// @Router /groups/{id} [delete]
//func DeleteGroup(c *gin.Context) {
//	idStr := c.Param("id")
//	id, err := strconv.ParseUint(idStr, 10, 64)
//	if err != nil {
//		HandleError(c, err)
//		return
//	}
//
//	if err := groupService.Delete(uint(id)); err != nil {
//		HandleError(c, err)
//		return
//	}
//
//	c.Status(http.StatusNoContent)
//}
