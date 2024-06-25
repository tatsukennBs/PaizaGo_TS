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
	var inpcount = strings.Split(sc.Text(), " ")

	customercount, err := strconv.Atoi(inpcount[0])
	ordercount, err := strconv.Atoi(inpcount[1])
	if err != nil {
		fmt.Printf("数値変換エラー: %v\n", err)
		return
	}

	var customer = []Customer{}

	for i := 0; i < customercount; i++ {

		sc.Scan()
		age, err := strconv.Atoi(sc.Text())
		if err != nil {
			fmt.Printf("数値変換エラー: %v\n", err)
			return
		}

		customer = append(customer, *NewCustomer(i+1, age, 0))
		//customer = append(customer, *NewCustomer(i+1, age, 0))
	}

	//var alcflg bool
	for j := 0; j < ordercount; j++ {
		sc.Scan()
		order := strings.Split(sc.Text(), " ")
		number, err := strconv.Atoi(order[0])
		ammount, err := strconv.Atoi(order[2])
		if err != nil {
			fmt.Printf("数値変換エラー: %v\n", err)
			return
		}

		//アルコールの注文があればフラグをたてる
		// if order[1] == "alcohol" {
		// 	alcflg = true
		// }

		// if customer[number-1].age < 20 {
		// 	customer[number-1].calculate_child_total_ammount(order[1], ammount)
		// } else {
		// 	customer[number-1].calculate_total_ammount(order[1], ammount, alcflg)
		// }

		if order[1] == "food" {
			if customer[number-1].age < 20 {
				customer[number-1].takefood(ammount)
			} else {
				customer[number-1].takefood(ammount)
			}
		} else if order[1] == "softdrink" {
			customer[number-1].takeSoftDrink(ammount)

		} else if order[1] == "alcohol" {
			if customer[number-1].age >= 20 {
				customer[number-1].takeAlcohol(ammount)
			}
		}
	}

	for _, v := range customer {
		fmt.Println(v.totalammount)
	}
}

type Customer struct {
	number       int
	age          int
	totalammount int
	//alcflg       bool
}

func NewCustomer(number int, age int, totalammount int) *Customer {
	return &Customer{
		number:       number,
		age:          age,
		totalammount: totalammount,
	}
}

type CustomerAdult struct {
	Customer
	alcflg bool
}

func NewCustomerAdult(number int, age int, totalammount int, flg bool) *CustomerAdult {
	return &CustomerAdult{
		Customer{
			number:       number,
			age:          age,
			totalammount: totalammount,
		},
		alcflg: flg,
	}
}

// func (cstmr *Customer) calculate_total_ammount(order string, ammount int, alcflg bool) {

// 	//合計金額に加算する
// 	cstmr.totalammount += ammount

// 	if alcflg && order == "food" {
// 		cstmr.totalammount -= 200
// 	}
// }

// func (cstmr *Customer) calculate_child_total_ammount(order string, ammount int) {

// 	if order != "alcohol" {
// 		//合計金額に加算する
// 		cstmr.totalammount += ammount
// 	}
// }

func (cstmr *Customer) takefood(ammount int) {
	//合計金額に加算する
	cstmr.totalammount += ammount
}

func (cstmr *Customer) takeSoftDrink(ammount int) {
	//合計金額に加算する
	cstmr.totalammount += ammount
}

func (cstmr *CustomerAdult) takefood(ammount int) {
	//合計金額に加算する
	cstmr.totalammount += ammount
	if cstmr.alcflg {
		cstmr.totalammount -= 200
	}
}

func (cstmr *CustomerAdult) takeAlcohol(ammount int) {
	//合計金額に加算する
	cstmr.totalammount += ammount
	cstmr.alcflg = true
}
