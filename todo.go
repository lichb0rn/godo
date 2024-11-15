package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (t *Todos) Add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}
	*t = append(*t, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("index out of range")
		fmt.Println(err)
		return err
	}

	return nil
}

func (todos *Todos) Delete(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	*todos = append(t[0:index], t[index+1:]...)
	return nil
}

func (todos *Todos) Toggle(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	isCompleted := t[index].Completed
	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}
	t[index].Completed = !isCompleted
	return nil
}

func (todos *Todos) Edit(index int, title string) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title
	return nil
}

func (todos *Todos) Print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for index, t := range *todos {
		completed := "❌"
		completedAt := ""

		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(
			strconv.Itoa(index),
			t.Title,
			completed,
			t.CreatedAt.Format(time.RFC1123),
			completedAt,
		)
	}
	table.Render()
}
