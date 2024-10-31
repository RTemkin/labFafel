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
	operators := []string{}

	for _, token := range tokens {
		if token == "1" || token == "0" {
			values = append(values, token)
		} else if token == "not" {
			// Обработка операции NOT
			if len(values) > 0 {
				lastValue := values[len(values)-1]
				result, err := not(lastValue)
				if err != nil {
					return err.Error()
				}
				values = values[:len(values)-1]               // Удаляем последнее значение
				values = append(values, boolToString(result)) // Добавляем результат
			}
		} else {
			operators = append(operators, token)
		}
	}

	// Применение операторов к значениям
	for i := 0; i < len(operators); i++ {
		operator := operators[i]
		if i < len(values)-1 {
			valA := values[i]
			valB := values[i+1]
			result := ApplyOperator(valA, valB, operator)
			values[i] = result                             // Записываем результат на место операндов
			values = append(values[:i+1], values[i+2:]...) // Удаляем второй операнд
			i--                                            // Приводим i в соответствие с изменённым массивом
		}
	}

	if len(values) > 0 {
		return values[0] // Возврат результата
	}
	return "нет значений"
}

func main() {
	comment := "1 And 1 Or 1 And 0 not"
	fmt.Println(LogicOper(comment))
}
