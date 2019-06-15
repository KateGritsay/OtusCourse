package main

import (
	"strconv"
	"unicode"
	"fmt"
)

func Test() {
	fmt.Println("sadfasfd")
}

var startString string = "a4bc5de"

var runeString = []rune (startString)

func Exercise2(){

	for  i := 0; i < len(runeString); i++ {
		if unicode.IsLetter(runeString[i]) {
			fmt.Printf("%v", string(runeString[i]))

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
				fmt.Printf("Error parse string to int")
			}
			for x := 0; x < int(h)-1; x++ {
				fmt.Printf("%v", string(runeString[i]))
			}

		} else {
			continue
		}

	}
}