package controller

import (
	"blackwebfilter/bootstrap"
	"blackwebfilter/domain"
	"blackwebfilter/internal/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WordFilterController struct {
	WordFilterUsecase domain.WordFilterUsecase
	Env               *bootstrap.Env
}

// @Summary Sensitive Word Verify for URL
// @Tags filter
// @Description 验证url的文本内容是否存在敏感词。
// @Produce application/json
// @Param data body domain.WordFilterRequest true "URL"
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Router /api/v1/sensitive [post]
func (wc *WordFilterController) Verify(c *gin.Context) {
	var request domain.WordFilterRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	resv, err := wc.WordFilterUsecase.Verify(c, request.URL)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "无法检测敏感词," + err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: resv})
}
