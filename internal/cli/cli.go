package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/whitelinkin/internal/events"
	"github.com/whitelinkin/internal/tasks"
)

func Run() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if checkInput := scanner.Scan(); !checkInput {
			fmt.Println("Ошибка ввода!")
			return
		}

		text := scanner.Text()

		input := strings.Fields(text)

		if len(input) == 0 {
			fmt.Println("Пустая команда")
			events.AddEvent(text, "Пустая команда")
			continue
		}

		cmd := input[0]
		switch cmd {
		case "help":
			tasks.Help()
			events.AddEvent(text, "")
		case "add":
			if len(input) < 3 {
				fmt.Println("Недостаточно аргументов")
				events.AddEvent(text, "Недостаточно аргументов")
			} else {
				title := input[1]
				body := strings.Join(input[2:], " ")
				tasks.AddTask(title, body)
				events.AddEvent(text, "")
			}
		case "list":
			tasks.List()
			events.AddEvent(text, "")
		case "del":
			if len(input) < 2 {
				fmt.Println("Недостаточно аргументов")
				events.AddEvent(text, "Недостаточно аргументов")
			} else {
				title := input[1]
				tasks.Delete(title)
				events.AddEvent(text, "")
			}
		case "done":
			if len(input) < 2 {
				fmt.Println("Недостаточно аргументов")
				events.AddEvent(text, "Недостаточно аргументов")
			} else {
				title := input[1]
				tasks.Done(title)
				events.AddEvent(text, "")
			}
		case "events":
			events.ListEvents()
			events.AddEvent(text, "")
		case "exit":
			return
		default:
			fmt.Printf("Неизвестная команда: %s", input)
			events.AddEvent(text, "Неизвестная команда")
		}
	}
}
