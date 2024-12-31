package main

import (
	"flag" //for parsing cli arguments
	"fmt"
	"os"
)

type CmdFlags struct { //all the commands
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
	LeftQuestsList bool
}

func NewCmdFlags() *CmdFlags { //returns pointer to a cmdFlags struct with the values
	cf := CmdFlags{} //initiate empty struct

	//define the flags:
	//"add" is the cli flag, "" is the default value and next is the flag description
	//can see them with -help
	flag.StringVar(&cf.Add, "add", "", "Add a new quest with \"title:urgency:due date\" (urgency value 0 to 5)")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a quest by index & other infos like, \"index:title:urgency:dueDate\"")
	flag.IntVar(&cf.Del, "del", -1, "Type the quest index to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Type the quest index to toggle 'completed'")
	flag.BoolVar(&cf.LeftQuestsList, "leftQuests", false, "List all the remaining quests")
	flag.BoolVar(&cf.List, "list", false, "List all quests in a pretty table")

	//extra info for -help
	flag.Usage = func() {
		flag.PrintDefaults()
		fmt.Println(`
Examples:
	quest -add "Eat chiken : 3 : 01 Dec 2024 5PM"
	quest -edit "2:-:4:-" (changes urgency at idx 2 to 4)
	quest -list
	quest -del 1
		`)
	}

	//proccess the cli arg and put the values in resp fields in cf
	//e.g. for '-add "henlo"' the value "henlo" is put into cf.Add
	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch { //to see which flag was called in the cli
	case cf.List:
		todos.print() //print our table
	case cf.LeftQuestsList:
		todos.leftQuestPrint() 
	case cf.Add != "": // -add "henlo"
		todos.add(cf.Add)
	case cf.Edit != "": // // -edit "2:-:3:-"
		err := todos.edit(cf.Edit)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)

	case cf.Del != -1: // -del 2 (-1 is invalid idx)
		todos.delete(cf.Del)

	default:
		fmt.Println("Invalid command")
	}
}
