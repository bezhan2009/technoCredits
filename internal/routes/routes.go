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

	// Auth
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
	// Protected routes (require authentication)
	api := r.Group("")
	api.Use(middlewares.CheckUserAuthentication)
	{
		// Cards expense routes
		cards := api.Group("/cards")
		{
			cards.GET("", controllers.GetAllCardsUser)
			cards.POST("", controllers.CreateCardExpense)
			cards.PATCH("/:id", controllers.UpdateCardExpense)
			cards.DELETE("/:id", controllers.DeleteCardExpense)
		}

		// Card expense payers routes
		payers := api.Group("/cards/payers")
		{
			payers.POST("", controllers.CreateCardExpensePayer)
			payers.PATCH("/:id", controllers.UpdateCardExpensePayer)
			payers.DELETE("/:id", controllers.DeleteCardExpensePayer)
		}

		// Card expense users routes
		users := api.Group("/cards/users")
		{
			users.POST("", controllers.CreateCardExpenseUser)
			users.PATCH("/:id", controllers.UpdateCardExpenseUser)
			users.DELETE("/:id", controllers.DeleteCardExpenseUser)
		}
	}

	return r
}
