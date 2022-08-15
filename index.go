package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("How old are you? ")

	reader := bufio.NewReader(os.Stdin)
	input, error := reader.ReadString('\n')

	if error != nil {
		log.Fatal(error)
	}

	input = strings.TrimSpace(input)
	age, error := strconv.ParseFloat(input, 64)

	if age >= 18 {
		fmt.Println("You are old man")
	} else {
		fmt.Println("You are young man")
	}
}
