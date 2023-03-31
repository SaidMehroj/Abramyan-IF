package main

import "fmt"

func If28(ans int) {
	if ans%100 == 0 && ans%400 != 0 {
		fmt.Println("Невысокосный год")
	} else if ans%4 == 0 {
		fmt.Println("Высокосный год")
	} else {
		fmt.Println("Невысокосный год")
	}
}

func If29(ans int) {
	if ans == 0 {
		fmt.Println("Нулевое число")
	} else if ans%2 == 0 && ans > 0 {
		fmt.Println("Четное положительное число")
	} else if ans%2 == 0 && ans < 0 {
		fmt.Println("Четное отрицательное число")
	} else if ans%2 != 0 && ans > 0 {
		fmt.Println("Нечетное положительное число")
	} else if ans%2 != 0 && ans < 0 {
		fmt.Println("Нечетное отрицательное число")
	}

}

func If30(ans int) {
	if ans%2 == 0 && ans >= 100 {
		fmt.Println("Четное трёхзначное число")
	} else if ans%2 != 0 && ans >= 100 {
		fmt.Println("Нечетное трёхзначное число")
	} else if ans%2 == 0 && ans >= 10 && ans < 100 {
		fmt.Println("Четное двухзначное число")
	} else if ans%2 != 0 && ans >= 10 && ans < 100 {
		fmt.Println("Нечетное двухзначное число")
	} else if ans%2 == 0 && ans < 10 {
		fmt.Println("Четная цифра")
	} else if ans%2 != 0 && ans < 10 {
		fmt.Println("Нечетная цифра")
	}
}
