package routers

import (
	controllers "sema/Delivery/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(UserCtrl *controllers.UserController, ChtCtrl *controllers.ChatController) {
	router := gin.Default()

	// User routes 
	userRoutes := router.Group("/user")
	{
		userRoutes.GET("")
	}

	// Chat routes 
	chatRoutes := router.Group("/chat") 
	{
		chatRoutes.GET("")
	}
	
	// Run the router
	router.Run();
}