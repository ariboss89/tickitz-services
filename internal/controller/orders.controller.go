package controller

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/ariboss89/tickitz-services/internal/dto"
	"github.com/ariboss89/tickitz-services/internal/service"
	"github.com/ariboss89/tickitz-services/pkg"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService *service.OrderService
}

func NewOrderController(orderService *service.OrderService) *OrderController {
	return &OrderController{
		orderService: orderService,
	}
}

// Get Schedules godoc
// @Summary      Get Schedules
// @Tags         order
// @Produce      json
// @Param        date		query string  true  "Movie Date"
// @Param        time		query string  true  "Movie Time"
// @Param        location		query string  true  "Cinema City"
// @Param        movie_id		query string  true  "Movie Id"
// @Success      200  {object}  dto.Schedules
// @Failure 		 500 {object} dto.ResponseError
// @Failure			 404 {object} dto.ResponseError
// @Router       /order/schedule [get]
// @Security			BearerAuth
func (o OrderController) GetSchedules(c *gin.Context) {
	date := c.Query("date")
	time := c.Query("time")
	location := c.Query("location")
	movieId := c.Query("movie_id")
	movieIdInt, _ := strconv.Atoi(movieId)

	var ordersQuery dto.OrdersQuery
	if err := c.ShouldBindQuery(&ordersQuery); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "internal server error",
		})
		return
	}

	data, err := o.orderService.GetSchedules(c.Request.Context(), date, time, location, movieIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "OK",
		"data": data,
	})
}

// Create Order godoc
// @Summary      Create Order
// @Tags         order
// @accept			 json
// @Produce      json
// @Param        auth	 body dto.CreateOrder  true  "Create Order Body"
// @Success      200  {object}  dto.CreateOrder
// @Failure 		 500 {object} dto.ResponseError
// @Failure			 404 {object} dto.ResponseError
// @Failure			 400 {object} dto.ResponseError
// @Router       /order/ [post]
// @Security			BearerAuth
func (o OrderController) CreateOrder(c *gin.Context) {
	// seats := c.QueryArray("seat")
	token, isExist := c.Get("token")
	if !isExist {
		c.AbortWithStatusJSON(http.StatusForbidden, dto.Response{
			Msg:     "Forbidden Access",
			Success: false,
			Data:    []any{},
			Error:   "Access Denied",
		})
		return
	}

	accessToken, _ := token.(pkg.JWTClaims)

	var newOrder dto.CreateOrder

	if err := c.ShouldBindJSON(&newOrder); err != nil {
		str := err.Error()
		if strings.Contains(str, "Field") {
			c.JSON(http.StatusBadRequest, dto.ResponseError{
				Msg:     "Invalid Body",
				Success: false,
				Error:   "invalid body",
			})
			return
		}
		if strings.Contains(str, "Empty") {
			c.JSON(http.StatusBadRequest, dto.ResponseError{
				Msg:     "Invalid Body",
				Success: false,
				Error:   "invalid body",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "internal server error",
		})
		return
	}

	userId := accessToken.Id

	if _, err := o.orderService.CreateOrder(c.Request.Context(), userId, newOrder); err != nil {
		str := err.Error()
		if strings.Contains(str, "empty") {
			c.JSON(http.StatusBadRequest, dto.ResponseError{
				Msg:     "Invalid Body",
				Success: false,
				Error:   str,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "Order created successfully",
		"data": []any{},
	})
}
