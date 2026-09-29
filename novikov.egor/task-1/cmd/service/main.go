package main

import "fmt"

func main() {
	var frst int
	_, error := fmt.Scan(&frst)
	if error != nil {
		fmt.Println("Invalid first operand")
		return
	}
	var sec int
	_, error = fmt.Scan(&sec)
	if error != nil {
		fmt.Println("Invalid second operand")
		return
	}
	var action string
	_, error = fmt.Scan(&action)
	if error != nil {
		fmt.Println("Invalid operation")
	}
	if action == "+" {
		fmt.Println(frst + sec)
	} else if action == "-" {
		fmt.Println(frst - sec)
	} else if action == "*" {
		fmt.Println(frst * sec)
	} else if action == "/" {
		if sec == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(frst / sec)
	}
}
