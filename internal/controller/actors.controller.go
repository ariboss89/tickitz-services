package controller

import (
	"net/http"

	"github.com/ariboss89/tickitz-services/internal/dto"
	"github.com/ariboss89/tickitz-services/internal/service"
	"github.com/gin-gonic/gin"
)

type ActorController struct {
	actorService *service.ActorService
}

func NewActorController(actorService *service.ActorService) *ActorController {
	return &ActorController{
		actorService: actorService,
	}
}

// Get All Actors godoc
// @Summary      Get All Actors
// @Tags         actor
// @Produce      json
// @Success      200  {object}  dto.Actors
// @Failure 		 500 {object} dto.ResponseError
// @Router       /actors/ [get]
func (m ActorController) GetAllActors(c *gin.Context) {
	data, err := m.actorService.GetAllActors(c.Request.Context())
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
