package main

import (
	"fmt"
	"time"
)

type Human interface {
	Eat(like string, dislike string)
	Sleep(startTime time.Time, endTime time.Time)
}

type Adult struct {
	name string
	age  int
}

// インターフェースを実装している
func (a *Adult) Sleep(startTime time.Time, endTime time.Time) {

	fmt.Println(startTime.Format("2006-01-02 15:04:05"))
	fmt.Println(endTime.Format("2006-01-02 15:04:05"))
}

func main() {

	//ad := Adult{"suzuki", 35}
	//ad.Sleep(time.Now(), time.Now().Add(time.Duration(time.Hour*6)))

	var input string
	fmt.Scan(&input)
}
