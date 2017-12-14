package main

import (
	"fmt"
	"os"

	goflags "github.com/jessevdk/go-flags"
)

type AddCommand struct {
	All bool `short:"a" long:"all" description:"Add all files"`
}

var addCommand AddCommand

func (x *AddCommand) Execute(args []string) error {
	fmt.Printf("Adding (all=%v): %#v\n", x.All, args)
	return nil
}

// func aainit() {
// 	parser.AddCommand("add",
// 		"Add a file",
// 		"The add command adds a file to the repository. Use -a to add all files.",
// 		&addCommand)
// }

func main() {
	parser := goflags.NewParser(&addCommand, goflags.HelpFlag|goflags.PassDoubleDash|goflags.IgnoreUnknown)
	parser.AddCommand("add",
		"Add a file",
		"The add command adds a file to the repository. Use -a to add all files.",
		&addCommand)
	parser.CommandHandler = func(command goflags.Commander, args []string) error {
		return command.Execute(args)
	}
	parser.ParseArgs(os.Args[1:])
}
