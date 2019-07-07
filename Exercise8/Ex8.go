//Параллельное исполнение
//Сделать функцию для параллельного выполнения N заданий.
//Принимает на вход слайс с заданиями `[]func()error`, число заданий которые можно выполнять параллельно `N`
// и максимальное число ошибок после которого нужно приостановить обработку.
// Учесть что задания могу выполняться разное время.


package Ex8

import (
	"errors"
	"fmt"
)

func Parallel(tasks []func(in chan int)error, tasksCount int, errorCount int) string {

	totalErrors := 0
	if tasksCount > len(tasks){
		tasksCount = len(tasks)
	}
	for batch:=0; batch < len(tasks); batch += tasksCount {

		var chanCountErrors= make(chan int, tasksCount)

		for i := batch; i < tasksCount+batch; i++ {
			go tasks[i](chanCountErrors)
		}


		//for {
		//	i, ok := <-chanCountErrors
		//
		//	if !ok {
		//		break
		//	}
		//	totalErrors += i
		//}

		tmp := 0
		for j := range chanCountErrors {
			totalErrors += j
			tmp++
			if tmp == tasksCount {
				break
			}
		}

		close(chanCountErrors)

		if totalErrors >= errorCount {
			break
		}
	}
	return "Done"
}

func Request(){
	tasksCount := 1
	errorCount := 2

	tasks := []func(in chan int)error{}
	for i:=0; i<10; i++ {
		tasks = append(tasks, func(in chan int) error {
			fmt.Println("Some text")

			in <- 0
			return nil
		})
	}

	for i:=0; i<10; i++ {
		tasks = append(tasks, func(in chan int) error {
			tmp :=0
			for i :=0; i<10000000; i++ {
				tmp += i
			}
			fmt.Println(tmp)

			in <- 0
			return nil
		})
	}

	for i:=0; i<10; i++ {
		tasks = append(tasks, func(in chan int) error {
			in <- 1
			fmt.Println("Fake error")
			return errors.New("Fake error")
		})

	}

	fmt.Println(Parallel (tasks, tasksCount, errorCount))
}