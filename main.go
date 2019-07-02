package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)


func worker(functions []func()error, maxTaskQty int, maxErrQty int) {

	var wg sync.WaitGroup

	quit := make(chan bool, maxErrQty * 2)

	for index, function := range functions {
		if index == maxTaskQty {
			break
		}

		// just for demonstration
		time.Sleep(1 * time.Second)

		wg.Add(1)
		go func(f func()error) {
			defer wg.Done()

			if len(quit) >= maxErrQty {
				return
			}

			err := f()
			if err != nil {
				quit <- true
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
		func() error {
			log.Println("4th call")
			return nil
		},
		func() error {
			log.Println("5th call")
			return nil
		},
	}

	worker(functions, 4, 2)
}
