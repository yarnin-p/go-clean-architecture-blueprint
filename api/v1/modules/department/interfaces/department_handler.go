package interfaces

import "github.com/gin-gonic/gin"

type DepartmentHandler interface {
	GetAllDepartmentsHandler(c *gin.Context)
}
