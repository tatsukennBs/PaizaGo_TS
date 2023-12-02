//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	nickname string
	old      string
	birth    string
	state    string
}

func main() {
	//構造体の変数宣言してそれ経由でないのと関数呼び出しできない（javaのインスタンスみたいな感じかも）
	var stu Student
	stu.UpdateStudent()
}

type UpdateInfo struct {
	changeNo   int
	changeName string
}

func (st *Student) UpdateStudent() {

	// 標準入力を取得する
	sc := bufio.NewScanner(os.Stdin)
	// 1行分の入力を取得する(クラスの人数を取得する)
	sc.Scan()
	var number = strings.Split(sc.Text(), " ")

	count, _ := strconv.Atoi(number[0])
	changeCount, _ := strconv.Atoi(number[1])

	// 人数分ループ、取得した行を空白区切で分割し配列に格納する
	studentslice := []Student{}
	for i := 0; i < count; i++ {

		//次の行を取得する
		sc.Scan()
		//テキストを取得して空白文字で分割して配列に格納する。
		var array = strings.Split(sc.Text(), " ")

		st.nickname = array[0]
		st.old = array[1]
		st.birth = array[2]
		st.state = array[3]

		studentslice = append(studentslice, *st)
	}
	//fmt.Println(studentslice)

	// 変更対象の行数分ループ、取得した行を空白区切で分割し配列に格納する。
	updateSlice := []UpdateInfo{}
	for i := 0; i < changeCount; i++ {

		//次の行を取得する
		sc.Scan()
		//テキストを取得して空白文字で分割して配列に格納する。
		var array = strings.Split(sc.Text(), " ")

		var updateInfo UpdateInfo

		updateInfo.changeNo, _ = strconv.Atoi(array[0])
		updateInfo.changeName = array[1]

		updateSlice = append(updateSlice, updateInfo)
	}

	//fmt.Println(updateSlice)

	// 更新用の配列の要素分ループ、変更スライスの番号で指定した番号の名前を更新する。
	for _, v := range updateSlice {
		studentslice[v.changeNo-1].nickname = v.changeName
	}

	for index, _ := range studentslice {
		fmt.Printf("%v %v %v %v\n", studentslice[index].nickname, studentslice[index].old, studentslice[index].birth, studentslice[index].state)
	}
}

// func (st *Student) SortStudent() {

// 	// 標準入力を取得する
// 	sc := bufio.NewScanner(os.Stdin)
// 	// 1行分の入力を取得する(クラスの人数を取得する)
// 	sc.Scan()
// 	count, _ := strconv.Atoi(sc.Text())
// 	fmt.Printf("１行目の数: %v", count)

// 	// 人数分ループ、取得した行を空白区切で分割し配列に格納する
// 	studentslice := []Student{}
// 	for i := 0; i < count; i++ {

// 		//次の行を取得する
// 		sc.Scan()
// 		//テキストを取得して空白文字で分割して配列に格納する。
// 		var array = strings.Split(sc.Text(), " ")

// 		st.nickname = array[0]
// 		st.old, _ = strconv.Atoi(array[1])
// 		st.birth = array[2]
// 		st.state = array[3]

// 		studentslice = append(studentslice, *st)
// 	}

// 	fmt.Printf("sort前: %v", studentslice)
// 	sort.Sort(StudentAge(studentslice))
// 	fmt.Printf("sort後: %v", studentslice)

// 	for index, _ := range studentslice {
// 		fmt.Printf("%v %v %v %v\n", studentslice[index].nickname, studentslice[index].old, studentslice[index].birth, studentslice[index].state)
// 	}
// }

// type StudentAge []Student

// func (arr StudentAge) Len() int {
// 	return len(arr)
// }

// func (arr StudentAge) Less(i, j int) bool {
// 	return arr[i].old < arr[j].old
// }

// func (arr StudentAge) Swap(i, j int) {
// 	arr[i], arr[j] = arr[j], arr[i]
// }

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
