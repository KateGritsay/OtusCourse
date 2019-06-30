//Реализовать двусвязанный список.
//Что такое двусвязный список: https://en.wikipedia.org/wiki/Doubly_linked_list
//
//Ожидаемые типы (псевдокод):
//
//```
//List      // тип контейнер
//  Len()   // длинна списка
//  First() // первый Item
//  Last()  // последний Item
//  PushFront(v interface{}) // добавить значение в начало
//  PushBack(v interface{})  // добавить значение в конец
//
//Item   // элемент списка
//  Value() interface{}  // возвращает значение
//  Nex() *Item          // следующий Item
//  Prev() *Item         // предыдущий
//  Remove()             // удалить Item из списка
//```

package Ex6

import "fmt"

func Request() {
	list := List {}
	list.PushFront(10)
	list.PushFront(20)
	list.PushFront(30)
	list.PushFront(40)

	list.Remove(list.lastItem)
	list.Remove(list.firstItem)
	list.Remove(list.lastItem)
	list.Remove(list.firstItem)
	list.Remove(list.lastItem)

	fmt.Println(list)
}

type List struct {
	firstItem *Item
	lastItem *Item
	length int
}
type Item struct {
	prev *Item
	next *Item
	value interface{}
}

func (list *List) Len() int {
	return list.length
}

func (list *List) First() *Item {
	return list.firstItem
}

func (list *List) Last() *Item {
	return list.lastItem
}

func (list *List) PushFront(value interface{}) {

	item := Item{}
	item.value = value
	item.next = list.firstItem
	item.prev = nil

	if list.firstItem != nil{
		list.firstItem = &item
		item.next.prev = &item
	} else {
		list.firstItem = &item
		list.lastItem = &item
	}

	list.length++
}

func (list *List) PushBack(value interface{}) {
	item := Item{}
	item.value = value
	item.next = nil
	item.prev = list.lastItem

	if list.lastItem != nil{
		list.lastItem = &item
		item.prev.next = &item
	} else {
		list.lastItem = &item
	}

	list.length++
}

func (item *Item) Prev() *Item {
	return item.prev
}

func (item *Item) Next() *Item {
	return item.next
}

func (list *List) Remove(item *Item) {
	if list.length == 0 {
		return
	}
	if item.prev != nil {
		elementNext := item.prev
		elementNext.next = item.next
	} else {
		list.firstItem = item.next
	}
	if item.next != nil {
		elementPrev := item.next
		elementPrev.prev = item.prev
	} else {
		list.lastItem = item.prev
	}
	list.length--
}