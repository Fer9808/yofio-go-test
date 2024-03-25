package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/Fer9808/yofio-go-test/internal/api/services"
	"github.com/Fer9808/yofio-go-test/internal/pkg/persistence"
	"github.com/Fer9808/yofio-go-test/pkg/http-err"

	"github.com/gin-gonic/gin"
)

type InvestmentReq struct {
	Investment int32 `json:"investment" binding:"required"`
}

type InvestmentRep struct {
	CreditType300 int32 `json:"credit_type_300"`
	CreditType500 int32 `json:"credit_type_500"`
	CreditType700 int32 `json:"credit_type_700"`
}

// CreateAssignments godoc
// @Summary Realiza el proceso de asignación de un crédito
// @Description Create Assigments
// @Produce json
// @Param Investment
// @Success 200 {object} InvestmentRep
// @Failure 400  {object}  httputil.HTTPError
// @Router /api/credit-assignment [post]
func CreateAssignments(c *gin.Context) {
	var investmentReq InvestmentReq
	creditService := services.NewCreditServiceImpl()
	_ = c.BindJSON(&investmentReq)

	x, y, z, err := creditService.Assign(investmentReq.Investment)
	if err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
	} else {
		resp := InvestmentRep{
			CreditType300: x,
			CreditType500: y,
			CreditType700: z,
		}
		c.JSON(http.StatusOK, resp)
	}
}

// GetStatistics godoc
// @Summary Realiza el proceso de obtener estadisticas de la asignación de créditos
// @Description Get Statistics
// @Produce json
// @Success 200 {object} Statistics
// @Failure 400  {object}  httputil.HTTPError
// @Router /api/statistics [post]
func GetStatistics(c *gin.Context) {
	s := persistence.GetAssignmentsRepository()
	if statistics, err := s.All(); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("Error al obtener las estadisticas"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, statistics)
	}
}
