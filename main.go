package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	//fmt.Println(Exercise2("a3b3c3"))
	fmt.Println(Exercise3("Об этом Sohu сообщает Sohu китайское издание Sohu ем самым, пишут авторы материала, Россия вынудила буквально весь мир сомневаться в военном превосходстве США. Соединенные Штаты Америки, признает Sohu, начали разрабатывать гиперзвуковое оружие задолго до Российской Федерации — еще в период холодной войны. 		Сегодня же, сообщает Sohu, Россия демонстрируется свои гиперзвуковые ракеты"))
}


func Exercise2(startString string) string{

	var runeString = []rune (startString)
	var result string

	for  i := 0; i < len(runeString); i++ {
		if !unicode.IsLetter(runeString[i]) {
			continue
		}
		result += string(runeString[i])

		if i == len(runeString)-1 {
			break
		}
		if unicode.IsLetter(runeString[i+1]){
			continue
		}

		num:= string(runeString[i+1])
		for j := i + 2; j < len(runeString); j++ {
			if unicode.IsLetter(runeString[j]) {
				break
			} else {
				num += string(runeString[j])
			}
		}
		h, err := strconv.ParseInt(num, 0, 16)

		if err != nil {
			log.Printf("Error parse string to int")
		}
		result += strings.Repeat(string(runeString[i]), int(h)-1)
	}
	return result;
}

type Ex3Output struct {
	Key string
	Value int
}
func Exercise3(inputText string) []Ex3Output {
	inputText = strings.ReplaceAll(inputText, ",", "")
	inputText = strings.ReplaceAll(inputText, ".", "")
	inputText = strings.ReplaceAll(inputText, ":", "")
	inputText = strings.ReplaceAll(inputText, "!", "")
	inputText = strings.ReplaceAll(inputText, "?", "")
	inputText = strings.ReplaceAll(inputText, ";", "")
	inputText = strings.ToLower(inputText)
	
	var separatedText  = strings.Split(inputText," ")
	result := []Ex3Output{}

	for index, val := range separatedText {
		if val == " " {
			continue
		}
		var counter int = 0
		tail := separatedText [index:]
		for indexTail, valTail := range tail {
			if valTail == " "{
				continue
			}
			if val == valTail {
				counter++
				tail[indexTail] = " "
			}
		}
		result = append(result, Ex3Output{val, counter})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Value > result[j].Value
	})
	return result[0:10]
}
