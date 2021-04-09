package repositories

import (
	"context"
	"go-clean-architecture-blueprint/api/v1/models"
	"gorm.io/gorm"
)

type departmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) *departmentRepository {
	return &departmentRepository{
		db: db,
	}
}

func (r *departmentRepository) GetAllDepartments(ctx context.Context, departments *[]models.Department) (err error) {
	if err = r.db.WithContext(ctx).Model(departments).Find(departments).Error; err != nil {
		return err
	}
	return nil
}
