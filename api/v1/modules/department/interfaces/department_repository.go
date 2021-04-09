package interfaces

import (
	"context"
	"go-clean-architecture-blueprint/api/v1/models"
)

type DepartmentRepository interface {
	GetAllDepartments(ctx context.Context, departments *[]models.Department) (err error)
}
