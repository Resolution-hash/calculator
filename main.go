package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите выражение:")

	expression, _ := reader.ReadString('\n')
	showResult(expression)
}

func showResult(expression string) {
	parts := strings.Split(expression, " ")
	if len(parts) != 3 {
		fmt.Println("Некорректный формат выражения")
		os.Exit(1)
	}

	var num1, num2 int64
	var err error

	if isRoman(strings.TrimSpace(parts[0])) {
		if !isRoman(strings.TrimSpace(parts[2])) {
			err := fmt.Errorf("Ошибка: оба числа должны быть римскими или арабскими")
			fmt.Println(err)
			os.Exit(1)
		}

		num1, err = romanToArabic(strings.TrimSpace(parts[0]))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		checkNumRange(num1)

		num2, err = romanToArabic(strings.TrimSpace(parts[2]))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		checkNumRange(num2)

		operator := parts[1]
		result := calculation(operator, num1, num2)

		fmt.Println(arabicToRoman(result))
	} else {
		num1, err = strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		checkNumRange(num1)
		num2, err = strconv.ParseInt(strings.TrimSpace(parts[2]), 10, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		checkNumRange(num2)
		operator := parts[1]

		fmt.Printf("%d", calculation(operator, num1, num2))
	}
}

func checkNumRange(num int64) {
	if num > 10 || num < 1 {
		fmt.Println("Ошибка. Допустимые числа: от 1 до 10 включительно.")
		os.Exit(1)
	}
}

func calculation(operator string, num1 int64, num2 int64) int64 {
	var result int64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			err := fmt.Errorf("Деление на ноль невозможно")
			fmt.Println(err)
			os.Exit(1)
		}
		result = num1 / num2
	default:
		err := fmt.Errorf("Неподдерживаемый оператор: %s", operator)
		fmt.Println(err)
		os.Exit(1)
	}

	return result
}

func isRoman(input string) bool {
	romanChars := "IVXLCDM"
	for _, char := range input {
		if !strings.ContainsRune(romanChars, char) {
			return false
		}
	}
	return true
}

func romanToArabic(romanNum string) (int64, error) {
	romanMap := map[rune]int{'I': 1, 'V': 5, 'X': 10}

	var result, prevValue int

	for _, char := range romanNum {
		if value, found := romanMap[char]; found {
			result += value
			if prevValue < value {
				result -= 2 * prevValue
			}
			prevValue = value
		}
	}
	return int64(result), nil
}

func arabicToRoman(arabicNum int64) string {
	if arabicNum <= 0 || arabicNum > 20 {
		err := fmt.Errorf("Число должно быть в диапазоне от 1 до 10.")
		fmt.Println(err)
		os.Exit(1)
	}

	romanMap := map[int]string{
		1:  "I",
		4:  "IV",
		5:  "V",
		9:  "IX",
		10: "X",
	}

	values := []int{10, 9, 5, 4, 1}

	romanNum := ""

	for _, value := range values {
		for arabicNum >= int64(value) {
			romanNum += romanMap[value]
			arabicNum -= int64(value)
		}
	}

	return romanNum
}
