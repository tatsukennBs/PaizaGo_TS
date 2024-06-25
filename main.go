package main

import (
	"fmt"
	"time"
)

type Result struct {
	Error error
}

func f(from string, cancel chan struct{}) chan Result {
	resultChan := make(chan Result, 1)

	go func() {
		for {
			select {
			case <-cancel:
				fmt.Println("goroutine処理終了")
				resultChan <- Result{nil}
				close(resultChan)
				return

			default:
				fmt.Println("loop処理実行")
				time.Sleep(1 * time.Second)
				fmt.Println("looplast")
				fmt.Println(time.Now())

				if from == "goroutineErr" {
					resultChan <- Result{fmt.Errorf("Error: %s", ",エラーですよ")}
					close(resultChan)
					return
				}
			}
		}
	}()
	return resultChan
}

func main() {

	cancel := make(chan struct{})
	result := f("goroutine", cancel)
	fmt.Println(result)

	go func() {
		for v := range result {
			if v.Error != nil {
				fmt.Printf("error: %v\n", v.Error)
				continue
			}
			// fmt.Printf("Response: %v\n", "aaa")
		}
	}()

	time.Sleep(10 * time.Second)
	close(cancel)

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
