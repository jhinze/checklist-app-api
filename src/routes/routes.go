package routes

import (
	"checklist/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//Health check
	r.GET("/health", controllers.GetHealth)

	//API group
	v1 := r.Group("/v1")
	{
		//List
		v1.GET("list", controllers.RetrieveLists)
		v1.GET("list/:listId", controllers.RetrieveList)
		v1.POST("list", controllers.CreateList)
		v1.DELETE("list/:listId", controllers.DeleteList)
		v1.PUT("list/:listId", controllers.UpdateList)

		//Item
		v1.GET("list/:listId/item/", controllers.RetrieveItems)
		v1.GET("list/:listId/item/:itemId", controllers.RetrieveItem)
		v1.POST("list/:listId/item", controllers.CreateItem)
		v1.DELETE("list/:listId/item/:itemId", controllers.DeleteItem)
		v1.PUT("/list/:listId/item/:itemId", controllers.UpdateItem)
		v1.PUT("/list/:listId/item/:itemId/complete", controllers.CompleteItem)
		v1.PUT("/list/:listId/item/:itemId/incomplete", controllers.IncompleteItem)

		//Util
		v1.GET("list/:listId/itemsSummary", controllers.RetrieveItemsSummary)
	}

	return r
}
