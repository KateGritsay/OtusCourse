//Частотный анализ
//Написать функцию, которая получает на вход текст и возвращает
//10 самых часто встречающихся слов без учета словоформ

package Exercises

import (
	"sort"
	"strings"
	"unicode"
)

type Ex4Output struct {
	Key   string
	Value int
}

func Exersice4(inputText string) []Ex4Output {
	inputText = strings.ToLower(inputText)

	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	var wordCount = make(map[string]int)
	var separatedText = strings.FieldsFunc(inputText, f)
	result := []Ex4Output{}
	
	for index, val := range separatedText {
		if val == " " {
			continue
		}

		wordCount[val] = 0
		tail := separatedText[index:]
		for indexTail, valTail := range tail {
			if valTail == " " {
				continue
			}
			if val == valTail {
				wordCount[val]++
				tail[indexTail] = " "
			}
		}
	}

	for indexWordCount, elementWordCount := range wordCount {
		result = append(result, Ex4Output{indexWordCount, elementWordCount})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Value > result[j].Value
	})
	if len(result) < 10 {
		return result
	}
	return result[0:10]
}

