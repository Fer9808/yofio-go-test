package test

import (
	"errors"
	"testing"

	"github.com/Fer9808/yofio-go-test/internal/api/services"
	"github.com/Fer9808/yofio-go-test/internal/pkg/models"
	"github.com/Fer9808/yofio-go-test/internal/pkg/persistence"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock del repositorio
type MockAssignmentsRepository struct {
	mock.Mock
}

func (m *MockAssignmentsRepository) Add(assignment *models.Assignments) error {
	args := m.Called(assignment)
	return args.Error(0)
}

func (m *MockAssignmentsRepository) All() (persistence.StatisticsResponse, error) {
	args := m.Called()
	return args.Get(0).(persistence.StatisticsResponse), args.Error(1)
}

func TestAssign_Success(t *testing.T) {
	mockRepo := new(MockAssignmentsRepository)
	creditService := services.NewCreditServiceImpl(mockRepo)

	mockRepo.On("Add", mock.AnythingOfType("*models.Assignments")).Return(nil)

	investment := int32(3000)
	_, _, _, err := creditService.Assign(investment)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAssign_Failure(t *testing.T) {
	mockRepo := new(MockAssignmentsRepository)
	creditService := services.NewCreditServiceImpl(mockRepo)

	mockRepo.On("Add", mock.AnythingOfType("*models.Assignments")).Return(errors.New("db error"))

	investment := int32(300)
	_, _, _, err := creditService.Assign(investment)

	assert.Nil(t, err, "db error")
	mockRepo.AssertExpectations(t)
}

func TestAssign_NoValidCombination(t *testing.T) {
	mockRepo := new(MockAssignmentsRepository)
	service := services.NewCreditServiceImpl(mockRepo)

	mockRepo.On("Add", mock.AnythingOfType("*models.Assignments")).Return(nil)

	// Intentar asignar una cantidad que sabemos no puede ser combinada con éxito
	// Ejemplo: Un valor que no es divisible ni por 300, 500, ni 700
	investment := int32(1) // Este valor debe garantizar que no se encuentre una combinación
	credit300, credit500, credit700, err := service.Assign(investment)

	assert.Equal(t, int32(0), credit300, "No debería asignarse créditos de $300")
	assert.Equal(t, int32(0), credit500, "No debería asignarse créditos de $500")
	assert.Equal(t, int32(0), credit700, "No debería asignarse créditos de $700")
	assert.Error(t, err, "Debería retornar un error indicando que no se pudo asignar crédito")
	mockRepo.AssertExpectations(t)
}

func TestAllAssignments_Success(t *testing.T) {
	mockRepo := new(MockAssignmentsRepository)
	expectedStats := persistence.StatisticsResponse{
		TotalAssignments:              10,
		SuccessfulAssignments:         8,
		UsuccessfulAssignments:        2,
		AverageSuccessfulInvestment:   500,
		AverageUnsuccessfulInvestment: 200,
	}

	// Configura el mock para devolver las estadísticas esperadas sin error
	mockRepo.On("All").Return(expectedStats, nil)

	stats, err := mockRepo.All()

	assert.NoError(t, err)
	assert.Equal(t, expectedStats, stats)
	mockRepo.AssertExpectations(t) // Verificar que se cumplieron las expectativas del mock
}
