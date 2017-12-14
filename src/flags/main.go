package main

import (
	"fmt"
	"os"

	goflags "github.com/jessevdk/go-flags"
)

type FlagsStruct struct {
	Docs DocumentFlagsStruct `command:"docs" description:"Manage Documents"`
}

type DocumentFlagsStruct struct {
	Create CreateDocumentFlagsStruct `command:"new" description:"Create Document"` // TODO
}

type CreateDocumentFlagsStruct struct {
	Title          string            `long:"title" description:"Set Document Title"`
	AppId          string            `long:"app-id" description:"Set Document AppId" required:"true"`
	CreatedBy      string            `long:"created-by" description:"Set Document Creator" required:"true"`
	CreatedByAppId string            `long:"created-by-app-id" description:"Set Document Creator App Id" required:"true"`
	Bizdata        map[string]string `long:"bizdata" description:"Set Document Bizdata"`
	Pages          []string          `long:"pages" description:"Set Pages File Paths" required:"true"`
}

func parseFlags() error {
	var flags FlagsStruct
	parser := goflags.NewParser(&flags, goflags.HelpFlag|goflags.PassDoubleDash|goflags.IgnoreUnknown)
	parser.CommandHandler = func(command goflags.Commander, args []string) error {
		return command.Execute(args)
	}
	_, err := parser.ParseArgs(os.Args[1:])
	return err
}

// type PostCommand struct {
// 	// options or subcommands
// }

func (_ CreateDocumentFlagsStruct) Usage() string {
	// return usage info
	return "[new command options]"
}

func (command CreateDocumentFlagsStruct) Execute(args []string) error {
	fmt.Println("title is", command.Title, args)
	return nil
}

func main() {
	if err := parseFlags(); err != nil {
		fmt.Println(err)
	}
}
