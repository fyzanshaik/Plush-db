package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/manifoldco/promptui"
)

func main() {
	fmt.Println("Welcome to PlushDB 🍥")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Would you like to create a database? (y/n): ")
		userInput, _, _ := reader.ReadRune()
		_, _ = reader.ReadString('\n')

		if userInput == 'y' {
			fmt.Println("Creating database!")
			break
		} else if userInput == 'n' {
			fmt.Println("Thank you for using PlushDB!")
			return
		} else {
			fmt.Println("Invalid input, please enter 'y' or 'n'.")
		}
	}

	var numFields int
	for {
		fmt.Print("Enter the number of fields (columns): ")
		numFieldsInput, _ := reader.ReadString('\n')
		numFieldsInput = numFieldsInput[:len(numFieldsInput)-1]

		numFieldsParsed, err := strconv.Atoi(numFieldsInput)
		if err != nil || numFieldsParsed <= 0 {
			fmt.Println("Please enter a valid number greater than 0.")
		} else {
			numFields = numFieldsParsed
			break
		}
	}

	fields := make(map[string]string)
	typesOfFields := []string{"string", "int", "rune", "bool"}

	for i := 0; i < numFields; i++ {
		fmt.Printf("Enter the name for field %d: ", i+1)
		fieldName, _ := reader.ReadString('\n')
		fieldName = fieldName[:len(fieldName)-1]

		prompt := promptui.Select{
			Label: "Select the type of the field:",
			Items: typesOfFields,
		}

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fields[fieldName] = result
		fmt.Printf("Field %d - Name: %s, Type: %s\n", i+1, fieldName, result)
	}

	fields["PrimaryKey"] = "int"

	fmt.Println("Final Schema:")
	for fieldName, fieldType := range fields {
		fmt.Printf("Field: %s, Type: %s\n", fieldName, fieldType)
	}

}
