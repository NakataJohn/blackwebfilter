package controller

import (
	"blackwebfilter/domain"
	"blackwebfilter/internal/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FilterController struct {
	Dc domain.DomainFilterUsecase
	Wc domain.WordFilterUsecase
}

// @Summary Comprehensive evaluation for URL
// @Tags filter
// @Description 综合验证url是否是恶意连接，先过域名黑名单后检测敏感词；
// @Produce application/json
// @Param data body domain.FilterRequest true "URL"
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Router /api/v1/filter [post]
func (fc *FilterController) Verify(c *gin.Context) {
	var request domain.DomainFilterRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	d_resv, err := fc.Dc.Verify(c, request.URL)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	// 不在域名黑名单中则检测敏感词
	if d_resv == "Not found" {
		s_resv, err := fc.Wc.Verify(c, request.URL)
		if err != nil {
			logger.Error(err.Error())
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "无法检测敏感词," + err.Error()})
			return
		}
		// 敏感词不存时
		if s_resv == d_resv {
			c.JSON(http.StatusOK, domain.SuccessResponse{Message: d_resv})
		} else {
			c.JSON(http.StatusOK, domain.SuccessResponse{Message: s_resv})
		}
	} else {
		c.JSON(http.StatusOK, domain.SuccessResponse{Message: d_resv})
	}
}
