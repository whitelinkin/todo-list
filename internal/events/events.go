package events

import (
	"fmt"
	"time"
)

var StorageEvents []Event

type Event struct {
	Input       string
	Description string
	CreatedAt   time.Time
}

func AddEvent(input, description string) {
	newEvent := Event{
		Input:       input,
		Description: description,
		CreatedAt:   time.Now(),
	}
	StorageEvents = append(StorageEvents, newEvent)
}

func ListEvents() {
	if len(StorageEvents) == 0 {
		fmt.Println("Событий пока нет")
		return
	}

	for _, event := range StorageEvents {
		fmt.Printf("[%s] Ввод: %s | Описание: %s\n",
			event.CreatedAt.Format("2006-01-02 15:04:05"),
			event.Input,
			event.Description,
		)
	}
}
