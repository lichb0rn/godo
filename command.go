package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Edit   string
	Delete int
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cmd := CmdFlags{}

	flag.StringVar(&cmd.Add, "add", "", "Add a todo with a title")
	flag.StringVar(&cmd.Edit, "edit", "", "Edit a todo by index with a new title")
	flag.IntVar(&cmd.Delete, "delete", -1, "Delete a todo by index")
	flag.IntVar(&cmd.Toggle, "toggle", -1, "Toggle a todo by index")
	flag.BoolVar(&cmd.List, "list", false, "List all todos")

	flag.Parse()

	return &cmd
}

func (cmd *CmdFlags) Execute(todos *Todos) {
	switch {
	case cmd.Add != "":
		todos.Add(cmd.Add)
	case cmd.Edit != "":
		parts := strings.SplitN(cmd.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid format for edit. Please use id:new_title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error, invalid index value. Please use a number")
			os.Exit(1)
		}

		todos.Edit(index, parts[1])
	case cmd.Delete != -1:
		todos.Delete(cmd.Delete)
	case cmd.Toggle != -1:
		todos.Toggle(cmd.Toggle)
	case cmd.List:
		todos.Print()

	default:
		fmt.Println("Error, no command provided")
	}
}
