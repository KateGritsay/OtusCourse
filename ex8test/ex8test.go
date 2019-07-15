package ex8test

import "fmt"
import "errors"

func Parallel (tasks []func()error, tasksMaxNumber int, errorsMaxNumber int) string {

tasksLimitCh := make (chan struct{}, tasksMaxNumber) //limit parallel functions
tasksResultCh := make (chan interface{}, len(tasks)) //all results
tasksErrorsCh := make (chan error, errorsMaxNumber ) //only result errors

for _, i := range tasks {
	go func(task func()error) {
		tasksLimitCh <- struct {}{}
		switch {
			case len(tasksErrorsCh) < errorsMaxNumber:
			doingTasks(task, tasksResultCh, tasksErrorsCh)
			<-tasksLimitCh

			case len(tasksLimitCh) >= tasksMaxNumber:
				break
		}
	}(i)
}

for range tasksResultCh{
	count := 0
	count++
	if count >= tasksMaxNumber || len(tasksErrorsCh) >= errorsMaxNumber {
		break
	}
}
	return "Done"
}

func doingTasks (task func()error, tasksResultCh chan<- interface{},tasksErrorsCh chan<- error){
	result := task()
	tasksResultCh <- result
	if result != nil{
		tasksErrorsCh <- result
	}
}

func Request () {
	tasksMaxNumber := 3
	errorsMaxNumber := 3

	tasks := []func() error{}
	for i := 0; i < 10; i++ {
		tasks = append(tasks, func() error {
			fmt.Println("Testing string")
			return nil
		})
	}

	for i := 0; i < 10; i++ {
		tasks = append(tasks, func() error {
			tmp := 0
			for i := 0; i < 100000; i++ {
				tmp += i
			}
			fmt.Println(tmp)

			return nil
		})
	}

	for i := 0; i < 10; i++ {
		tasks = append(tasks, func() error {
			fmt.Println("Fake error")
			return errors.New("Fake error")
		})

	}

	fmt.Println(Parallel(tasks, tasksMaxNumber, errorsMaxNumber))
}