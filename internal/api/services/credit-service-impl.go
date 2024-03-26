package services

import (
	"errors"
	"log"

	models "github.com/Fer9808/yofio-go-test/internal/pkg/models"
	"github.com/Fer9808/yofio-go-test/internal/pkg/persistence"
)

type CreditAssigner interface {
	Assign(investment int32) (int32, int32, int32, error)
}

type CreditServiceImpl struct {
	repo persistence.AssignmentsRepositorInterface
}

func NewCreditServiceImpl(repo persistence.AssignmentsRepositorInterface) *CreditServiceImpl {
	return &CreditServiceImpl{repo: repo}
}

func (c CreditServiceImpl) Assign(investment int32) (int32, int32, int32, error) {
	// Se realiza la iteración para asignar todos los posibles créditos de $700
	for z := investment / 700; z >= 0; z-- {
		remainingAfter700s := investment - z*700
		// Se realiza la iteración para asignar todos los posibles créditos de $500 con el restante
		for y := remainingAfter700s / 500; y >= 0; y-- {
			remainingAfter500s := remainingAfter700s - y*500
			// Se realiza la iteración para asignar todos los créditos de $300 con el restante
			if remainingAfter500s%300 == 0 {
				x := remainingAfter500s / 300

				// Registrar asignación exitosa
				assignment := models.Assignments{
					Investment:    investment,
					CreditType300: x,
					CreditType500: y,
					CreditType700: z,
					Success:       true,
				}

				if err := c.repo.Add(&assignment); err != nil {
					log.Println(err)
				}

				return x, y, z, nil
			}
		}
	}

	// Registrar intento fallido
	assignment := models.Assignments{
		Investment: investment,
		Success:    false,
	}

	if err := c.repo.Add(&assignment); err != nil {
		log.Println(err)
	}

	return 0, 0, 0, errors.New("Error al asignar crédito")
}
