package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	calculate()
}

func fromRoman(Roman string) int {
	RomNum := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for i := 0; i < 10; i++ {
		if RomNum[i] == Roman {
			return i + 1
		}
	}
	return 0
}

func toRoman(num int) string {
	values := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	roman := ""
	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			num -= values[i]
			roman += romans[i]
		}
	}
	return roman
}

func detect(expr string) (string, string, string, int) {
	arabicPattern := `^([1-9]|10) ([+\-*/]) ([1-9]|10)$`
	romanPattern := `^(X|IX|IV|V?I{0,3}) ([+\-*/]) (X|IX|IV|V?I{0,3})$`
	arabicRe := regexp.MustCompile(arabicPattern)
	romanRe := regexp.MustCompile(romanPattern)
	if arabicRe.MatchString(expr) {
		matches := arabicRe.FindStringSubmatch(expr)
		return matches[1], matches[2], matches[3], 1
	} else if romanRe.MatchString(expr) {
		matches := romanRe.FindStringSubmatch(expr)
		return matches[1], matches[2], matches[3], 2
	}
	return "", "", "", 0
}

func input() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = input[:len(input)-1]
	return input
}

func process(a, b int, c string) int {
	var res int
	if c == "+" {
		res = a + b
	} else if c == "-" {
		res = a - b
	} else if c == "*" {
		res = a * b
	} else if c == "/" && b != 0 {
		res = a / b
	} else {
		panic("Деление на ноль")
	}
	return res
}

func calculate() {
	var number1, number2 int
	str := input()
	num1, operat, num2, res := detect(str)
	if res == 1 {
		number1, _ = strconv.Atoi(num1)
		number2, _ = strconv.Atoi(num2)
	} else if res == 2 {
		number1 = fromRoman(num1)
		number2 = fromRoman(num2)
	} else {
		panic("Неверный ввод")
	}
	result := process(number1, number2, operat)
	if res == 1 {
		fmt.Println(result)
	} else if result > 0 {
		fmt.Println(toRoman(result))
	} else {
		panic("Невозможный ответ")
	}
}
