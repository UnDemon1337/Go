package main

import (
	"fmt"
	"strconv"
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
	case 0:
		example2()
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
	case 10:
		task10()
	case 11:
		task11()
	}
}

func example1() {

	var str string = "5"
	convertNum, _ := strconv.Atoi(str)
	fmt.Println(convertNum * 2) // вывод 10

	var num int = 5
	convertStr := strconv.Itoa(num)
	fmt.Println(convertStr + convertStr) // вывод 55

	var floatNum1 float64 = 6.6
	intNum1 := int(floatNum1) // приведение (casting) вещественного числа к целому
	fmt.Println(intNum1)

	var intNum2 int = 6
	floatNum2 := float64(intNum2) // приведение (casting) целого числа к вещественному
	fmt.Println(floatNum2 + 0.4)  // Вывод 6.4

	a := 5
	b := 2
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)

	fmt.Println(12 / 3)
	fmt.Println(12 / 4)
	fmt.Println(12 / 5)
	fmt.Println(12 / 6)
	fmt.Println(12 / 7)
	fmt.Println(12 / 12)
	fmt.Println(12 / 14)

	fmt.Println(12 % 3)
	fmt.Println(12 % 4)
	fmt.Println(12 % 5)
	fmt.Println(12 % 6)
	fmt.Println(12 % 7)
	fmt.Println(12 % 12)
	fmt.Println(12 % 14)

	var num1 int = 2 + 2*2
	fmt.Println(num1)

	var num2 int = (2 + 2) * 2
	fmt.Println(num2)

	someVar := 1                                             // Инициализируем переменную someVar значением 1.
	fmt.Println("Исходное значение:", someVar)               // Выводим начальное значение: 1.
	someVar = -someVar                                       // Применяем унарный минус, превращая 1 в -1.
	fmt.Println("Из положительного отрицательное:", someVar) // Выводим: -1.
	someVar = -someVar                                       // Ещё раз применяем унарный минус. Теперь -1 становится 1.
	fmt.Println("Из отрицательного положительное:", someVar) // Выводим: 1.
}

func task1() {
	var num int
	fmt.Scan(&num)
	fmt.Println(num * num)
}

func task2() {
	var num int
	fmt.Scan(&num)
	fmt.Println(num / 1000)
}

func task3() {
	var x int
	fmt.Scan(&x)
	x2 := x * x   // Возводим число в квадрат
	x4 := x2 * x2 // Возводим число в четвертую степень
	x6 := x4 * x2 // Умножаем число x4 на x2, получаем число в шестой степени
	fmt.Println(x6)
}

func task4() {
	var num1, num2, num3 int
	fmt.Scan(&num1, &num2, &num3)
	fmt.Println(num1 * num2 * num3)
}

func task5() {
	var num1, num2 uint16
	fmt.Scan(&num1, &num2)
	fmt.Println(num2 / num1)
}

func task6() {
	var num1, num2 uint16
	fmt.Scan(&num1, &num2)
	fmt.Println(num2 % num1)
}

func task7() {
	var num int
	fmt.Scan(&num)
	fmt.Println(num)
	num++
	fmt.Println(num)
	num++
	fmt.Println(num)
}

func task8() {
	var num1, num2, num3, num4 int
	fmt.Scan(&num1, &num2, &num3, &num4)
	fmt.Println(3 * (num1 + num2 + num3 + num4))
}

func task9() {
	var a, b, n int
	fmt.Scan(&a, &b, &n)
	fmt.Println((n*(a*100+b))/100, (n*(a*100+b))%100)
}

func task10() {
	var min int
	fmt.Scan(&min)
	fmt.Println(min, "мин - это", min/60, "час", min%60, "минут.")
}

func task11() {
	var num int
	fmt.Scan(&num)
	fmt.Println("Следующее за числом", num, "число:", num+1)
	fmt.Println("Для числа", num, "предыдущее число:", num-1)
}

func example2() {
	num1 := 5
	str1 := strconv.Itoa(num1)
	str2 := "5"
	num2, _ := strconv.Atoi(str2)
	fmt.Println(str1, num2)
}
