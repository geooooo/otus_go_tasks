package main

import (
	"fmt"

	"github.com/lazybark/go-helpers/npt"
)

// Hello now()
// Завести Go репозиторий на GitHub,
// написать программу печатающую текущее время / точное время
// с использованием библиотеки NTP.

func main() {
	hour, minute, second := now()
	fmt.Printf("%d:%d:%d\n", hour, minute, second)
}

func now() (int, int, int) {
	npt := npt.Now()
	time := npt.Time()

	hour := time.Hour()
	minute := time.Minute()
	second := time.Second()

	return hour, minute, second
}
