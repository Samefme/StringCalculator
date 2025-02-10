package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Разделяем входную строку на 3 переменные
func op(input string) (left, right, znac string) {
	runes := []rune(input)
	leftRunes := []rune{}
	rightRunes := []rune{}
	var operatorFound bool

	for i, char := range runes {
		if i == 0 && runes[i] != '"' {
			panic("Первым аргументом не должно быть число. Так же аргумент должен передаваться в кавычках!")
		}
		if char == '+' || char == '-' || char == '*' || char == '/' {
			if runes[i-1] == ' ' && runes[i+1] == ' ' {
				znac = string(char)
				operatorFound = true
				continue
			}
		}
		if !operatorFound {
			leftRunes = append(leftRunes, char)
		} else {
			rightRunes = append(rightRunes, char)
		}
	}

	one := strings.TrimSpace(string(leftRunes))
	two := strings.TrimSpace(string(rightRunes))

	left = strings.ReplaceAll(one, "\"", "")
	right = strings.ReplaceAll(two, "\"", "")

	if len(left) > 10 || len(right) > 10 {
		panic("Строка не должна превышать 10 символов!")
	}

	return left, right, znac
}

// Суммируем
func sum(left, right string) string {

	return left + right
}

// Вычитаем
func minus(left, right string) string {

	return strings.Replace(left, right, "", 1)
}

// Умножаем
func mult(left, right string) string {

	number, err := strconv.Atoi(right)
	if err != nil {
		panic("При уменожении, второй аргумент должен быть целым числом!")
	}

	if number < 1 || number > 10 {
		panic("Число не входит в диапазон от 1 до 10!")
	}

	return strings.Repeat(left, number)
}

// Делим
func div(left, right string) string {

	number, err := strconv.Atoi(right)
	if err != nil {
		panic("При делении, второй аргумент должен быть целым числом!")
	}

	if number < 1 || number > 10 {
		panic("Число не входит в диапазон от 1 до 10!")
	}

	out := len(left) / number

	return left[:out]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введите данные:")

	for scanner.Scan() {
		input := scanner.Text()

		left, right, znac := op(input)

		// Выбираеам действие в зависимости от знака между переменными
		var result string

		switch znac {
		case "+":
			result = sum(left, right)
		case "-":
			result = minus(left, right)
		case "*":
			result = mult(left, right)
		case "/":
			result = div(left, right)
		}
		// Если результат вычисления больше 40, в конце добавляется ...
		if len(result) > 40 {
			result = result[:40] + "..."
		}

		fmt.Printf("\"%s\"\n", result)
	}

}
