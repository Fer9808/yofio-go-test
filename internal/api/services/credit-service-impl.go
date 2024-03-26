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
	max700s := investment / 700
	remaining := investment % 700

	max500s := remaining / 500
	remaining = remaining % 500

	if remaining%300 == 0 {
		max300s := remaining / 300

		// Registrar asignación exitosa
		assignment := models.Assignments{
			Investment:    investment,
			CreditType300: max300s,
			CreditType500: max500s,
			CreditType700: max700s,
			Success:       true,
		}

		if err := c.repo.Add(&assignment); err != nil {
			log.Println(err)
		}

		return max300s, max500s, max700s, nil
	}

	for ; max700s >= 0; max700s-- {
		for try500s := int32(0); try500s <= max500s; try500s++ {
			remainingAfter700sAnd500s := investment - max700s*700 - try500s*500
			if remainingAfter700sAnd500s%300 == 0 {
				try300 := remainingAfter700sAnd500s / 300
				// Registrar asignación exitosa
				assignment := models.Assignments{
					Investment:    investment,
					CreditType300: try300,
					CreditType500: try500s,
					CreditType700: max700s,
					Success:       true,
				}

				if err := c.repo.Add(&assignment); err != nil {
					log.Println(err)
				}

				return try300, try500s, max700s, nil
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
