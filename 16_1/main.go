package main

import (
	"fmt"
	"os"
)

// Цель: Реализовать "каркас" микросервиса,
// считывающий конфиг из файла, создающий логгер/логгеры
// с указанными уровнями детализации.
//
// Необходимо доработать код сервиса "Календарь"
// из предыдущего задания, добавив в него:
// * Обработку аргументов командной строки
// * Чтение файла конфигурации (параметр --config в командной строке)
// * Создание логгеров и настройка уровня логирования
// * Создание и запуск hello-world web-сервера
//
// Параметры, передаваемые через аргументы командной строки:
// * --config - путь к конфигу
//
// Параметры, которые должны быть в конфиге:
// * http_listen - ip и port на котором должен слушать web-сервер
// * log_file - путь к файлу логов
// * log_level - уровень логирования (error / warn / info / debug)

func main() {
	pathToConfig := InitArgs()

	config, error := InitConfig(pathToConfig)
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}

	fmt.Printf("%v\n", config)

	InitLogger(config.PathToLogFile, config.LogLevel)

	Log("test")

	CloseLogger()
}
