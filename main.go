package main

import (
	"bufio"
	"fmt"
	"os"
)

const ResetColor = "\033[0m"
const Red = "\033[31m"
const Green = "\033[32m"
const Bold = "\033[1m"
const ResetBold = "\033[22m"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	question1(scanner)
	question2(scanner)
	question3(scanner)
}

func question1(scanner *bufio.Scanner) {
	// todo: add code example
	println(Red + "Когда инициализируются глобальные переменные уровня пакета?" + ResetColor)
	println("1. Во время инициализации пакета: перед вызовом функции main()")
	println("2. Во время инициализации пакета: перед вызовом функции init()")
	println("3. При первом обращении к переменной")

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			println("Sad to see you go!")
			break
		}
		if text == "2" {
			println("Congratz!")
			break
		}
		println("Wrong answer! Try again!")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}

func question2(scanner *bufio.Scanner) {
	// todo: add code example
	const q = Red + "В каком порядке произойдет инициализация переменных:" + Green + `
	package main

	var a = 1
	var b = a + c
	var c = f()
	var d = 2

	func f() {
		...
	}

	func main() {
		...
	}
	` + ResetColor
	println(q)
	println("1. a,b,c,d")
	println("2. a,d,c,b")
	println("3. a,c,b,d")
	println("4. a,c, для остальных переменных порядок не определен")

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			println("Sad to see you go!")
			break
		}
		if text == "3" {
			println("Correct!")
			break
		}
		println("Wrong! Try again!")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}

func question3(scanner *bufio.Scanner) {
	q := Red + `Как работет инструкция defer? В какой момент расчитываются значения парамеров функции при отложенном вызове? Что выведет этот код?` + Green +
		`for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}` + ResetColor
	println(q)
	println("1. 4, 3, 2, 1, 0")
	println("2. 0, 1, 2, 3, 4")
	println("3. 0")
	println("4. 4")
	println("5. 4, 4, 4, 4")

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			println("Sad to see you go!")
			break
		}
		if text == "1" {
			println("Correct!")
			break
		}
		println("Wrong! Try again!")
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}
