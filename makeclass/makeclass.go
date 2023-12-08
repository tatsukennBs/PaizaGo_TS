package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// 1行目の数値を取得する
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	count, err := strconv.Atoi(sc.Text())

	if err != nil {
		fmt.Printf("数値変換エラー: %v\n", err)
		return
	}

	var employee = []Employee{}

	for i := 0; i < count; i++ {
		//1行読み込み
		sc.Scan()
		var array = strings.Split(sc.Text(), " ")

		expression := array[0]

		var emp Employee
		emp.number, err = strconv.Atoi(array[1])
		if err != nil {
			fmt.Printf("数値変換エラー: %v\n", err)
			continue
		}
		switch expression {
		case "make":
			emp.name = array[2]
			employee = append(employee, emp)

		case "getnum":
			//emp.number = employee[emp.number-1].number
			fmt.Println(employee[emp.number-1].getnum())

		case "getname":
			//emp.name = employee[emp.number-1].name
			fmt.Println(employee[emp.number-1].getname())
		}
	}
}

type Employee struct {
	number int
	name   string
}

func (emp *Employee) getnum() int {
	return emp.number

}

func (emp *Employee) getname() string {
	return emp.name
}
