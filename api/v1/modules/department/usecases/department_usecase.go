package usecases

import (
	"context"
	"go-clean-architecture-blueprint/api/v1/models"
	"go-clean-architecture-blueprint/api/v1/modules/department/interfaces"
)

type departmentUseCase struct {
	departmentRepo interfaces.DepartmentRepository
}

func NewDepartmentUseCase(appointRepo interfaces.DepartmentRepository) *departmentUseCase {
	return &departmentUseCase{
		departmentRepo: appointRepo,
	}
}

func (u *departmentUseCase) GetAllDepartmentsService(ctx context.Context, departments *[]models.Department) (err error) {
	if err = u.departmentRepo.GetAllDepartments(ctx, departments); err != nil {
		return err
	}
	return nil
}
