package bootstrap

import (
	"blackwebfilter/internal/fileutil"
	"blackwebfilter/internal/logger"
	"bufio"
	"fmt"
	"os"

	"github.com/bits-and-blooms/bloom/v3"
	"go.uber.org/zap"
)

func NewBloomFilter(env *Env) (*bloom.BloomFilter, error) {
	filter := bloom.NewWithEstimates(10000000, 0.001)
	if fileutil.Exists(env.BlackDomainsPath) {
		file, err := os.Open(env.BlackDomainsPath)
		if err != nil {
			logger.Error("Environment of BLACK_DOMAINS_PATH can't be loaded: ", zap.Error(err))
			return nil, err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			filter.Add(scanner.Bytes())
		}
		return filter, nil
	} else {
		return nil, fmt.Errorf("file of blackweb list is not exists")
	}
}

func FilterTest(data string) bool {
	app := App()
	env := app.Env
	filter, err := NewBloomFilter(env)
	if err != nil {
		panic(err)
	}
	return filter.Test([]byte(data))
}
