package models

type Department struct {
	QId      string `json:"q_id" gorm:"column:qID;type:char(6)"`
	QName    string `json:"q_name" gorm:"column:qName;type:char(25)"`
	DeptCode string `json:"dept_code" gorm:"column:deptCode;type:char(6);not null"`
	DeptDesc string `json:"dept_desc" gorm:"column:deptDesc;type:char(35)"`
}

// Override table name
func (Department) TableName() string {
	return "Touch_department"
}
