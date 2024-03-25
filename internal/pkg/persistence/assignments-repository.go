package persistence

import (
	"github.com/Fer9808/yofio-go-test/internal/pkg/db"
	models "github.com/Fer9808/yofio-go-test/internal/pkg/models"
)

type AssignmentsRepository struct{}

var assignmentsRepository *AssignmentsRepository

type StatisticsResponse struct {
	TotalAssignments              int64   `json:"total_assignments"`
	SuccessfulAssignments         int64   `json:"successful_assignments"`
	UsuccessfulAssignments        int64   `json:"unsuccessful_assignments"`
	AverageSuccessfulInvestment   float64 `json:"average_successful_investment"`
	AverageUnsuccessfulInvestment float64 `json:"average_unsuccessful_investment"`
}

func GetAssignmentsRepository() *AssignmentsRepository {
	if assignmentsRepository == nil {
		assignmentsRepository = &AssignmentsRepository{}
	}
	return assignmentsRepository
}

func (r *AssignmentsRepository) Add(assignments *models.Assignments) error {
	err := db.GetDB().Save(&assignments).Error
	return err
}

func (r *AssignmentsRepository) All() (StatisticsResponse, error) {
	var total, successfulCount, unsuccessfulCount int64
	var successfulInvestmentTotal, unsuccessfulInvestmentTotal float64
	err := db.GetDB().Model(&models.Assignments{}).Count(&total).Error
	err = db.GetDB().Model(&models.Assignments{}).Where("success = ?", true).Count(&successfulCount).Error
	err = db.GetDB().Model(&models.Assignments{}).Where("success = ?", false).Count(&unsuccessfulCount).Error
	db.GetDB().Table("assignments").Select("AVG(investment)").Where("success = ?", true).Row().Scan(&successfulInvestmentTotal)
	db.GetDB().Table("assignments").Select("AVG(investment)").Where("success = ?", false).Row().Scan(&unsuccessfulInvestmentTotal)
	return StatisticsResponse{
		TotalAssignments:              total,
		SuccessfulAssignments:         successfulCount,
		UsuccessfulAssignments:        unsuccessfulCount,
		AverageSuccessfulInvestment:   successfulInvestmentTotal,
		AverageUnsuccessfulInvestment: unsuccessfulInvestmentTotal,
	}, err
}
