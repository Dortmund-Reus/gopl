// Charcount computes counts of Unicode characters.
package charchount

import (
	"bufio"
	"io"
	"log"
	"unicode"
)

func charCount(r io.Reader) (
	runes map[rune]int, props map[string]int, sizes map[int]int, invalid int) {

	runes = make(map[rune]int)   // char frequency
	props = make(map[string]int) // unicode.Properties frequency
	sizes = make(map[int]int)    // rune length frequency
	invalid = 0                  // invalid char frequency

	in := bufio.NewReader(r)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("charCount: %s", err)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		for name, rangeTable := range unicode.Properties {
			if unicode.In(r, rangeTable) {
				props[name]++
			}
		}
		runes[r]++
		sizes[n]++
	}
}
