package bootstrap

import (
	"blackwebfilter/internal/logger"
	"fmt"

	ahocorasick "github.com/BobuSumisu/aho-corasick"
	"go.uber.org/zap"
)

func NewTrieBuilder(env *Env) (*ahocorasick.Trie, error) {
	builder := ahocorasick.NewTrieBuilder()
	err := builder.LoadStrings(env.SensitiveWordsPath)
	if err != nil {
		logger.Error("Environment of SENSITIVE_WORDS_PATH can't be loaded: ", zap.Error(err))
		return nil, err
	}
	trie := builder.Build()
	return trie, err
}

func TrieTest(word string) int {
	app := App()
	env := app.Env
	trie, err := NewTrieBuilder(env)
	if err != nil {
		panic(err)
	}
	matches := trie.MatchString(word)
	for _, match := range matches {
		fmt.Printf("Matched pattern %v %q at position %d.\n", string(match.Match()),
			match.Pattern(), match.Pos())
	}
	return len(matches)
}
