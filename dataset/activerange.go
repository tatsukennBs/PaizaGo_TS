//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//1問目
	// sc := bufio.NewScanner(os.Stdin)

	// // 1行目の数値を取得する
	// sc.Scan()
	// var numArray = strings.Split(sc.Text(), " ")
	// number, _ := strconv.Atoi(numArray[1])

	// sc.Scan()
	// var array = strings.Split(sc.Text(), " ")
	// fmt.Println(array[number-1])

	//2問目(20min)
	// sc := bufio.NewScanner(os.Stdin)

	// // 1行目の数値を取得する（使用しない）
	// sc.Scan()
	// // 2行目の数値列を取得して配列に格納する
	// sc.Scan()
	// var numArray = strings.Split(sc.Text(), " ")
	// //3行目の数値を取得する（使用しない）
	// sc.Scan()
	// //4行目の数値列を取得して配列に格納する
	// sc.Scan()
	// var selectArray = strings.Split(sc.Text(), " ")

	// for _, v := range selectArray {

	// 	index, _ := strconv.Atoi(v)
	// 	fmt.Printf("%v\n", numArray[index-1])
	// }

	//3問目(1h)
	// sc := bufio.NewScanner(os.Stdin)

	// // 数列を取得する
	// sc.Scan()

	// var numArray = strings.Split(sc.Text(), " ")

	// //ソートインターフェースを実装したsortint型として定義する
	// var numArrayInt sortint

	// //分割後配列から取得した値を数値変換してint型配列に格納する
	// for _, v := range numArray {

	// 	num, err := strconv.Atoi(v)
	// 	if err != nil {
	// 		fmt.Printf("数値変換エラー: %v\n", err)
	// 	}
	// 	numArrayInt = append(numArrayInt, num)
	// }

	// fmt.Printf("sort前：%v", numArrayInt)
	// sort.Sort(numArrayInt)
	// fmt.Printf("sort後：%v", numArrayInt)

	// answer := numArrayInt[len(numArrayInt)-1] - numArrayInt[0]
	// fmt.Printf("%v", answer)

	//4問目
	sc := bufio.NewScanner(os.Stdin)

	// 数列を取得する
	sc.Scan()

	var numArray = strings.Split(sc.Text(), " ")

	// 操作回数を取得して数値に変換する
	executeCnt, err := strconv.Atoi(numArray[1])
	if err != nil {
		fmt.Printf("数値変換エラー: %v\n", err)
	}

	// 2行目を取得する（デフォルトの数列）
	sc.Scan()
	var resultArray = strings.Split(sc.Text(), " ")

	// 操作回数分ループする
	for i := 0; i < executeCnt; i++ {
		//3行目以降を取得する
		sc.Scan()
		var executeArray = strings.Split(sc.Text(), " ")

		fmt.Println(&resultArray)
		fmt.Println(executeArray)
		Execute(&resultArray, executeArray)

	}

}

func Execute(resultArray *[]string, executeArray []string) {

	switch executeArray[0] {
	case "0":
		fmt.Println("push back")
		*resultArray = append(*resultArray, executeArray[1])

	case "1":
		fmt.Println("pop back")
		if len(*resultArray) > 0 {
			*resultArray = (*resultArray)[:len(*resultArray)-1]
		}

	case "2":
		fmt.Println("print start")
		for index, value := range *resultArray {

			if index == len(*resultArray)-1 {
				fmt.Printf("%v", value)
				fmt.Printf("\n")
			} else {
				fmt.Printf("%v ", value)
			}
		}
	}
}

// エイリアスとして定義
// type sortint []int

// func (arr sortint) Len() int {
// 	return len(arr)
// }

// // sortのやり方を規定する
// func (arr sortint) Less(i, j int) bool {
// 	return arr[i] < arr[j]
// }

// func (arr sortint) Swap(i, j int) {
// 	arr[i], arr[j] = arr[j], arr[i]
// }
