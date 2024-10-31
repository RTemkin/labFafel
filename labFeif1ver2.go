package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func AddMap(input1, input2 string, slace *[]map[string]string) {
	newMap := map[string]string{
		input1: input2,
	}

	*slace = append(*slace, newMap)
}

func Game(lineYes, lineNo *[]map[string]string) bool {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Я угадаю животное, загадывай его и отвечай на мои вопросы")
	time.Sleep(time.Second)

	fmt.Println("Животное большое?")

	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)

	if input == "да" {
		fmt.Println("Это Слон!?")

		input1, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input1 = strings.TrimSpace(input1)
		input1 = strings.ToLower(input1)

		if input1 == "да" {
			fmt.Println("Ура!!!!!")
			return true
		} else {
			i := 0
			for i < len(*lineYes) {
				for key := range (*lineYes)[i] {
					fmt.Println(key + "?")

					input4, err := reader.ReadString('\n')
					if err != nil {
						log.Fatal(err)
					}
					input4 = strings.TrimSpace(input4)
					input4 = strings.ToLower(input4)

					if input4 == "да" {
						fmt.Printf("Это %s !? \n", (*lineYes)[i][key])

						input5, err := reader.ReadString('\n')
						if err != nil {
							log.Fatal(err)
						}
						input5 = strings.TrimSpace(input5)
						input5 = strings.ToLower(input5)

						if input5 == "да" {
							fmt.Println("Ура!!!!!")

							return true
						}

					}
				}

				i++
			}

			fmt.Println("Я не знаю кто это")
			fmt.Println("Охарактеризуй его")

			input2, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			input2 = strings.TrimSpace(input2)
			input2 = strings.ToLower(input2)

			fmt.Println("кто это?")

			input3, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			input3 = strings.TrimSpace(input3)
			input3 = strings.ToLower(input3)

			AddMap(input2, input3, lineYes)

			fmt.Println("Я запомнил")
			time.Sleep(time.Second)

			return false
		}

	} else if input == "нет" {
		fmt.Println("Это Мышь!?")
		input1, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input1 = strings.TrimSpace(input1)
		input1 = strings.ToLower(input1)

		if input1 == "да" {
			fmt.Println("Ура!!!!!")
			return true
		} else {
			i := 0
			for i < len(*lineNo) {
				for key := range (*lineNo)[i] {
					fmt.Println(key + "?")

					input4, err := reader.ReadString('\n')
					if err != nil {
						log.Fatal(err)
					}
					input4 = strings.TrimSpace(input4)
					input4 = strings.ToLower(input4)

					if input4 == "да" {
						fmt.Printf("Это %s!? \n", (*lineNo)[i][key])

						input5, err := reader.ReadString('\n')
						if err != nil {
							log.Fatal(err)
						}
						input5 = strings.TrimSpace(input5)
						input5 = strings.ToLower(input5)

						if input5 == "да" {
							fmt.Println("Ура!!!!!")

							return true
						}

					}
				}

				i++
			}

			fmt.Println("Я не знаю кто это")
			fmt.Println("охарактеризуй его")
			input2, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			input2 = strings.TrimSpace(input2)
			input2 = strings.ToLower(input2)

			fmt.Println("кто это?")
			input3, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			input3 = strings.TrimSpace(input3)
			input3 = strings.ToLower(input3)

			AddMap(input2, input3, lineNo)

			fmt.Println("Я запомнил")
			time.Sleep(time.Second)

			return false

		}

	} else {
		fmt.Println("Не разбочиво, напишите Да или Нет")
		time.Sleep(time.Second)

	}

	return false
}

func main() {

	lineYes := make([]map[string]string, 0)
	lineNo := make([]map[string]string, 0)

	for {
		if Game(&lineYes, &lineNo) {
			break
		}
	}
}
