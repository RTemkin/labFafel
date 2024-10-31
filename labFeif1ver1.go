package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type TreeNode struct {
	sign string
	yes  *TreeNode
	no   *TreeNode
}

type Tree struct {
	TreeHead *TreeNode
}

func create(str string) *TreeNode {
	newTreeNode := &TreeNode{
		sign: str,
		yes:  nil,
		no:   nil,
	}
	return newTreeNode
}

func (tr *Tree) atYesNo(yn string, trN *TreeNode) {

	temp1 := &trN
	yn = strings.ToLower(yn)

	if yn == "да" {
		tr.TreeHead.yes = *temp1
	} else if yn == "нет" {
		tr.TreeHead.no = *temp1
	} else {
		err := errors.New("Не верное обозначение")
		fmt.Println(err)
		fmt.Println("Исправить")
		return

	}
}

func (tr *Tree) addSign(str, str2 string) {
	temp1 := &TreeNode{str, nil, nil}
	temp2 := tr.TreeHead.yes
	temp3 := &TreeNode{str2, nil, nil}

	tr.TreeHead.yes = temp1
	temp1.no = temp2
	temp2.yes = temp3
}

func (trN *TreeNode) addNode(sig, str string) {
	newNode := &TreeNode{sig, nil, nil}

	str = strings.ToLower(str)

	if str == "да" {
		trN.yes = newNode
	} else {
		trN.no = newNode
	}

}

func (trN *TreeNode) addTwoNode(str1, str2 string) {
	newNode1 := &TreeNode{str2, nil, nil}
	newNode := &TreeNode{str1, newNode1, nil}

	trN.no = newNode

}

func (tr *Tree) playTree() bool {
	currNode := tr.TreeHead
	reader := bufio.NewReader(os.Stdin)

	for currNode != nil {
		fmt.Println(currNode.sign + "?")

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		input = strings.ToLower(input)

		if input == "да" {
			currNode = currNode.yes
		} else if input == "нет" {
			if currNode.no != nil {
				currNode = currNode.no
			} else {
				fmt.Println("Я не знаю кто это.")
				fmt.Println("расскажи про него")

				input2, err := reader.ReadString('\n')
				if err != nil {
					log.Fatal(err)
				}
				input2 = strings.TrimSpace(input2)
				input2 = strings.ToLower(input2)

				fmt.Println("Кто это?")

				input3, err := reader.ReadString('\n')
				if err != nil {
					log.Fatal(err)
				}
				input3 = strings.TrimSpace(input3)
				input3 = strings.ToLower(input3)

				currNode.addTwoNode(input2, input3)
				return false
			}
		} else {
			fmt.Println("Неверный ответ. Введите 'да' или 'нет'.")
			continue
		}

		if currNode.yes == nil {
			fmt.Println("Это", currNode.sign, "?!")

			input1, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			input1 = strings.TrimSpace(input1)
			input1 = strings.ToLower(input1)

			if input1 == "да" {
				fmt.Println("Я угадал!!!!")
				return true
			} else if currNode.no == nil {
				fmt.Println("Я не знаю кто это.")
				fmt.Println("расскажи про него")

				input2, err := reader.ReadString('\n')
				if err != nil {
					log.Fatal(err)
				}
				input2 = strings.TrimSpace(input2)
				input2 = strings.ToLower(input2)

				fmt.Println("Кто это?")

				input3, err := reader.ReadString('\n')
				if err != nil {
					log.Fatal(err)
				}
				input3 = strings.TrimSpace(input3)
				input3 = strings.ToLower(input3)

				currNode.addTwoNode(input2, input3)
				return false
			} else {
				currNode = currNode.no
			}

		}
	}
	return false
}

func main() {

	fmt.Println("Привет, давай сыграем")
	time.Sleep(time.Second)
	fmt.Println("Обучи меня и я буду угадывать животное")
	time.Sleep(time.Second)
	fmt.Println("Отвечай Да или нет, поехали")

	root := Tree{
		TreeHead: &TreeNode{
			sign: "Большое",
			yes:  &TreeNode{"Слон", nil, nil},
			no:   &TreeNode{"Мышь", nil, nil},
		},
	}

	for {
		if root.playTree() {
			break
		}
	}

}
