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
	fmt.Println("Simple Calculator Shell")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", " ", -1)

		inputs := strings.Fields(text)
		input1 := inputs[0]
		input2 := inputs[1]
		input3 := inputs[2]

		in1, err1 := strconv.Atoi(input1)
		in3, err3 := strconv.Atoi(input3)
		if err1 != nil {
			// handle error
			fmt.Println(err1)
			os.Exit(2)
		}
		if err3 != nil {
			// handle error
			fmt.Println(err3)
			os.Exit(2)
		}

		if strings.Contains(input2, "+") {
			result := in1 + in3
			fmt.Println(text, " = ", result)
		}

		if strings.Contains(input2, "-") {
			result := in1 - in3
			fmt.Println(text, " = ", result)
		}

		if strings.Contains(input2, "*") {
			result := in1 * in3
			fmt.Println(text, " = ", result)
		}

		if strings.Contains(input2, "x") {
			result := in1 * in3
			fmt.Println(text, " = ", result)
		}

		if strings.Contains(input2, "X") {
			result := in1 * in3
			fmt.Println(text, " = ", result)
		}

		if strings.Contains(input2, "/") {
			result := in1 / in3
			fmt.Println(text, " = ", result)
		}

	}

}
