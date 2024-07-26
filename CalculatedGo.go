package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func add(x, y float64) float64 {
	return x + y
}

func subtract(x, y float64) float64 {
	return x - y
}

func multiply(x, y float64) float64 {
	return x * y
}

func divide(x, y float64) float64 {
	if y != 0 {
		return x / y
	}
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Вы можете в любой момент завершить процесс нажав '!'.")

	var (
		num1     float64
		num2     float64 // Объявление переменных вне цикла
		operator string
	)

	for {
		// Ввод первого числа
		for {
			fmt.Print("Введите первое число: ")
			input, _ := reader.ReadString('\n')
			input = input[:len(input)-1]
			if input == "!" {
				fmt.Println("Процесс завершен")
				return
			}

			//проверка числа
			var err error
			num1, err = strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("Неверно. Пожалуйста, введите корректное число.")
			} else {
				break
			}
		}

		// Ввод оператора
		for {
			fmt.Print("Введите оператор (+, -, *, /): ")
			input, _ := reader.ReadString('\n')
			input = input[:len(input)-1]
			if input == "!" {
				fmt.Println("Процесс завершен")
				return
			}

			operator = input
			if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
				fmt.Println("Неверный оператор. Пожалуйста, введите +, -, *, или /.")
			} else {
				break
			}
		}

		// Ввод второго числа
		for {
			fmt.Print("Введите второе число: ")
			input, _ := reader.ReadString('\n')
			input = input[:len(input)-1]
			if input == "!" {
				fmt.Println("Процесс завершен")
				return
			}

			var err error
			num2, err = strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("Неверно. Пожалуйста, введите корректное число.")
			} else {
				break
			}
		}

		// Вычисление результата
		var result float64
		switch operator {
		case "+":
			result = add(num1, num2)
		case "-":
			result = subtract(num1, num2)
		case "*":
			result = multiply(num1, num2)
		case "/":
			result = divide(num1, num2)
		}

		fmt.Printf("Результат: %.2f\n", result)

		// Проверка продолжения
		fmt.Print("Хотите продолжить? 1.(да) / 2.(нет): ")
		input, _ := reader.ReadString('\n')
		input = input[:len(input)-1]
		if input != "1" {
			fmt.Print("Процесс завершен")
			break
		}
	}
}
