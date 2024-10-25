package usecase

import (
	"blackwebfilter/domain"
	"blackwebfilter/internal/logger"
	"blackwebfilter/internal/parseutil"
	"context"
	"time"

	"github.com/bits-and-blooms/bloom/v3"
)

type domainFilterUsecase struct {
	filter         *bloom.BloomFilter
	contextTimeout time.Duration
}

func NewDomainFilterUsecase(timeout time.Duration, bloomfilter *bloom.BloomFilter) domain.DomainFilterUsecase {
	return &domainFilterUsecase{
		filter:         bloomfilter,
		contextTimeout: timeout,
	}
}

func (du *domainFilterUsecase) Verify(c context.Context, urlStr string) (string, error) {
	_, cancel := context.WithTimeout(c, du.contextTimeout)
	defer cancel()
	subdomains := parseutil.GetSubDomains(urlStr)

	for _, domainStr := range subdomains {
		if du.filter.Test([]byte(domainStr)) {
			logger.Infof("URL %s Matched Black Domain \"%s\" \n", urlStr, domainStr)
			return "Bingo! Black domain found❀❀❀ " + domainStr, nil
		}
	}

	logger.Infof("URL %s Have Not Matched Any Black Domains.\n", urlStr)
	return "Not found", nil
}
