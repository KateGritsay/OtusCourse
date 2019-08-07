//Частотный анализ
//Написать функцию, которая получает на вход текст и возвращает
//10 самых часто встречающихся слов без учета словоформ

package exercise4

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func Request()  {
	fmt.Println(Exersice4("Я за пришла за яйцами, я в я магазин за в яйцами за"))
}

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

	for _, word := range separatedText{
		wordCount[word]++
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

