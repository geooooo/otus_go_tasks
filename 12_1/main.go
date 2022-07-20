package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// Цель: Реализовать утилиту envdir на Go.
// Эта утилита позволяет запускать программы получая переменные окружения
// из определенной директории. См man envdir
// Пример go-envdir /path/to/evndir command arg1 arg2
//
// Реализовать функцию вида ReadDir(dir string) (map[string]string, error),
// которая сканирует указанный каталог и возвращает все переменные окружения,
// определенные в нем.
//
// Реализовать функцию вида RunCmd(cmd []string, env map[string]string) int,
// которая запускает программу с аргументами (cmd) c переопределнным окружением.
//
// Реализовать функцию main, анализирующую аргументы командной строки и
// вызывающую ReadDir и RunCmd

func main() {
	directoryName, commandArguments, error := initArgs()
	if error != nil {
		fmt.Println(error)
		os.Exit(0)
	}

	environmentVars, error := ReadDir(directoryName)
	if error != nil {
		fmt.Println(error)
		os.Exit(0)
	}

	returnCode := RunCmd(commandArguments, environmentVars)

	os.Exit(returnCode)
}

func initArgs() (string, []string, error) {
	flag.Parse()

	directoryName := flag.Arg(0)
	if directoryName == "" {
		return "", nil, errors.New("missed directory error")
	}

	commandArguments := flag.Args()[1:]
	if len(commandArguments) == 0 {
		return "", nil, errors.New("missed command error")
	}

	return directoryName, commandArguments, nil
}

func ReadDir(directoryName string) (map[string]string, error) {
	fileInfos, error := ioutil.ReadDir(directoryName)
	if error != nil {
		return nil, errors.New("reading directory error")
	}

	environmentVars := map[string]string{}
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}

		fileName := fileInfo.Name()
		filePath := directoryName + string(os.PathSeparator) + fileName
		data, error := os.ReadFile(filePath)
		if error != nil {
			return nil, errors.New("reading file error")
		}

		environmentVars[fileName] = strings.Trim(string(data), " \n\t")
	}

	return environmentVars, nil
}

func RunCmd(commandArguments []string, environmentVars map[string]string) int {
	cmd := exec.Command(commandArguments[0], commandArguments[1:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()

	for name, value := range environmentVars {
		environmentVar := name + "=" + value
		cmd.Env = append(cmd.Env, environmentVar)
	}

	error := cmd.Run()
	if error != nil {
		exitError := error.(*exec.ExitError)
		return exitError.ExitCode()
	}

	return 0
}
