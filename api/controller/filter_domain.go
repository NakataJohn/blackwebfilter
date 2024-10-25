package controller

import (
	"blackwebfilter/bootstrap"
	"blackwebfilter/domain"
	"blackwebfilter/internal/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DomainFilterController struct {
	DomainFilterUsecase domain.DomainFilterUsecase
	Env                 *bootstrap.Env
}

// @Summary Black Domain Verify for URL
// @Tags filter
// @Description 验证url的域名是否在黑名单中
// @Produce application/json
// @Param data body domain.DomainFilterRequest true "URL"
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Router /api/v1/domains [post]
func (dc *DomainFilterController) Verify(c *gin.Context) {
	var request domain.DomainFilterRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	resv, err := dc.DomainFilterUsecase.Verify(c, request.URL)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: resv})
}
