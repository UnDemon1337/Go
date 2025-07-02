package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// var s1, s2, s3 int
	// countWords, err := fmt.Scan(&s1, &s2, &s3)
	// if err != nil {
	// 	fmt.Println("Error scanning input:", err)
	// 	return
	// }
	// fmt.Println("Scanned words:", countWords)
	// fmt.Println(s1, s2, s3)

	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Print("Введите строку: ")
	// _ = scanner.Scan()
	// name := scanner.Text()
	// fmt.Println("Введенная строка:", name)

	// var num float64
	// fmt.Scan(&num)
	// fmt.Println(num)

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
	case 5:
		task5()
	}
}

func task1() {
	var name string
	fmt.Scan(&name)
	fmt.Println("Привет,", name)
}

func task2() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	book := scanner.Text()
	fmt.Println(book, "- лучшая книга!")
}

func task3() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	world1 := scanner.Text()
	_ = scanner.Scan()
	world2 := scanner.Text()
	_ = scanner.Scan()
	world3 := scanner.Text()
	fmt.Println(world1)
	fmt.Println(world2)
	fmt.Println(world3)
}

func task4() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	world1 := scanner.Text()
	_ = scanner.Scan()
	world2 := scanner.Text()
	_ = scanner.Scan()
	world3 := scanner.Text()
	fmt.Println(world1)
	fmt.Println(world2)
	fmt.Println(world3)
}

func task5() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	separator := scanner.Text()
	_ = scanner.Scan()
	world1 := scanner.Text()
	_ = scanner.Scan()
	world2 := scanner.Text()
	_ = scanner.Scan()
	world3 := scanner.Text()
	fmt.Print(world1 + separator + world2 + separator + world3)
}
