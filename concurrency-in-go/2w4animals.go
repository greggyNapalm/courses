package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type AnimalCommand struct {
	token0 string
	token1 string
	token2 string
}

type Cow struct {
}

func (a Cow) Eat() {
	fmt.Println("grass")
}
func (a Cow) Move() {
	fmt.Println("walk")
}
func (a Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {
}

func (a Bird) Eat() {
	fmt.Println("worms")
}
func (a Bird) Move() {
	fmt.Println("fly")
}
func (a Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
}

func (a Snake) Eat() {
	fmt.Println("mice")
}
func (a Snake) Move() {
	fmt.Println("slither")
}
func (a Snake) Speak() {
	fmt.Println("hsss")
}

func ReadStdInLine() string {
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	return strings.TrimSuffix(line, "\n")
}

func getCommand() (AnimalCommand, error) {
	//newanimal xyz bird
	//query xyz move
	fmt.Printf(">")
	tokens := strings.Fields(ReadStdInLine())
	if len(tokens) != 3 {
		return AnimalCommand{}, errors.New("unsupported amount of words/tokens in command. It should be three of them.")
	}
	if !(strings.EqualFold(tokens[0], "newanimal") || strings.EqualFold(tokens[0], "query")) {
		return AnimalCommand{}, errors.New("unsupported first token in command:`" + tokens[0] + "`. Should be `newanimal` or `query`")
	}
	return AnimalCommand{tokens[0], tokens[1], tokens[2]}, nil
}

func main() {
	animals := make(map[string]Animal)
	cow := Cow{}
	bird := Bird{}
	snake := Snake{}
	for {
		command, err := getCommand()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		if strings.EqualFold(command.token0, "newanimal") {
			switch command.token2 {
			case "cow":
				animals[command.token1] = cow
			case "bird":
				animals[command.token1] = bird
			case "snake":
				animals[command.token1] = snake
			default:
				fmt.Fprintf(os.Stderr, "Wrong animal kind name. Should be on of:cow, bird, snake")
				os.Exit(1)
			}
			fmt.Println("Created it!")
		}
		if strings.EqualFold(command.token0, "query") {
			switch command.token2 {
			case "eat":
				animals[command.token1].Eat()
			case "move":
				animals[command.token1].Move()
			case "speak":
				animals[command.token1].Speak()
			default:
				fmt.Fprintf(os.Stderr, "Wrong action name. Should be on of:eat, move, speak")
				os.Exit(1)
			}
		}
	}
}
