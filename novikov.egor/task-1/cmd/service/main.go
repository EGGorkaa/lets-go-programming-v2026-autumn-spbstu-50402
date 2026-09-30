package main

import "fmt"

func main() {
	var frst int
	_, err := fmt.Scan(&frst)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	var sec int
	_, err = fmt.Scan(&sec)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	var action string
	_, err = fmt.Scan(&action)
	if err != nil {
		fmt.Println("Invalid operation")
	}
	switch action {
	case "+":
		fmt.Println(frst + sec)
	case "-":
		fmt.Println(frst - sec)
	case "*":
		fmt.Println(frst * sec)
	case "/":
		if sec == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(frst / sec)
	default:
		fmt.Println("Invalid operation")
		return
	}
}
