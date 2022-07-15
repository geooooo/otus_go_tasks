package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Распаковка строки
// Создать Go функцию, осуществляющую примитивную распаковку строки,
// содержащую повторяющиеся символы / руны, например:

// * "a4bc2d5e" => "aaaabccddddde"
// * "abcd" => "abcd"
// * "45" => "" (некорректная строка)

// Дополнительное задание: поддержка escape - последовательности
// * `qwe\4\5` => `qwe45` (*)
// * `qwe\45` => `qwe44444` (*)
// * `qwe\\5` => `qwe\\\\\` (*)

func main() {
	fmt.Println(UnpackString(`qwe\4\5`))
}

func UnpackString(packedString string) string {
	result := ""

	letters := strings.Split(packedString, "")
	lettersCount := len(letters)

	for i := 0; i < lettersCount; i++ {
		letterAsString := letters[i]
		letterAsRune, _ := utf8.DecodeRuneInString(letterAsString)

		hasFirstEscape := letterAsString == `\`
		hasEscapedLetter := false

		if hasFirstEscape && i < lettersCount-1 {
			nextLetterAsString := letters[i+1]
			nextLetterAsRune, _ := utf8.DecodeRuneInString(nextLetterAsString)
			hasSecondEscape := nextLetterAsString == `\`

			if hasSecondEscape || unicode.IsNumber(nextLetterAsRune) {
				letterAsString = nextLetterAsString
				hasEscapedLetter = true
				i++
			} else {
				return ""
			}
		}

		if unicode.IsLetter(letterAsRune) || hasEscapedLetter {
			additionalRepeatCount := 0

			for ; i < lettersCount-1; i++ {
				nextLetterAsString := letters[i+1]
				nextLetterAsRune, _ := utf8.DecodeRuneInString(nextLetterAsString)

				if !unicode.IsNumber(nextLetterAsRune) {
					break
				}

				if nextLetterAsString == "0" {
					return ""
				}

				digit, _ := strconv.Atoi(nextLetterAsString)

				if additionalRepeatCount == 0 {
					additionalRepeatCount = digit
				} else {
					additionalRepeatCount = additionalRepeatCount*10 + digit
				}
			}

			repeatCount := 1
			if additionalRepeatCount > 0 {
				repeatCount = additionalRepeatCount
			}

			result = result + strings.Repeat(letterAsString, repeatCount)
		} else {
			return ""
		}
	}

	return result
}
