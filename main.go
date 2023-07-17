package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkroman(x string) int { // проверка римского числа
	romarab := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	if n, ok := romarab[x]; ok {
		return n
	} else {
		return 0
	}
}

func checkoperant(x string) int { // проверка является ли введенное значение числом и соответствует ли тз
	n, err := strconv.Atoi(x)
	if err != nil {
		return 0
	} else if n < 1 || n > 10 {
		panic("число не соответствует диапазону")
	} else {
		return n
	}
}

func checoperator(x string) bool { // проверка знака на соответствие тз
	sign := [4]string{"/", "+", "-", "*"}
	for _, i := range sign {
		if x == i {
			return true
		}
	}
	return false
}

func calculate(n1, n2 int, zn string) int { // осуществление мат. операции
	var result int
	switch zn {
	case "/":
		result = n1 / n2
	case "*":
		result = n1 * n2
	case "+":
		result = n1 + n2
	case "-":
		result = n1 - n2
	}
	return result
}

func arabrom(x int) string { // Функция переводит результат вычисления функции calculate() из арабской с.с. в римскую с.с.
	arabromone := map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
	}

	arabromten := map[int]string{
		1:  "X",
		2:  "XX",
		3:  "XXX",
		4:  "XL",
		5:  "L",
		6:  "LX",
		7:  "LXX",
		8:  "LXXX",
		9:  "XC",
		10: "C",
	}

	var romnum string
	if x < 1 || x > 100 {
		panic("в римской системе нет нуля и отрицательных чисел")
	} else {
		ten := x / 10
		one := x % 10
		if ten != 0 {
			romnum += arabromten[ten]
		}
		if one != 0 {
			romnum += arabromone[one]
		}
	}
	return romnum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	incom := strings.Split(text, " ")
	if len(incom) < 3 {
		panic("не смогу посчитать, это не мат. операция")
	} else if len(incom) > 3 {
		panic("умею только *,/,+,- и считать только до 10")
	}
	num1, operator, num2 := incom[0], incom[1], incom[2]

	if checoperator(operator) == true {
		arab1 := checkoperant(num1)
		arab2 := checkoperant(num2)
		if arab1 == 0 && arab2 == 0 {
			rim1 := checkroman(num1)
			rim2 := checkroman(num2)
			if rim1 == 0 || rim2 == 0 {
				panic("римские числа не входят в диапазон")
			} else {
				fmt.Print(arabrom(calculate(rim1, rim2, operator)))
			}
		} else if (arab1 != 0 && checkroman(num2) != 0) || (checkroman(num1) != 0 && arab2 != 0) {
			panic("не умею одновременно считать римские и арабские цифры")
		} else {
			fmt.Print(calculate(arab1, arab2, operator))
		}
	} else {
		panic("цифры вижу, что делать непонимаю, проверь знак")
	}
}
