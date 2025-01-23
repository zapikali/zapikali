package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Предварительно подготовим регулярное выражение:
	re := regexp.MustCompile(`^"([^"]{0,10})"\s*([\+\-\*\/])\s*(?:"([^"]{0,10})"|([1-9]|10))$`)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение в формате:")
	fmt.Println(`  "строка" + "строка"`)
	fmt.Println(`  "строка" - "строка"`)
	fmt.Println(`  "строка" * число (1..10)`)
	fmt.Println(`  "строка" / число (1..10)`)
	fmt.Println("Нажмите Ctrl+D (Linux/macOS) или Ctrl+Z (Windows) для выхода.")
	fmt.Println()

	for {
		// Считаем строку из stdin
		fmt.Print(">>> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			// Если мы встретили EOF (Ctrl+D/Ctrl+Z), прерываем цикл
			if err == io.EOF {
				fmt.Println("\nВыход.")
				break
			}
			// Любая другая ошибка чтения
			fmt.Printf("Ошибка чтения: %v\n", err)
			// Переходим на следующую итерацию, не завершая программу
			continue
		}

		// Уберём \n и возможные пробелы по краям
		input = strings.TrimSpace(input)

		// Если пользователь ввёл пустую строку — пропустим итерацию
		if input == "" {
			continue
		}

		// Пытаемся сопоставить с шаблоном
		matches := re.FindStringSubmatch(input)
		if matches == nil {
			fmt.Println("Неверный формат ввода.")
			// Переходим на следующую итерацию
			continue
		}

		firstString := matches[1]  // первая строка без кавычек
		operator := matches[2]     // оператор (+, -, * или /)
		secondString := matches[3] // вторая строка (если строка)
		secondNumber := matches[4] // число (если число)

		var result string

		switch operator {
		case "+", "-":
			// Для + или - второй операнд должен быть строкой
			if secondString == "" {
				fmt.Println("Для операций + или - вторым операндом должна быть строка в кавычках.")
				continue
			}

			if operator == "+" {
				// Конкатенация
				result = firstString + secondString
			} else {
				// Вычитание: удаляем первое вхождение secondString в firstString
				index := strings.Index(firstString, secondString)
				if index != -1 {
					result = firstString[:index] + firstString[index+len(secondString):]
				} else {
					result = firstString
				}
			}

		case "*", "/":
			// Для * или / второй операнд должен быть числом 1..10
			if secondNumber == "" {
				fmt.Println("Для операций * или / вторым операндом должно быть число 1..10.")
				continue
			}
			num, err := strconv.Atoi(secondNumber)
			if err != nil || num < 1 || num > 10 {
				fmt.Println("Вторым операндом должно быть целое число от 1 до 10.")
				continue
			}

			if operator == "*" {
				// Повторяем строку num раз
				var sb strings.Builder
				for i := 0; i < num; i++ {
					sb.WriteString(firstString)
				}
				result = sb.String()
			} else {
				// Деление: берём первые floor(len/num) символов
				length := len(firstString)
				cut := length / num
				result = firstString[:cut]
			}

		default:
			fmt.Println("Неизвестный оператор.")
			continue
		}

		// Если результат получился больше 40 символов, обрезаем и добавляем "..."
		if len(result) > 40 {
			result = result[:40] + "..."
		}

		fmt.Printf("Результат: \"%s\"\n\n", result)
	}
}
