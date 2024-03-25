package test

import (
	"testing"

	"github.com/Fer9808/yofio-go-test/internal/pkg/config"
	"github.com/Fer9808/yofio-go-test/internal/pkg/db"
	"github.com/Fer9808/yofio-go-test/internal/pkg/models"
	"github.com/Fer9808/yofio-go-test/internal/pkg/persistence"

	"github.com/stretchr/testify/assert"
)

func Setup() {
	config.Setup("./config.yml")
	db.SetupDB()
	db.GetDB().Exec("DELETE FROM assignments")
	db.GetDB().Exec("INSERT INTO assignments(id, investment, credit_300, credit_500, credit_700, success) VALUES(1, 500, 0, 1, 0, true);")
	db.GetDB().Exec("INSERT INTO assignments(id, investment, credit_300, credit_500, credit_700, success) VALUES(2, 400, 0, 0, 0, false);")
}

func TestAddAssignment(t *testing.T) {
	Setup()
	repo := persistence.GetAssignmentsRepository()

	assignment := models.Assignments{
		Investment:    1500,
		CreditType300: 5,
		CreditType500: 0,
		CreditType700: 0,
		Success:       true,
	}

	err := repo.Add(&assignment)
	assert.NoError(t, err, "No debería haber error al añadir una nueva asignación")

	var count int64
	db.GetDB().Model(&models.Assignments{}).Count(&count)
	assert.Equal(t, int64(3), count, "Debería haber exactamente tres asignaciones en la base de datos")
}

func TestAllAssignments(t *testing.T) {
	Setup()
	repo := persistence.GetAssignmentsRepository()

	response, err := repo.All()
	assert.NoError(t, err, "No debería haber error al recuperar estadísticas")

	assert.GreaterOrEqual(t, response.TotalAssignments, int64(2), "Debería haber dos asignaciones en total")
	assert.GreaterOrEqual(t, response.SuccessfulAssignments, int64(1), "Debería haber un asignación exitosas")
	assert.GreaterOrEqual(t, response.UsuccessfulAssignments, int64(1), "Debería haber una asignación no exitosas")
}
