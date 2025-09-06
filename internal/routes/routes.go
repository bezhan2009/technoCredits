package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "technoCredits/docs"
	"technoCredits/internal/controllers"
	"technoCredits/internal/controllers/middlewares"
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

	profile := r.Group("/profile", middlewares.CheckUserAuthentication)
	{
		profile.GET("", controllers.GetMyDataUser)
		profile.PATCH("", controllers.UpdateUser)
		profile.PATCH("/password", controllers.UpdateUsersPassword)
	}
	//
	//groupRoutes := r.Group("/groups", middlewares.CheckUserAuthentication)
	//{
	//	groupRoutes.POST("/", controllers.CreateGroup)
	//	groupRoutes.GET("/:id", controllers.GetGroupByID)
	//	groupRoutes.PUT("/:id", controllers.UpdateGroup)
	//	groupRoutes.DELETE("/:id", controllers.DeleteGroup)
	//}

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
