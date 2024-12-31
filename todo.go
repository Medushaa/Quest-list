package main

import (
	"errors"
	"fmt"
	"time"

	//for the table:
	"os"
	"strconv"
	"strings"

	"github.com/aquasecurity/table"
	"github.com/liamg/tml"
)

// a single task struct
type Todo struct {
	Title       string
	Completed   bool
	Urgency     int
	DueDate     string
	CompletedAt *time.Time
}

// 'type' create abstruction for the []Todo type. So, we can later declare
// variable of type Todos. (in main.go)
type Todos []Todo //slice/list to hold all the "Todo"s

// method to add new todo to the slice Todos
func (todos *Todos) add(details string) {
	parts := strings.SplitN(details, ":", 3)

	if len(parts) != 3 { // Ensure the format is correct
		fmt.Println("Error: Invalid format for add. Use 'title : urgency : due date'")
		os.Exit(1)
	}

	title := strings.TrimSpace(parts[0])
	urgency, err := strconv.Atoi(strings.TrimSpace(parts[1])) // Parse urgency as integer
	if err != nil || urgency < 0 || urgency > 5 {
		fmt.Println("Error: Invalid urgency level. It must be a number from 0 to 5.")
		os.Exit(1)
	}

	dueDate := strings.TrimSpace(parts[2]) // Parse due date as string (validation optional)

	// Create the new todo
	todo := Todo{
		Title:       title,
		Completed:   false,
		Urgency:     urgency,
		DueDate:     dueDate,
		CompletedAt: nil,
	}

	*todos = append(*todos, todo)
	fmt.Println("New quest sucessfully added!👍")
}

// delete a todo at that index
func (todos *Todos) delete(index int) error {
	t := *todos //slice data
	if err := t.validateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)
	fmt.Printf("Quest at index %v successfully deleted.", index)
	return nil
}

// if user inputs a invalid index (error handling)
func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err.Error())
		return err
	}
	return nil
}

// toggle task completion T/F
func (todos *Todos) toggle(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	isCompleted := t[index].Completed

	if !isCompleted { //if false
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}

	t[index].Completed = !isCompleted //toggle T<->F

	fmt.Printf("Quest at index %v marked as completed ✅. Good Job!", index)

	return nil
}

// change title, urgency and duedate at a task index
func (todos *Todos) edit(details string) error {
	//split input into parts: index, title, urgency, and due date
	parts := strings.SplitN(details, ":", 4)
	if len(parts) != 4 {
		return fmt.Errorf("invalid format for edit. Use 'index:title:urgency:due date'")
	}

	index, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return fmt.Errorf("invalid index: %v", err)
	}

	if err := (*todos).validateIndex(index); err != nil {
		return err
	}

	//get the todo to edit. address of the todo at index in the slice (pointer)
	t := &(*todos)[index]

	//Update fields if they are not placeholders
	if parts[1] != "-" { // Title
		t.Title = strings.TrimSpace(parts[1])
	}
	if parts[2] != "-" { // Urgency
		urgency, err := strconv.Atoi(strings.TrimSpace(parts[2]))
		if err != nil {
			return fmt.Errorf("invalid urgency: %v", err)
		}
		t.Urgency = urgency
	}
	if parts[3] != "-" { // Due Date
		t.DueDate = strings.TrimSpace(parts[3])
	}

	fmt.Printf("Quest at index %v has been edited.", index)
	return nil
}

// print the pretty table in the cli
func (todos *Todos) print() {
	tabl := table.New(os.Stdout)
	tabl.SetRowLines(false)
	tabl.SetHeaderStyle(table.StyleBold)
	tabl.SetLineStyle(table.StyleBrightCyan)

	tabl.SetHeaders("#", "Quest Title", "Completed", "Urgency", "Due Date", "Completed At")

	timeFormat := "Mon, 02 Jan, 03:04 PM"

	getUrgencyEmoji := func(urgency int) string {
		switch urgency {
		case 0:
			return "⚪"
		case 1:
			return "🟢"
		case 2:
			return "🟡🟡"
		case 3:
			return "🟠🟠🟠"
		case 4:
			return "🔴🔴🔴🔴"
		case 5:
			return "📢❗⏰🚨🆘"
		default:
			return "❓" // Unknown urgency
		}
	}

	for index, t := range *todos {
		completed := " ❌ "
		completedAt := ""
		title := ""
		urgency := getUrgencyEmoji(t.Urgency)

		switch t.Urgency {
		case 1:
			title = tml.Sprintf("<green>" + t.Title + "</green>")
		case 2:
			title = tml.Sprintf("<yellow>" + t.Title + "</yellow>")
		case 3:
			title = tml.Sprintf("<magenta>" + t.Title + "</magenta>")
		case 4:
			title = tml.Sprintf("<red>" + t.Title + "</red>")
		case 5:
			title = tml.Sprintf("<red>" + t.Title + "</red>")
		default:
			title = tml.Sprintf("<white>" + t.Title + "</white>")
		}

		if t.Completed {
			completed = " ✅ "
			title = tml.Sprintf("<black>" + t.Title + "</black>")
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(timeFormat) //time.RFC1123 (that was the default)
			}
			urgency = ""
		}

		tabl.AddRow(
			strconv.Itoa(index),
			title,
			completed,
			urgency, //use color here next
			t.DueDate,
			completedAt,
		)
	}

	tabl.Render()
}

// print the pretty table in the cli
func (todos *Todos) leftQuestPrint() {
	tabl := table.New(os.Stdout)
	tabl.SetRowLines(false)
	tabl.SetHeaderStyle(table.StyleBold)
	tabl.SetLineStyle(table.StyleBrightCyan)

	tabl.SetHeaders("#", "Quest Title", "Completed", "Urgency", "Due Date")

	getUrgencyEmoji := func(urgency int) string {
		switch urgency {
		case 0:
			return "⚪"
		case 1:
			return "🟢"
		case 2:
			return "🟡🟡"
		case 3:
			return "🟠🟠🟠"
		case 4:
			return "🔴🔴🔴🔴"
		case 5:
			return "📢❗⏰🚨🆘"
		default:
			return "❓" // Unknown urgency
		}
	}

	for index, t := range *todos {
		if t.Completed {
			continue
		}

		completed := " ❌ "
		title := ""
		urgency := getUrgencyEmoji(t.Urgency)

		switch t.Urgency {
		case 1:
			title = tml.Sprintf("<green>" + t.Title + "</green>")
		case 2:
			title = tml.Sprintf("<yellow>" + t.Title + "</yellow>")
		case 3:
			title = tml.Sprintf("<magenta>" + t.Title + "</magenta>")
		case 4:
			title = tml.Sprintf("<red>" + t.Title + "</red>")
		case 5:
			title = tml.Sprintf("<red>" + t.Title + "</red>")
		default:
			title = tml.Sprintf("<white>" + t.Title + "</white>")
		}

		tabl.AddRow(
			strconv.Itoa(index),
			title,
			completed,
			urgency, //use color here next
			t.DueDate,
		)
	}

	tabl.Render()
}
