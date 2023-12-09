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

		empnum, err := strconv.Atoi(array[1])
		if err != nil {
			fmt.Printf("数値変換エラー: %v\n", err)
			continue
		}
		switch expression {
		case "make":
			employee = append(employee, *NewEmployee(empnum, array[2]))
		case "getnum":
			//emp.number = employee[emp.number-1].number
			//employy[i]の中身はEmploeeの型なので、要素をインスタンスのように扱い.getnum()で取得できる
			fmt.Println(employee[empnum-1].getnum())

		case "getname":
			//emp.name = employee[emp.number-1].name
			//employy[i]の中身はEmploeeの型なので、要素をインスタンスのように扱い.getnum()で取得できる
			fmt.Println(employee[empnum-1].getname())
		case "change_num":
			newnum, err := strconv.Atoi(array[2])
			if err != nil {
				fmt.Printf("数値変換エラー: %v\n", err)
				continue
			}
			employee[empnum-1].change_num(newnum)

		case "change_name":
			newname := array[2]
			employee[empnum-1].change_name(newname)
		}
	}
}

type Employee struct {
	number int
	name   string
}

// コンストラクタの作成
func NewEmployee(number int, name string) *Employee {
	return &Employee{
		number: number,
		name:   name,
	}
}

func (emp *Employee) getnum() int {
	return emp.number

}

func (emp *Employee) getname() string {
	return emp.name
}

/* 値（num,name)更新用のメソッドを追加 */
func (emp *Employee) change_num(newnum int) {
	emp.number = newnum
}

func (emp *Employee) change_name(newname string) {
	emp.name = newname
}
