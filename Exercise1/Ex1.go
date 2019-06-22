//Написать программу печатающую текущее время / точное время с использованием библиотеки NTP


package Ex1

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
)

func Request() {
	currentTime, err := ntp.Time("time.apple.com")
	if err != nil {
		log.Print(err)
	} else {
		fmt.Println(currentTime)
	}
}