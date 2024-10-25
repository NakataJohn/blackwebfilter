package usecase

import (
	"blackwebfilter/domain"
	"blackwebfilter/internal/logger"
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	ahocorasick "github.com/BobuSumisu/aho-corasick"
	"github.com/PuerkitoBio/goquery"
)

type wordFilterUsecase struct {
	trie           *ahocorasick.Trie
	contextTimeout time.Duration
}

func NewSensitiveWordUsecase(timeout time.Duration, trie *ahocorasick.Trie) domain.WordFilterUsecase {
	return &wordFilterUsecase{
		trie:           trie,
		contextTimeout: timeout,
	}
}

func (wu *wordFilterUsecase) Verify(c context.Context, urlStr string) (string, error) {
	_, cancel := context.WithTimeout(c, wu.contextTimeout)
	defer cancel()

	texts, err := getDocTexts(urlStr)
	if err != nil {
		return "", err
	}
	logger.Infof("%s body texts:%s", urlStr, texts)
	matches := wu.trie.MatchString(texts)

	if len(matches) > 0 {
		for _, match := range matches {
			logger.Infof("URL %s Matched Word \"%s\" at position %d.\n", urlStr, string(match.Match()), match.Pos())
		}
		return "Bingo! Sensitive words found❀❀❀" + strconv.Itoa(len(matches)) + " times", nil
	}

	logger.Infof("URL %s Have Not Matched Any Sensitive Words \n", urlStr)
	return "Not found", nil
}

// 获取网页正文
func getDocTexts(urlStr string) (string, error) {
	res, err := http.Get(urlStr)
	if err != nil {
		// logger.Errorf("Get URL:%s. An error occurred, couse of:%s", urlStr, err.Error())
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		// logger.Errorf("status code error: %d %s", res.StatusCode, res.Status)
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		// logger.Errorf("Query URL:%s. An error occurred, couse of:%s", urlStr, err.Error())
		return "", err
	}
	texts := strings.Trim(doc.Text(), " ")
	return texts, err
}
