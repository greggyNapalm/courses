package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

type AnimalCommand struct {
	name string
	verb string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func ReadStdInLine() string {
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	return strings.TrimSuffix(line, "\n")
}

func getCommand() AnimalCommand {
	fmt.Printf(">")
	words := strings.Fields(ReadStdInLine())
	return AnimalCommand{words[0], words[1]}
}

func main() {
	animals := map[string]Animal{
		"cow":   Animal{"grass", "walk", "moo"},
		"bird":  Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}
	for {
		command := getCommand()
		switch command.verb {
		case "eat":
			animals[command.name].Eat()
		case "move":
			animals[command.name].Move()
		case "speak":
			animals[command.name].Speak()
		}
	}
}
