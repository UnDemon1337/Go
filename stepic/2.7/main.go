package main

import (
	"fmt"
	"math"
)

func main() {

	var task int
	_, err := fmt.Scan(&task)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	switch task {
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

func task1() {
	var num1, num2, num3 int
	fmt.Scan(&num1, &num2, &num3)
	fmt.Println(num1+num2+num3, num1*num2*num3)
}

func task2() {
	var a int
	_, _ = fmt.Scan(&a)
	fmt.Println(a / 100 % 10)
}

func task3() {
	var a int
	_, _ = fmt.Scan(&a)
	fmt.Println(a/100%10 + a/10%10 + a%10)
}

func task4() {
	var a, b float64
	fmt.Scan(&a, &b)
	fmt.Println(math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2)))
}
