package main

import "fmt"

func main() {
	println(chainAdd(5)(10))            // 15
	println(chainAdd(5)(2))             // 7
	println(chainAdd(5)(-10))           // -5
	println(chainFormat("123")("2132")) // 15
	consoleLogic("abubu")

}

func chainAdd(number int) func(number2 int) int {
	return func(number2 int) int {
		return number + number2
	}
}

func consoleLogic(text string) func(something interface{}) interface{} {
	fmt.Println(text)
	return func(something interface{}) interface{} {
		fmt.Println(something)
		return nil
	}
}

func chainFormat(text string) func(text2 string) string {
	return func(text2 string) string {
		return fmt.Sprintf("%s and %s", text, text2)
	}
}
