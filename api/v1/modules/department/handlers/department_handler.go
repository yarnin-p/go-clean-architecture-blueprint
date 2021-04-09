package handlers

import (
	"github.com/gin-gonic/gin"
	"go-clean-architecture-blueprint/api/v1/models"
	"go-clean-architecture-blueprint/api/v1/modules/department/interfaces"
	"go-clean-architecture-blueprint/utils"
	"net/http"
)

type departmentHandler struct {
	departmentUseCase interfaces.DepartmentUseCase
}

func NewDepartmentHandler(departmentUseCase interfaces.DepartmentUseCase) *departmentHandler {
	return &departmentHandler{
		departmentUseCase: departmentUseCase,
	}
}

func (h *departmentHandler) GetAllDepartmentsHandler(c *gin.Context) {
	var departments []models.Department
	err := h.departmentUseCase.GetAllDepartmentsService(c.Request.Context(), &departments)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "", err.Error(), make([]models.Department, 0))
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "", "", departments)
	return
}
