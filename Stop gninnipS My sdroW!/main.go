package kata

import (
  "strings"
)

func SpinWords(str string) string {
  aTokens := strings.Split(str, " ")
  for i, word := range aTokens {
    if len([]rune(word)) > 4 {
      aTokens[i] = reverse(word)
    }
  }
  return strings.Join(aTokens, " ")
}// SpinWords

func reverse(s string) string {
  runes := []rune(s)
  for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
    runes[i], runes[j] = runes[j], runes[i]
  }
  return string(runes)
}
