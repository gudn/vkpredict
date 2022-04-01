package norm

import (
	"strings"
	"unicode"
)

var replaceBySpaces = "-/()*@;:"

func Norm(input string) string {
	var b strings.Builder
	splitted := strings.Fields(input)
	wasSpace := false
	for i, v := range splitted {
		if i > 0 && !wasSpace {
			b.WriteRune(' ')
			wasSpace = true
		}
		for _, c := range v {
			if unicode.IsLetter(c) || unicode.IsNumber(c) {
				b.WriteRune(c)
				wasSpace = false
			}
			if strings.ContainsRune(replaceBySpaces, c) && !wasSpace {
				b.WriteRune(' ')
				wasSpace = true
			}
		}
	}
	return b.String()
}
