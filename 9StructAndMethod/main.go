package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rectangle struct {
	width  float64
	height float64
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

func (r *Rectangle) Perimeter() float64 {
	return (r.width + r.height) * 2
}

func main() {
	var Param Rectangle
	var ErrCustom = errors.New("Введенное значение не является числом больше нуля")
	scanner := bufio.NewScanner(os.Stdin) // Создание нового сканера для ввода

	for {
		fmt.Print("Введите длинну прямоугольника: ")
		scanner.Scan()
		input_w := scanner.Text()
		w, err := strconv.ParseFloat(strings.TrimSpace(input_w), 64)
		if err != nil {
			fmt.Print(ErrCustom.Error())
		} else if w <= 0.0 {
			fmt.Println("Введите значение больше нуля")
		} else {
			Param.width = w
			fmt.Print("Введите ширину прямоугольника: ")
			scanner.Scan()
			input_h := scanner.Text()
			h, err := strconv.ParseFloat(strings.TrimSpace(input_h), 64)
			if err != nil {
				fmt.Print(ErrCustom.Error())
			} else if h <= 0.0 {
				fmt.Println("Введите значение больше нуля")
			} else {
				Param.height = h
				fmt.Printf("Площадь прямоугольника: %.4f\n", Param.Area())
				fmt.Printf("Периметр прямоугольника: %.4f\n", Param.Perimeter())
			}
		}
	}
}
