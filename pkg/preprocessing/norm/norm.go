package norm

import (
	"strings"
	"unicode"
)

func Norm(input string) string {
	var b strings.Builder
	splitted := strings.Fields(input)
	for i, v := range splitted {
		if i > 0 {
			b.WriteRune(' ')
		}
		for _, c := range v {
			if unicode.IsLetter(c) || unicode.IsNumber(c) {
				b.WriteRune(c)
			}
		}
	}
	return b.String()
}
