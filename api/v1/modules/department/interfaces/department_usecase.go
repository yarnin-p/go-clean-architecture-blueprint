package interfaces

import (
	"context"
	"go-clean-architecture-blueprint/api/v1/models"
)

type DepartmentUseCase interface {
	GetAllDepartmentsService(ctx context.Context, departments *[]models.Department) (err error)
}
