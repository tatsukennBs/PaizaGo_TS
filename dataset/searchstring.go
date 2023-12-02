package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//問題１、２　19:54-20:00
	sc := bufio.NewScanner(os.Stdin)

	// 1行目の数値を取得する
	sc.Scan()

	fmt.Println(sc.Text())

	//問題３　20:00-XXX

}
