package kata

import (
	"strings"
)

var encoder string
var isDuble bool = false
var isFirst bool = false

func DuplicateEncode(word string) string {
	encoder = ""
	word = strings.ToLower(word)
	for _, v := range word {
		for _, v2 := range word {
			if v == v2 {
				if isFirst {
					isDuble = true
				}
				isFirst = true
			}
		}
		if isDuble {
			encoder += ")"
		} else {
			encoder += "("
		}
		isFirst = false
		isDuble = false
	}

	return encoder
}
