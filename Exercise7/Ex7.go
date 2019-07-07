//Функция логирования Otus
//Задача: написать функцию логирования LogOtusEvent, на вход которой приходят события типа HwAccepted (домашняя работа принята) и HwSubmitted (студент отправил дз)
// для этого - спроектировать и реализовать интерфейс OtusEvent.
// Для события HwAccepted мы хотим логирровать дату, айди и грейд, для HwSubmitter - дату, id и комментарий, например:
//
//2019-01-01 submitted 3456 "please take a look at my homework"
//2019-01-01 accepted 3456 4
//
//
//type HwAccepted struct {
//Id int
//Grade int
//}
//
//type HwSubmitted struct {
//Id int
//Code string
//Comment string
//}
//
//interface OtusEvent {
//}
//
//function LogOtusEvent(e OtusEvent, w io.Writer) {
//
//}

package Ex7

import (
	"fmt"
	"io"
	"os"
	"time"
)

type HwAccepted struct {
	Id int
	Grade int
}

type HwSubmitted struct {
	Id int
	Code string
	Comment string
}

type OtusEvent interface {
	Happening() string
}

var dt = time.Now().Format("2006-01-01")

func (ha HwAccepted) Happening() string {
	return fmt.Sprint(dt, " accepted ", ha.Id, ha.Grade, " \r" )
}

func (hs HwSubmitted) Happening() string {
	return fmt.Sprint(dt, " submitted ", hs.Id, hs.Comment )
}

func LogOtusEvent (h OtusEvent, w io.Writer) {
	fmt.Fprintf(w,"%s", h.Happening())
}

func Request()  {
	accepted := HwAccepted{123232, 4}
	file, _ := os.Create("/Users/Kate/Documents/GitHub/OtusCourse/dat.txt")
	LogOtusEvent(accepted,file)

	submitted := HwSubmitted{3453, ""," please take a look at my homework"}
	LogOtusEvent(submitted,file)
}