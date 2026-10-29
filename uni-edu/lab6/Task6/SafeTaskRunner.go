package main

import (
	"errors"
	"fmt"
)

type Task struct {
	Name string
	Run  func() error
}

func SafeRun(tasks []Task) error {
	var errs []error

	for _, task := range tasks {
		// Обертка для перехвата паники внутри задачи
		err := func() (errRet error) {
			defer func() {
				if r := recover(); r != nil {
					errRet = fmt.Errorf("panic: %v", r)
				}
			}()
			return task.Run()
		}()

		if err != nil {
			wrapped := fmt.Errorf("task %s: %w", task.Name, err)
			errs = append(errs, wrapped)
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
