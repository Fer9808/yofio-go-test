package models

type Assignments struct {
	Investment    int32 `gorm:"column:investment;not null;" json:"investment" form:"investment"`
	CreditType300 int32 `gorm:"column:credit_300;" json:"credit_300" form:"credit_300"`
	CreditType500 int32 `gorm:"column:credit_500;" json:"credit_500" form:"credit_500"`
	CreditType700 int32 `gorm:"column:credit_700;" json:"credit_700" form:"credit_700"`
	Success       bool  `gorm:"column:success;" json:"success" form:"success"`
}
