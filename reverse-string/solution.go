package reversor

import "strings"

// ReverseRunes returns its argument string reversed rune-wise left to right.
// Stolen from: https://go.dev/doc/code#ImportingLocal
func ReverseRunes(s string) string {
	r := []rune(s)
	for lower, upper := 0, len(r)-1; lower < len(r)/2; lower, upper = lower+1, upper-1 {
		r[lower], r[upper] = r[upper], r[lower]
	}
	return string(r)
}

// Home made solution
func Reverse(text string) string {
	var reversed string = ""
	for _, char := range text {
		reversed = string(char) + reversed
	}
	return reversed
}

// taken from Sabrina at => https://github.com/Sabrina-Incinga/labora-go/blob/c01ed40afbe3bc6bc8dbe907e91d0afbc0d00c5d/seccion-2/desafio_invertir_palabras/invertirPalabras.go#L13
func InvertWords(word string) string {
	pushLettersToSlice := func(s string, wordSlice *[]string) {
		for _, v := range s {
			b := func(x rune) {
				*wordSlice = append(*wordSlice, string(x))
			}
			defer b(v)
		}
	}

	var slice []string
	pushLettersToSlice(word, &slice)

	return strings.Join(slice, "")

}
