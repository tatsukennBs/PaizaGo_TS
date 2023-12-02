package main

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
)

func main() {

	//【n 行の出力】1,000 行以内の出力
	// sc := bufio.NewScanner(os.Stdin)
	// sc.Scan()
	// count, err := strconv.Atoi(sc.Text())

	// if err != nil {
	// 	fmt.Printf("数値変換エラーです：%v", err)
	// }

	// for i := 1; i <= count; i++ {
	// 	fmt.Println(i)
	// }

	// カンマ区切りで 2 つ出力
	// sc := bufio.NewScanner(os.Stdin)
	// sc.Scan()

	// var arr = strings.Split(sc.Text(), " ")

	// fmt.Printf("%v,%v", arr[0], arr[1])

	// バーティカルライン区切りで 3 つの文字列を出力 12/2 23:14-23:29
	// sc := bufio.NewScanner(os.Stdin)

	// var arr []string
	// for i := 0; i < 3; i++ {
	// 	sc.Scan()
	// 	arr = append(arr, sc.Text())
	// }

	// fmt.Printf("%v|%v|%v", arr[0], arr[1],arr[2])

	// カンマ区切りで 10 個出力 1  12/2 23:29-23:43
	// sc := bufio.NewScanner(os.Stdin)
	// sc.Scan()
	// var arr = strings.Split(sc.Text(), " ")

	// for _, v := range arr {

	// 	fmt.Printf("%v,", v)

	// }

	// カンマ区切りで 10 個出力 2  12/2 23:43-23:48
	// sc := bufio.NewScanner(os.Stdin)
	// sc.Scan()
	// var arr = strings.Split(sc.Text(), " ")

	// for i, v := range arr {

	// 	if i == len(arr)-1 {
	// 		fmt.Printf("%v", v)
	// 	} else {
	// 		fmt.Printf("%v,", v)
	// 	}
	// }

	//半角スペースとバーティカルライン区切りで 10 個出力 12/2 23:52-23:58
	sc := bufio.NewScanner(os.Stdin)

	for i := 0; i < 10; i++ {
		sc.Scan()

		if i == 9 {
			fmt.Printf("%v", sc.Text())
		} else {
			fmt.Printf("%v | ", sc.Text())
		}
	}

	//【特定の文字で区切り 1 行で出力】大きな数値を 3 けたごとにカンマ区切りで出力
	str := "1234567890"
	fmt.Println(str[0:2])

}
