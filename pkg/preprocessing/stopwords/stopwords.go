// Удаление стоп-слов
//
// Стоп-слова храняется в файле рядом (необязательно на разных строчках)
package stopwords

import (
	_ "embed"
	"strings"
)

//go:embed stopwords.txt
var contents string

var stopWords map[string]struct{}

func init() {
	words := strings.Fields(contents)
	stopWords = make(map[string]struct{}, len(words))
	for _, w := range words {
		stopWords[w] = struct{}{}
	}
}

func Stopwords(input string) string {
	var b strings.Builder
	words := strings.Fields(input)
	for _, w := range words {
		if _, ok := stopWords[w]; !ok {
			b.WriteString(w)
			b.WriteRune(' ')
		}
	}
	result := b.String()
	if result == "" {
		return input
	}
	return result
}
