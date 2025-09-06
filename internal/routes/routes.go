package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	_ "technoCredits/docs"
	"technoCredits/internal/app/service"
	"technoCredits/internal/controllers"
	"technoCredits/internal/controllers/middlewares"
	"technoCredits/internal/repository"
)

func InitRoutes(r *gin.Engine, db *gorm.DB) *gin.Engine {

	r.GET("/ping", controllers.Ping)

	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", controllers.SignUp)
		auth.POST("/sign-in", controllers.SignIn)
		auth.POST("/refresh", controllers.RefreshToken)
	}

	groupRepo := repository.NewGroupRepository(db)
	groupSrv := service.NewGroupService(groupRepo)
	controllers.InitGroupService(groupSrv)

	groupRoutes := r.Group("/groups", middlewares.CheckUserAuthentication)
	{
		groupRoutes.POST("/", controllers.CreateGroup)
		groupRoutes.GET("/:id", controllers.GetGroupByID)
		groupRoutes.PUT("/:id", controllers.UpdateGroup)
		groupRoutes.DELETE("/:id", controllers.DeleteGroup)
	}

	groupMemberRepo := repository.NewGroupMemberRepository(db)
	groupMemberSrv := service.NewGroupMemberService(groupMemberRepo)
	groupMemberHandler := controllers.NewGroupMemberHandler(groupMemberSrv)

	members := groupRoutes.Group("/:id/members", middlewares.CheckUserAuthentication)
	{
		members.POST("/", groupMemberHandler.AddMember)
		members.GET("/", groupMemberHandler.GetMembers)
		members.PUT("/:userID", groupMemberHandler.UpdateMemberRole)
		members.DELETE("/:userID", groupMemberHandler.RemoveMember)
	}

	settlementsRepo := repository.NewSettlementRepository(db)
	settlementsSrv := service.NewSettlementService(settlementsRepo)
	settlementsHandler := controllers.NewSettlementHandler(settlementsSrv)

	settlements := groupRoutes.Group("/:id/settlements", middlewares.CheckUserAuthentication)
	{
		settlements.POST("/", settlementsHandler.Create)
		settlements.POST("/", settlementsHandler.GetAll)
		settlements.PUT("/:id", settlementsHandler.Update)
		settlements.DELETE("/:id", settlementsHandler.Delete)
	}


	settlementsRepo := repository.NewSettlementRepository(db)
	settlementsSrv := service.NewSettlementService(settlementsRepo)
	settlementsHandler := controllers.NewSettlementHandler(settlementsSrv)

	settlements := groupRoutes.Group("/:id/settlements", middlewares.CheckUserAuthentication)
	{
		settlements.POST("/", settlementsHandler.Create)
		settlements.GET("/", settlementsHandler.GetAll)
		settlements.PUT("/:id", settlementsHandler.Update)
		settlements.DELETE("/:id", settlementsHandler.Delete)
	}

	return r
}
