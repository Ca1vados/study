package main

import (
	"errors"
	"fmt"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func main() {
	var dig int
	var ErrCustom = errors.New("Введенное значение не является целым неотрицательным числом")
	for {
		fmt.Print("Введите целое неотрицательное число для вычисления факториала или -1 для выхода из программы: ")
		_, err := fmt.Scan(&dig)
		if err != nil {
			fmt.Print(ErrCustom.Error())
			continue
		} else if dig < -1 {
			fmt.Println("Введите число больше нуля")
		} else if dig == -1 {
			break
		} else {
			fmt.Printf("Факториал числа %d: %d\n", dig, factorial(dig))
		}
	}
}
