package tasks

import (
	"fmt"
	"time"

	"github.com/k0kubun/pp"
)

var StorageTask []Task

type Task struct {
	Title     string
	Body      string
	Done      bool
	CreatedAt time.Time
	DoneAt    *time.Time
}

func Help() {
	pp.Println(`
	help 
	команда позволяет узнать доступные команды и их формат

    add {заголовок задачи из одного слова} {текст задачи из одного или нескольких слов}
	команда позволяет добавлять новые задачи в список задач

    list
	команда позволяет получить полный список всех задач

    del {заголовок существующей задачи}
	команда позволяет удалить задачу по её заголовку

    done {заголовок существующей задачи}
	команда позволяет отменить задачу как выполненную

    events
	команда позволяет получить список всех событий

    exit
	команда позволяет завершить выполнение программы
`)
}

func AddTask(title, body string) {
	newTask := Task{
		Title:     title,
		Body:      body,
		Done:      false,
		CreatedAt: time.Now(),
	}

	StorageTask = append(StorageTask, newTask)
}

func List() {
	for _, t := range StorageTask {
		status := "[ ]"
		if t.Done {
			status = "[✓]"
		}

		fmt.Printf("%s Задача: %s | Описание: %s | Создана: %s\n",
			status, t.Title, t.Body, t.CreatedAt.Format("2006-01-02 15:04"))

		if t.DoneAt != nil {
			fmt.Printf("	Выполнено: %s\n", t.DoneAt.Format("2006-01-02 15:04"))
		}
	}
}

func Done(title string) {
	found := false
	for i, _ := range StorageTask {
		if title == StorageTask[i].Title {

			StorageTask[i].Done = true

			now := time.Now()
			StorageTask[i].DoneAt = &now

			found = true

			break
		}
	}

	if found == false {
		fmt.Println("Задача не найдена")
	}
}

func Delete(title string) {
	found := false
	for i, _ := range StorageTask {
		if title == StorageTask[i].Title {
			StorageTask = append(StorageTask[:i], StorageTask[i+1:]...)

			found = true
			break
		}
	}

	if !found {
		fmt.Println("Задача не найдена")
	}
}
