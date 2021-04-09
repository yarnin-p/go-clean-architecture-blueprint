package view_models

type DepartmentOut struct {
	QId      string `json:"q_id" gorm:"column:qID;type:char(6)"`
	QName    string `json:"q_name" gorm:"column:qName;type:char(25)"`
	DeptCode string `json:"dept_code" gorm:"column:deptCode;type:char(6);not null"`
	DeptDesc string `json:"dept_desc" gorm:"column:deptDesc;type:char(35)"`
}
