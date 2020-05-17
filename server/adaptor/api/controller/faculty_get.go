package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/k-ueki/AGPlus/server/adaptor/repository"
	"github.com/k-ueki/AGPlus/server/application/service"
)

type (
	FacultyGetController struct {
		service.FacultyGetService
	}
)

func NewFacultyGetController(db *gorm.DB) *FacultyGetController {
	return &FacultyGetController{
		FacultyGetService: service.FacultyGetService{
			FacultyGetRepository: repository.FacultyGetRepository{
				DB: db,
			},
		},
	}
}

func (c *FacultyGetController) List(ctx *gin.Context) {

}
