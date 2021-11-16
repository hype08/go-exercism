//Package raindrops takes an integer value and "translates" it to a particular raindrop sound by determining the integer's prime factorization.
package raindrops

import "strconv"

func Convert(number int) (s string) {
	var words string
	if number%3 == 0 {
		words += "Pling"
	}

	if number%5 == 0 {
		words += "Plang"
	}

	if number%7 == 0 {
		words += "Plong"
	}

	if words != "" {
		return words
	} else {
		return strconv.Itoa(number)
	}
}
