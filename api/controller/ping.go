package controller

import (
	"blackwebfilter/bootstrap"
	"blackwebfilter/domain"
	"blackwebfilter/internal/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingController struct {
	PingUsecase domain.PingUsecase
	Env         *bootstrap.Env
}

// @Summary Ping
// @Tags Ping
// @Description api探活
// @Success 200 {object} domain.SuccessResponse
// @Failture 500 {object} domain.ErrorResponse
// @Router /api/v1/ping [get]
func (pc *PingController) Ping(c *gin.Context) {
	resv, err := pc.PingUsecase.Ping(c)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: resv})
}
