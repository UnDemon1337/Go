package main

import (
	"fmt"
)

func main() {

	var task int
	_, err := fmt.Scan(&task)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	switch task {
	case -1:
		example1()
	case 1:
		task1()
	case 2:
		task2()
	case 3:
		task3()
	case 4:
		task4()
	}
}

func example1() {
	var number1 int = 1324
	fmt.Println(number1 % 10)    // 4
	fmt.Println(number1 % 100)   // 24
	fmt.Println(number1 % 1000)  // 324
	fmt.Println(number1 % 10000) // 1324

	var number2 int = 1324
	fmt.Println(number2 / 10)    // 132
	fmt.Println(number2 / 100)   // 13
	fmt.Println(number2 / 1000)  // 1
	fmt.Println(number2 / 10000) // 0
}

func task1() {
	var num int = 1024
	_, _ = fmt.Scan(&num)
	fmt.Println(num % 10)
}

func task2() {
	var num int
	_, _ = fmt.Scan(&num)
	fmt.Println(num % 100 / 10)
}

func task3() {
	var num int
	_, _ = fmt.Scan(&num)
	fmt.Println(num/100 + num/10%10 + num%10)
}

func task4() {
	var num int
	_, _ = fmt.Scan(&num)
	fmt.Print(num % 10)
	fmt.Print(num / 10 % 10)
	fmt.Print(num / 100)
}
