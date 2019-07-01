package main

import (
	"fmt"
	"log"
	"sync"
)


func worker(functions []func()error, maxTaskQty int, maxErrQty int) {

	var wg sync.WaitGroup

	for index, function := range functions {
		if index == maxTaskQty {
			break
		}

		wg.Add(1)
		go func(f func()error) {
			defer wg.Done()

			err := f()
			if err != nil {
				log.Println("Error:", err.Error())
			}
		}(function)
	}

	wg.Wait()
}

func main() {
	functions := []func()error {
		func() error {
			log.Println("1st call")
			return nil
		},
		func() error {
			log.Println("2nd call")
			return fmt.Errorf("error 2")
		},
		func() error {
			log.Println("3d call")
			return fmt.Errorf("error 3")
		},
	}

	worker(functions, 3, 1)
}
