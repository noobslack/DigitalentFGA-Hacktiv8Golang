package controller

import (
	"assignment2/model"
	"assignment2/repository"
	"assignment2/util"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	ordersRepository repository.IOrderRepository
}

func NewOrderController(ordersRepository repository.IOrderRepository) *OrderController {

	return &OrderController{
		ordersRepository: ordersRepository,
	}

}

func (oc *OrderController) CreateOrder(ctx *gin.Context) {
	var newOrder model.Order

	err := ctx.ShouldBindJSON(&newOrder)

	if err != nil {
		var r model.Response = model.Response{
			Success: false,
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	createdOrder, err := oc.ordersRepository.CreateOrder(newOrder)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, createdOrder, ""))

}

func (oc *OrderController) GetOrder(ctx *gin.Context) {

	orders, err := oc.ordersRepository.GetOrder()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	fmt.Println(orders)
	ctx.JSON(http.StatusOK, util.CreateResponse(true, orders, ""))
}

func (oc *OrderController) UpdateOrder(ctx *gin.Context) {

	OrderID, err := strconv.Atoi(ctx.Param("order_id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.CreateResponse(false, nil, err.Error()))
		return
	}

	var updateOrder model.Order

	err = ctx.ShouldBindJSON(&updateOrder)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.CreateResponse(false, nil, err.Error()))
		return
	}

	updateOrder.Order_id = OrderID

	updateOrder, err = oc.ordersRepository.UpdateOrder(updateOrder)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, updateOrder, ""))

}

func (oc *OrderController) DeleteOrder(ctx *gin.Context) {
	var deleteOrder model.Order
	var err error
	orderID := ctx.Param("order_id")

	deleteOrder.Order_id, err = strconv.Atoi(orderID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	deleteOrder, err = oc.ordersRepository.DeleteOrder(deleteOrder.Order_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, deleteOrder, ""))

}
