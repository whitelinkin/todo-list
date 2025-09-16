package tasks

import "testing"

func TestAddTask(t *testing.T) {
	StorageTask = []Task{}

	AddTask("testTitle", "testBody")

	if len(StorageTask) != 1 {
		t.Errorf("ожидал 1 задачу, получил %d", len(StorageTask))
	}

	task := StorageTask[0]
	if task.Title != "testTitle" {
		t.Errorf("ожидал Title 'testTitle', получил '%s'", task.Title)
	}
	if task.Body != "testBody" {
		t.Errorf("ожидал Body 'testBody', получил '%s'", task.Body)
	}
	if task.Done != false {
		t.Errorf("ожидал Done=false, получил %v", task.Done)
	}
	if task.CreatedAt.IsZero() {
		t.Errorf("CreatedAt не должен быть пустым")
	}
}

func TestDone(t *testing.T) {
	StorageTask = []Task{}

	AddTask("doneTask", "body")

	Done("doneTask")

	if !StorageTask[0].Done {
		t.Errorf("ожидал Done=true, получил false")
	}
	if StorageTask[0].DoneAt == nil {
		t.Errorf("ожидал, что DoneAt установлен, но он nil")
	}
}

func TestDelete(t *testing.T) {
	StorageTask = []Task{}

	AddTask("task1", "body1")
	AddTask("task2", "body2")

	Delete("task1")

	if len(StorageTask) != 1 {
		t.Errorf("ожидал 1 задачу, получил %d", len(StorageTask))
	}

	if StorageTask[0].Title != "task2" {
		t.Errorf("ожидал 'task2', получил '%s'", StorageTask[0].Title)
	}
}

func TestList(t *testing.T) {
	StorageTask = []Task{}

	AddTask("task1", "body1")
	AddTask("task2", "body2")

	if len(StorageTask) != 2 {
		t.Errorf("ожидал 2 задачи, получил %d", len(StorageTask))
	}
}
