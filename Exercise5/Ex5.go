//Поиск максимума
//Написать функцию находящую максимальный элемент в слайсе
//с произвольными элементами ([]interface{}) с использованием
//пользовательской функции-компаратора.


package Ex5

import (
	"fmt"
	"sort"
)

func FindMax(someSlice []interface{}, s func(i, j int) bool) (result interface{}) {

	if len(someSlice) == 0 {
		return nil
	}
	sort.Slice(someSlice, s)
	return someSlice[0]
}


func Request () {
	requestSlice := []interface{}{}
	myMaxString := FindMax(requestSlice, func(i, j int) bool {
		return requestSlice[i].(string) < requestSlice[j].(string)
	})

	if myMaxString == nil {
		fmt.Print("Slice is empty\n")
	} else {
		fmt.Println("String maximum is", myMaxString.(string))
	}


	requestSlice = []interface{}{1,2,3}
	myMaxInt := FindMax(requestSlice, func(i, j int) bool {
		return requestSlice[i].(int) > requestSlice[j].(int)
	})

	if myMaxInt == nil {
		fmt.Print("Slice is empty\n")
	} else {
		fmt.Println("Int maximum is",myMaxInt.(int) )
	}
}

