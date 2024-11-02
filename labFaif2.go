package main

import (
	"errors"
	"fmt"
	"strings"
)

func not(val string) (bool, error) {
	if val == "1" {
		return false, nil
	} else if val == "0" {
		return true, nil
	} else {
		err := errors.New("неизвестное высказывание")
		return false, err
	}
}

func And(a, b string) bool {
	if a == "1" && b == "1" {
		return true
	} else {
		return false
	}
}

func Or(a, b string) bool {
	if a == "1" || b == "1" {
		return true
	} else {
		return false
	}
}

func Impl(a, b string) bool {
	if a == "1" && b == "0" {
		return false
	} else {
		return true
	}
}

func Equi(a, b string) bool {
	if a == b {
		return true
	} else {
		return false
	}
}

func boolToString(val bool) string {
	if val {
		return "1"
	}
	return "0"
}

func StrintToBool(val string) (bool, error) {
	if val == "1" {
		return true, nil
	} else if val == "0" {
		return false, nil
	} else {
		err := errors.New("нет данных")
		return false, err
	}
}

func ApplyOperator(a, b, operator string) string {
	switch operator {
	case "And":
		return boolToString(And(a, b))
	case "Or":
		return boolToString(Or(a, b))
	case "Impl":
		return boolToString(Impl(a, b))
	case "Equi":
		return boolToString(Equi(a, b))
	default:
		return "неизвестный оператор"
	}
}

func LogicOper(comment string) string {
	tokens := strings.Fields(comment)
	values := []string{}
	stack := []string{}

	for _, token := range tokens {
		if token == "(" {
			// Начало подвыражения, добавляем в стек
			stack = append(stack, token)
		} else if token == ")" {
			// Конец подвыражения, разрешаем стек
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				operator := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if len(values) < 2 {
					return "неправильное выражение" // Ошибка при недостатке значений
				}
				valB := values[len(values)-1]
				valA := values[len(values)-2]
				values = values[:len(values)-2] // Удаляем два последних значения
				result := ApplyOperator(valA, valB, operator)
				values = append(values, result) // Добавляем результат обратно
			}
			if len(stack) == 0 {
				return "неправильное выражение" // Закрывающая скобка без открывающей
			}
			stack = stack[:len(stack)-1] // Удаляем '('
		} else if token == "not" {
			// Обработка операции NOT
			if len(values) > 0 {
				lastValue := values[len(values)-1]
				result, err := not(lastValue)
				if err != nil {
					return err.Error()
				}
				values[len(values)-1] = boolToString(result) // Заменяем на результат
			}
		} else if token == "And" || token == "Or" || token == "Impl" || token == "Equi" {
			// Обработка операторов
			for len(stack) > 0 && precedence(stack[len(stack)-1]) >= precedence(token) {
				operator := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if len(values) < 2 {
					return "неправильное выражение" // Ошибка при недостатке значений
				}
				valB := values[len(values)-1]
				valA := values[len(values)-2]
				values = values[:len(values)-2] // Удаляем два последних значения
				result := ApplyOperator(valA, valB, operator)
				values = append(values, result) // Добавляем результат обратно
			}
			stack = append(stack, token) // Добавляем оператор в стек
		} else if token == "1" || token == "0" {
			// Добавляем значения в массив
			values = append(values, token)
		}
	}

	// Обработка оставшихся операторов в стеке
	for len(stack) > 0 {
		operator := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(values) < 2 {
			return "неправильное выражение" // Ошибка при недостатке значений
		}
		valB := values[len(values)-1]
		valA := values[len(values)-2]
		values = values[:len(values)-2] // Удаляем два последних значения
		result := ApplyOperator(valA, valB, operator)
		values = append(values, result) // Добавляем результат обратно
	}

	if len(values) > 0 {
		return values[0] // Возврат результата
	}
	return "нет значений"
}

// Определение приоритета операторов
func precedence(op string) int {
	switch op {
	case "not":
		return 3
	case "And":
		return 2
	case "Or":
		return 1
	case "Impl", "Equi":
		return 0
	}
	return -1
}

func main() {

	comment := "1 And 1 Or ( 0 Or 0 ) And ( 1 And not 0 )"
	fmt.Println(StrintToBool(LogicOper(comment)))
}
