package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var path string
	fmt.Println("Choose file path to be opened:")
	fmt.Scanln(&path)

	fmt.Println("Available commands:")
	fmt.Println("set {key} {value}")
	fmt.Println("get {key}")
	fmt.Println("")

	if _, error := os.Stat(path); error != nil {
		fmt.Println("ERROR>>File does not exist")
		return
	}

	for {

		input, _ := reader.ReadString('\n')

		isSetCommand, isGetCommand, error := getCommandTypeFromInput(input)
		parts := strings.Fields(input)

		if error != nil {
			fmt.Println(error)
			continue
		}

		if isSetCommand && len(parts) != 3 {
			fmt.Println("ERROR>>Set command required 3 parameters")
			continue
		} else if isGetCommand && len(parts) != 2 {
			fmt.Println("ERROR>>Get command required 2 parameters")
			continue
		}

		if isGetCommand {
			handleGet(path, parts[1])
		} else if isSetCommand {
			handleSet(path, parts[1], parts[2])
		}

	}
}

func handleGet(path string, key string) {

	var value, found, error = getKey(path, key)

	if error != nil  {
		fmt.Println("ERROR>>Unable to read data file")
	} else if !found {
		fmt.Println("ERROR>>The key is not set")
	}

	if !found {
		return
	}

	fmt.Printf("GET>>%s \n", value)
}

func handleSet(path string, key string, value string) {

	error := setKey(path, key, value)

	if error != nil  {
		fmt.Println(error)
	} else {
		fmt.Printf("SET>>%s to %s \n", key, value)
	}
}

func getCommandTypeFromInput(input string) (bool, bool, error)  {

	parts := strings.Fields(input)

	if parts[0] == "get" {
		return !is_set_command, is_get_command, nil
	} else if parts[0] == "set" {
		return is_set_command, !is_get_command, nil
	}

	return !is_set_command, !is_get_command, errors.New("unknown command type")

}

const is_set_command = true
const is_get_command = true