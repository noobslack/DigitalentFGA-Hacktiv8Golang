package main

import (
	"assignment2/controller"
	"assignment2/lib"
	"assignment2/repository"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := lib.InitDatabase()
	if err != nil {
		panic(err)
	}

	orderRepository := repository.NewOrderRepository(db)
	orderController := controller.NewOrderController(orderRepository)

	ginEngine := gin.Default()

	ginEngine.GET("/order", orderController.GetOrder)
	ginEngine.POST("/order", orderController.CreateOrder)
	ginEngine.PUT("/order/:order_id", orderController.UpdateOrder)
	ginEngine.DELETE("/order/:order_id", orderController.DeleteOrder)

	err = ginEngine.Run("localhost:8080")

	if err != nil {
		panic(err)
	}
}
