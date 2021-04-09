package routes

import (
	"github.com/gin-gonic/gin"
	"go-clean-architecture-blueprint/api/v1/modules/department/handlers"
	"go-clean-architecture-blueprint/api/v1/modules/department/interfaces"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc interfaces.DepartmentUseCase) {
	h := handlers.NewDepartmentHandler(uc)
	r := router.Group("/departments")
	{
		r.GET("/", h.GetAllDepartmentsHandler)
	}
}
