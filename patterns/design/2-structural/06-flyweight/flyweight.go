package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Text Formatting: FormattedText with capitalize []bool to indicate if they are capitalized; Introduce TextRange type with Start, End & bools; And BetterFormattedText with []*TextRange
// User Names: User with FullName; Introduce allNames slice

type FormattedText struct {
	Text       string
	Capitalize []bool
}

func (ft *FormattedText) String() string {
	b := strings.Builder{}

	for idx, char := range ft.Text {
		if ft.Capitalize[idx] {
			b.WriteRune(unicode.ToUpper(char))
		} else {
			b.WriteRune(char)
		}
	}

	return b.String()
}

func main() {
	ft := &FormattedText{"fortuna aduvait juvat", []bool{false, true, true, true, true, ...}}

	fmt.Println(ft)
}
