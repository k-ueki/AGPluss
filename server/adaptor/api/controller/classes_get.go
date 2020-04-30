package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/k-ueki/AGPlus/server/adaptor/repository"

	"github.com/gin-gonic/gin"
	"github.com/k-ueki/AGPlus/server/application/service"
)

type (
	ClassGetController struct {
		service.ClassGetService
	}
)

func NewClassGetController(db *gorm.DB) *ClassGetController {
	return &ClassGetController{
		ClassGetService: service.ClassGetService{
			ClassGetRepository: repository.ClassGetRepository{
				DB: db,
			},
		},
	}
}

func (c *ClassGetController) List(ctx *gin.Context) {
	classes, err := c.ClassGetService.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, classes)
	return
}

func (c *ClassGetController) Show(ctx *gin.Context) {
	//TODO: structでbindしたい
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors.New("failed to bind parameters"))
		return
	}

	class, err := c.ClassGetService.Show(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, class)
	return
}