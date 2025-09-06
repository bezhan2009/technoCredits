package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "technoCredits/docs"
	"technoCredits/internal/controllers"
	"technoCredits/internal/controllers/middlewares"
	"technoCredits/internal/controllers/websockets"
)

func InitRoutes(r *gin.Engine) *gin.Engine {

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

	r.GET("/users", controllers.GetAllUsers)

	profile := r.Group("/profile", middlewares.CheckUserAuthentication)
	{
		profile.GET("", controllers.GetMyDataUser)
		profile.PATCH("", controllers.UpdateUser)
		profile.PATCH("/password", controllers.UpdateUsersPassword)
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

		users := api.Group("/cards/users")
		{
			users.POST("", controllers.CreateCardExpenseUser)
			users.PATCH("/:id", controllers.UpdateCardExpenseUser)
			users.DELETE("/:id", controllers.DeleteCardExpenseUser)
		}

		// Groups routes
		groups := api.Group("/groups")
		{
			groups.GET("", controllers.GetGroupsUser)
			groups.POST("", controllers.CreateGroup)

			// Отдельная группа для операций с конкретной группой
			group := groups.Group("/:group_id")
			{
				group.GET("", controllers.GetGroupByID)
				group.PUT("", controllers.UpdateGroup)
				group.DELETE("", controllers.DeleteGroup)

				// Участники конкретной группы
				members := group.Group("/members")
				{
					members.GET("", controllers.GetGroupMembersByGroupID)
					members.POST("", controllers.CreateGroupMember)
				}
			}
		}

		// Отдельные операции с участниками (по ID участника)
		groupMembers := api.Group("/group-members")
		{
			groupMembers.PUT("/:id", controllers.UpdateGroupMember)
			groupMembers.DELETE("/:id", controllers.DeleteGroupMember)
		}

		// Settlements routes
		settlements := api.Group("/settlements")
		{
			settlements.GET("", controllers.GetAllSettlements)
			settlements.POST("", controllers.CreateSettlement)
			settlements.GET("/:id", controllers.GetSettlementByID)
			settlements.PUT("/:id", controllers.UpdateSettlement)
			settlements.DELETE("/:id", controllers.DeleteSettlement)
		}
	}

	r.GET("/notifications", middlewares.CheckUserAuthentication, websockets.RealTimeNotificationReader)

	return r
}
