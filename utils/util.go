package utils

import (
	"fmt"
	"unicode"
)

func ConcatHbaseCommandStr(command string) string {
	var chars []rune
	for _, letter := range command {
		ok, letters := SpecialLetters(letter)
		if ok {
			chars = append(chars, letters...)
		} else {
			chars = append(chars, letter)
		}
	}
	return fmt.Sprintf(`echo %s | hbase shell -n`, string(chars))
}

func SpecialLetters(letter rune) (bool, []rune) {
	if unicode.Is(unicode.Han, letter) {
		panic("Unsupported Chinese characters")
	}
	if unicode.IsPunct(letter) || unicode.IsSymbol(letter) {
		var chars []rune
		chars = append(chars, '\\', letter)
		return true, chars
	}
	return false, nil
}
