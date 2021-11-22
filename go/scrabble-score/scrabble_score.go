// Package scrabble helps calculate Scrabble score
package scrabble

import "strings"

func Score(word string) int {
	values := map[int]string{
		1: "AEIOULNRST",
		2: "DG",
		3: "BCMP",
		4: "FHVWY",
		5: "K",
		8: "JX",
		10: "QZ",
	}
  word = strings.ToUpper(word)

  var score int
  for _, c := range word {
    for k, v := range values {
      if strings.Contains(v, string(c)) {
        score += k
      }
    }
  }

  return score
}
