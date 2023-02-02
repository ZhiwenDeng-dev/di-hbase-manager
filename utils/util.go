package utils

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func EscapeHbaseCommandStr(command string) string {
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

//通过管道同步获取日志的函数
func SyncLog(reader io.ReadCloser, f *os.File, output *string) {

	buf := make([]byte, 1024, 1024)
	for {
		strNum, err := reader.Read(buf)
		if strNum > 0 {
			outputByte := buf[:strNum]
			*output = *output + string(outputByte)
			f.WriteString(string(outputByte))
		}
		if err != nil {
			//读到结尾
			if err == io.EOF || strings.Contains(err.Error(), "file already closed") {
				err = nil
			}
		}
	}
}
