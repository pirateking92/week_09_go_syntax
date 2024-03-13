package reverse

import (
	"strings"
)

func ReverseName(name string) string {
	const space = " "
	words := strings.Split(name, space)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, space)
}
