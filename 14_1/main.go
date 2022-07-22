package main

import (
	"fmt"
)

// Цель: В результате выполнения ДЗ должен получиться базовый скелет микросервиса,
// который будет развиваться в дальнейших ДЗ.
//
// Cоздать методы бизнес логики (методы у структур) для работы с этими структурами:
// - добавление событий в хранилище
// - удаление событий из хранилища
// - изменение событий в хранилище
// - листинг событий
// - пр. на усмотрение студента
//
// Создать объекты ошибок (error sentinels) соответсвующие бизнес ошибкам,
// например ErrDateBusy - данное время уже занято другим событием
//
// Реализовать хранение событий в памяти (т.е. просто складывать объекты в слайсы)

func main() {
	handler := EventHandler(func(event *Event) {
		fmt.Println(event)
	})

	event1 := Event{}
	event2 := Event{}
	event3 := Event{}
	event4 := Event{}

	store := EventStore{}

	if error := store.AddHandler(&handler); error != nil {
		fmt.Println(error)
	}
	if error := store.AddHandler(&handler); error != nil {
		fmt.Println(error)
	}

	if error := store.AddEvent(&event1); error != nil {
		fmt.Println(error)
	}
	if error := store.AddEvent(&event1); error != nil {
		fmt.Println(error)
	}
	if error := store.AddEvent(&event2); error != nil {
		fmt.Println(error)
	}
	if error := store.AddEvent(&event3); error != nil {
		fmt.Println(error)
	}

	if error := store.UpdateEvent(&event4, &event1); error != nil {
		fmt.Println(error)
	}
	if error := store.UpdateEvent(&event1, &event4); error != nil {
		fmt.Println(error)
	}

	if error := store.RemoveEvent(&event1); error != nil {
		fmt.Println(error)
	}
	if error := store.RemoveEvent(&event1); error != nil {
		fmt.Println(error)
	}
	if error := store.RemoveEvent(&event2); error != nil {
		fmt.Println(error)
	}
	if error := store.RemoveEvent(&event3); error != nil {
		fmt.Println(error)
	}
	if error := store.RemoveEvent(&event4); error != nil {
		fmt.Println(error)
	}

	if error := store.RemoveHandler(&handler); error != nil {
		fmt.Println(error)
	}
	if error := store.RemoveHandler(&handler); error != nil {
		fmt.Println(error)
	}

	fmt.Println(store.events, store.handlers)
}
