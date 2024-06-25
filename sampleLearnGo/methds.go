package main

import (
	"fmt"
)

func main() {

	//メソッド呼び出しは*変数で指定しなくてもメソッド定義に準じて変換が行われる
	r := rect{width: 100, height: 5}
	fmt.Println("area:", r.area())  //500
	fmt.Println("area:", r.perim()) //2000

	rp := &r
	fmt.Println("area:", rp.area())  // 500　(*rp.area()と同義)
	fmt.Println("area:", rp.perim()) // 2000　(*rp.perim()と同義)

}

type rect struct {
	width  int
	height int
}

// ポインタレシーバ
func (r *rect) area() int {
	return r.width * r.height
}

// 値レシーバ
func (r rect) perim() int {
	return 2 * r.width * 2 * r.height
}
