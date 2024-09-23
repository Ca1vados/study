package main

import (
	"errors"
	"fmt"
)

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return (r.width + r.height) * 2
}

func main() {
	var Param Rectangle
	var ErrCustom = errors.New("Введенное значение не является числом больше нуля")
	for {
		fmt.Print("Введите длинну прямоугольника: ")
		_, err := fmt.Scanln(&Param.width)
		if err != nil {
			fmt.Print(ErrCustom.Error())
		} else if Param.width <= 0 {
			fmt.Println("Введите значение больше нуля")
		} else {
			fmt.Print("Введите ширину прямоугольника: ")
			_, err := fmt.Scanln(&Param.height)
			if err != nil {
				fmt.Print(ErrCustom.Error())
			} else if Param.height <= 0 {
				fmt.Println("Введите значение больше нуля")
			} else {

				fmt.Printf("Площадь прямоугольника: %.4f\n", Rectangle.Area(Param))
				fmt.Printf("Периметр прямоугольника: %.4f\n", Rectangle.Perimeter(Param))
			}
		}
	}
}
