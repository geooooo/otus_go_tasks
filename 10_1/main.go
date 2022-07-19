package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

// Цель: Реализовать утилиту копирования файлов
// gocopy -­from /path/to/source ­-to /path/to/dest -­offset 1024 -­limit 2048
//
// Утилита должна принимать следующие аргументы:
// * файл источник (From)
// * файл копия (To)
// * Отступ в источнике (Offset), по умолчанию - 0
// * Количество копируемых байт (Limit), по умолчанию - весь файл из From
//
// Реализовать функцию вида Copy(from string, to string, limit int, offset int) error

func main() {
	sourceFileName, destinationFileName, countOfCopyBytes, sourceFileOffset := initArgs()

	error := Copy(sourceFileName, destinationFileName, countOfCopyBytes, sourceFileOffset)
	if error == nil {
		fmt.Println("Success")
	} else {
		fmt.Println(error)
	}
}

func Copy(
	sourceFileName string,
	destinationFileName string,
	countOfCopyBytes uint,
	sourceFileOffset uint,
) error {
	sourceFile, destinationFile, error := openFiles(sourceFileName, destinationFileName)
	if error != nil {
		return error
	}

	defer sourceFile.Close()
	defer destinationFile.Close()

	if countOfCopyBytes == 0 {
		sourceFileSize, error := getFileSize(sourceFile)
		if error != nil {
			return error
		}

		countOfCopyBytes = sourceFileSize - sourceFileOffset
	}

	error = setFileOffset(sourceFile, sourceFileOffset)
	if error != nil {
		return error
	}

	_, error = io.CopyN(destinationFile, sourceFile, int64(countOfCopyBytes))
	if error != nil {
		return errors.New("copy files error")
	}

	return nil
}

func getFileSize(file *os.File) (uint, error) {
	fileInfo, error := file.Stat()
	if error != nil {
		return 0, errors.New("get source file size error")
	}

	return uint(fileInfo.Size()), nil
}

func openFiles(sourceFileName string, destinationFileName string) (*os.File, *os.File, error) {
	sourceFile, error := os.Open(sourceFileName)
	if error != nil {
		return nil, nil, errors.New("open source file error")
	}

	destinationFile, error := os.Create(destinationFileName)
	if error != nil {
		return nil, nil, errors.New("open destination file error")
	}

	return sourceFile, destinationFile, nil
}

func setFileOffset(file *os.File, offset uint) error {
	_, error := file.Seek(int64(offset), io.SeekStart)
	if error != nil {
		return errors.New("set source file offset error")
	}

	return nil
}

func initArgs() (string, string, uint, uint) {
	sourceFileName := flag.String("from", "", "source file for a reading")
	destinationFileName := flag.String("to", "", "destination file for a writing")
	countOfCopyBytes := flag.Uint("limit", 0, "limit of bytes count for a reading")
	sourceFileOffset := flag.Uint("offset", 0, "offset in the source file")

	flag.Parse()

	return *sourceFileName, *destinationFileName, *countOfCopyBytes, *sourceFileOffset
}
