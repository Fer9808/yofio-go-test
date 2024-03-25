package services

import (
	"errors"
	"log"

	models "github.com/Fer9808/yofio-go-test/internal/pkg/models"
	"github.com/Fer9808/yofio-go-test/internal/pkg/persistence"
)

type CreditServiceImpl struct{}

func NewCreditServiceImpl() *CreditServiceImpl {
	return &CreditServiceImpl{}
}

func (c CreditServiceImpl) Assign(investment int32) (int32, int32, int32, error) {
	s := persistence.GetAssignmentsRepository()
	for x := investment / 300; x >= 0; x-- {
		for y := (investment - 300*x) / 500; y >= 0; y-- {
			for z := (investment - 300*x - 500*y) / 700; z >= 0; z-- {
				if 300*x+500*y+700*z == investment {
					// Registrar asignación exitosa
					assignment := models.Assignments{
						Investment:    investment,
						CreditType300: x,
						CreditType500: y,
						CreditType700: z,
						Success:       true,
					}

					if err := s.Add(&assignment); err != nil {
						log.Println(err)
					}

					return x, y, z, nil
				}
			}
		}
	}
	// Registrar intento fallido
	assignment := models.Assignments{
		Investment: investment,
		Success:    false,
	}

	if err := s.Add(&assignment); err != nil {
		log.Println(err)
	}

	return 0, 0, 0, errors.New("Error al asignar crédito")
}