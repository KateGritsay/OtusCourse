//Распаковка строки
//Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
//
//* "a4bc2d5e" => "aaaabccddddde"
//* "abcd" => "abcd"
//* "45" => "" (некорректная строка)
//* `qwe\4\5` => `qwe45` (*)
//* `qwe\45` => `qwe44444` (*)
//* `qwe\\5` => `qwe\\\\\` (*)

package Exercise3

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func Request() {
	requestString := "a4bc2d5e"
	resultString, err := Exercise3(requestString)

	if err != nil {
		log.Printf(err.Error())
	}

	fmt.Println(resultString)
}

func Exercise3(startString string) (string, error) {

	var runeString = []rune(startString)
	var result string

	for i := 0; i < len(runeString); i++ {
		if !unicode.IsLetter(runeString[i]) {
			continue
		}
		result += string(runeString[i])

		if i == len(runeString)-1 {
			break
		}
		if unicode.IsLetter(runeString[i+1]) {
			continue
		}

		num := string(runeString[i+1])
		for j := i + 2; j < len(runeString); j++ {
			if unicode.IsLetter(runeString[j]) {
				break
			} else {
				num += string(runeString[j])
			}
		}
		h, err := strconv.ParseInt(num, 0, 16)

		if err != nil {
			return "", err
		}
		result += strings.Repeat(string(runeString[i]), int(h)-1)
	}
	return result, nil
}
