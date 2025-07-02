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
	case 5:
		task5()
	case 6:
		task6()
	case 7:
		task7()
	case 8:
		task8()
	case 9:
		task9()
	}
}

func example1() {
	var d float64 = 2.9
	fmt.Println(int(d)) // 2

	var n float64 = -9
	fmt.Println(math.Abs(n)) // берет модуль числа и выводит на экран число 9

	var n1 float64 = 8
	fmt.Println(math.Sqrt(n1))

	// ошибка!
	// var n2 float64 = 8
	// var b2 int
	// b2 = math.Sqrt(n2)

	var n2 float64 = 8
	fmt.Println(math.Pow(n2, 2))

	// var n float64 = 8
	// var b int
	// b = math.Pow(n, 2) // ошибка! Результат записывается в целочисленную переменную!

	var n3 float64 = 8
	fmt.Println(math.Max(n3, 4)) // выводит на экран число 8

	fmt.Println(math.Min(n3, 4))
}

func task1() {
	var r float64
	_, _ = fmt.Scan(&r)
	fmt.Println(3.14 * r * r)
}

func task2() {
	var a, b float64
	_, _ = fmt.Scan(&a, &b)
	fmt.Println(0.5 * a * b)
}

func task3() {
	var f float64
	_, _ = fmt.Scan(&f)
	fmt.Println((f - 32.0) * 5 / 9)
}

func task4() {
	var a, b float64
	_, _ = fmt.Scan(&a, &b)
	fmt.Println((a + b) / 2)
}

func task5() {
	var num float64
	_, _ = fmt.Scan(&num)
	fmt.Println(num - float64(int(num)))
}

func task6() {
	var num float64
	_, _ = fmt.Scan(&num)
	fmt.Println(num / math.Pow(2, 13))
}

func task7() {
	var a, b float64
	_, _ = fmt.Scan(&a, &b)
	fmt.Println(a + b + math.Sqrt(math.Pow(a, 2)+math.Pow(b, 2)))
}

func task8() {
	var x1, x2, y1, y2 float64
	_, _ = fmt.Scan(&x1, &y1, &x2, &y2)
	fmt.Println(math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2)))
}

func task9() {
	var x1, x2, y1, y2 float64
	_, _ = fmt.Scan(&x1, &y1, &x2, &y2)
	fmt.Println(math.Abs(x1-x2) + math.Abs(y1-y2))
}
