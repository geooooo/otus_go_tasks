package main

import "fmt"

// Реализовать двусвязанный список на языке Go

// Item элемент списка
// Value() возвращает значение
// Next() следующий Item
// Prev() предыдущий Item

type IntItem struct {
	next  *IntItem
	prev  *IntItem
	value int
}

func (item *IntItem) Value() int {
	return item.value
}

func (item *IntItem) Next() *IntItem {
	return item.next
}

func (item *IntItem) Prev() *IntItem {
	return item.prev
}

// List тип контейнер
// Len() длинна списка
// First() первый Item
// Last() последний Item
// PushFront(i Item) добавить значение в начало
// PushBack(i Item) добавить значение в конец
// Remove(i Item) удалить элемент

type IntList struct {
	rootItem *IntItem
}

func (list *IntList) Len() int {
	if list.rootItem == nil {
		return 0
	}

	listLen := 0

	for item := list.rootItem; item.next != nil; item = item.next {
		listLen++
	}

	return listLen + 1
}

func (list *IntList) First() *IntItem {
	return list.rootItem
}

func (list *IntList) Last() *IntItem {
	if list.rootItem == nil {
		return nil
	}

	var lastItem *IntItem

	for lastItem = list.rootItem; lastItem.next != nil; lastItem = lastItem.next {
	}

	return lastItem
}

func (list *IntList) PushFront(newItem *IntItem) {
	if list.rootItem == nil {
		list.rootItem = newItem
	} else {
		newItem.next = list.rootItem
		list.rootItem.prev = newItem
		list.rootItem = newItem
	}
}

func (list *IntList) PushBack(newItem *IntItem) {
	if list.rootItem == nil {
		list.rootItem = newItem
		return
	}

	var lastItem *IntItem
	for lastItem = list.rootItem; lastItem.next != nil; lastItem = lastItem.next {
	}

	lastItem.next = newItem
	newItem.prev = lastItem
}

func (list *IntList) Remove(removedItem *IntItem) {
	isFound := false

	if removedItem == list.rootItem {
		if list.rootItem.next == nil {
			list.rootItem = nil

			isFound = true
		} else {
			nextItem := list.rootItem.next
			nextItem.prev = nil
			list.rootItem = nextItem

			isFound = true
		}
	} else {
		for item := list.rootItem; item != nil && !isFound; item = item.next {
			if item == removedItem {
				if item.next == nil {
					prevItem := item.prev
					prevItem.next = nil
				} else {
					prevItem := item.prev
					nextItem := item.next
					prevItem.next = nextItem
					nextItem.prev = prevItem
				}

				isFound = true
			}
		}
	}

	if isFound {
		removedItem.next = nil
		removedItem.prev = nil
	}
}

func main() {
	list := IntList{}

	fmt.Println(list.Len(), list.First(), list.Last())
}
