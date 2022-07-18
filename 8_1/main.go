package main

import (
	"errors"
	"fmt"
)

// Параллельное исполнение
// Написать функцию для параллельного выполнения N заданий
// (т.е. в N параллельных горутинах).
// Функция принимает на вход:
// - слайс с заданиями `[]func() error`;
// - число заданий которые можно выполнять параллельно (`N`);
// - максимальное число ошибок после которого нужно приостановить обработку.

type Job func() error

type JobCompleteSignal error

func main() {
	jobs := []Job{}

	for i := 0; i < 10000; i++ {
		jobs = append(jobs, func() error {
			for i := 0; i < 1000000; i++ {
			}
			return nil
		})
	}

	error := runTasks(jobs, 10, 1)

	if error != nil {
		fmt.Printf("%v\n", error)
	} else {
		fmt.Printf("Success\n")
	}
}

func runTasks(jobs []Job, parallelCount uint, maxErrorCount uint) error {
	if parallelCount == 0 || len(jobs) == 0 {
		return errors.New("invalid arguments")
	}

	jobCompleteChannel := make(chan JobCompleteSignal, parallelCount)

	defer func() {
		close(jobCompleteChannel)
	}()

	countOfRunnedJobs := uint(0)
	errorCount := uint(0)

	for _, job := range jobs {
		go runTask(jobCompleteChannel, job)

		countOfRunnedJobs++

		if countOfRunnedJobs == parallelCount {
			error := <-jobCompleteChannel

			if error != nil {
				errorCount++
			}

			if errorCount >= maxErrorCount {
				return errors.New("max error count")
			}

			countOfRunnedJobs--
		}
	}

	if countOfRunnedJobs == 0 {
		return nil
	}

	for error := range jobCompleteChannel {
		if error != nil {
			errorCount++
		}

		if errorCount == maxErrorCount {
			return errors.New("max error count")
		}

		countOfRunnedJobs--

		if countOfRunnedJobs == 0 {
			break
		}
	}

	return nil
}

func runTask(jobCompleteChannel chan<- JobCompleteSignal, job Job) {
	error := job()
	jobCompleteChannel <- error
}
