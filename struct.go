package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Student struct {
	nickname string
	old      int
	birth    string
	state    string
}

func main() {
	//構造体の変数宣言してそれ経由でないのと関数呼び出しできない（javaのインスタンスみたいな感じかも）
	var stu Student
	stu.SortStudent()
}

func (st *Student) SortStudent() {

	// 標準入力を取得する
	sc := bufio.NewScanner(os.Stdin)
	// 1行分の入力を取得する(クラスの人数を取得する)
	sc.Scan()
	count, _ := strconv.Atoi(sc.Text())
	fmt.Printf("１行目の数: %v", count)

	// 人数分ループ、取得した行を空白区切で分割し配列に格納する
	studentslice := []Student{}
	for i := 0; i < count; i++ {

		//次の行を取得する
		sc.Scan()
		//テキストを取得して空白文字で分割して配列に格納する。
		var array = strings.Split(sc.Text(), " ")

		st.nickname = array[0]
		st.old, _ = strconv.Atoi(array[1])
		st.birth = array[2]
		st.state = array[3]

		studentslice = append(studentslice, *st)
	}

	fmt.Printf("sort前: %v", studentslice)
	sort.Sort(StudentAge(studentslice))
	fmt.Printf("sort後: %v", studentslice)

	for index, _ := range studentslice {
		fmt.Printf("%v %v %v %v\n", studentslice[index].nickname, studentslice[index].old, studentslice[index].birth, studentslice[index].state)
	}
}

type StudentAge []Student

func (arr StudentAge) Len() int {
	return len(arr)
}

func (arr StudentAge) Less(i, j int) bool {
	return arr[i].old < arr[j].old
}

func (arr StudentAge) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// func (st *Student) SearchStudent() {

// 	sc := bufio.NewScanner(os.Stdin)
// 	//1行分の入力を取得する(クラスの人数を取得する)
// 	sc.Scan()
// 	count, _ := strconv.Atoi(sc.Text())

// 	//mapを作成
// 	//student := make(map[string]string)

// 	studentslice := []Student{}

// 	for i := 0; i < count; i++ {
// 		//1人分の情報を取得する。
// 		sc.Scan()
// 		//取得した情報を配列に格納する
// 		var array = strings.Split(sc.Text(), " ")

// 		st.nickname = array[0]
// 		st.old = array[1]
// 		st.birth = array[2]
// 		st.state = array[3]

// 		studentslice = append(studentslice, *st)
// 	}

// 	sc.Scan()
// 	searchage := sc.Text()

// 	//fmt.Println(studentslice)
// 	for index, _ := range studentslice {

// 		if studentslice[index].old == searchage {
// 			fmt.Println(studentslice[index].nickname)
// 		}

// 		// fmt.Println("User{")
// 		// fmt.Println("nickname : " + studentslice[index]["nickname"])
// 		// fmt.Println("old : " + studentslice[index]["old"])
// 		// fmt.Println("birth : " + studentslice[index]["birth"])
// 		// fmt.Println("state : " + studentslice[index]["state"])
// 		// fmt.Println("}")
// 	}
// }
