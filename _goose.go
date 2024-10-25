package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	goose "github.com/advancedlogic/GoOse"
)

func main() {

	var texts []string
	g := goose.New()
	article, _ := g.ExtractFromURL("https://www.baidu.com")
	println("title", article.Title)
	println("description", article.MetaDescription)
	println("keywords", article.MetaKeywords)
	println("content", article.CleanedText)
	println("url", article.FinalURL)
	println("top image", article.TopImage)
	article.Doc.Find("body").Each(func(i int, s *goquery.Selection) {
		texts = append(texts, s.Text())
	})
	for _, text := range texts {
		fmt.Println(text)
	}
}
