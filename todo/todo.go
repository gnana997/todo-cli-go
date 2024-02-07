package todo

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []item

func (t *Todos) Add(task string) {
	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *Todos) Complete(index int) error {
	ls := *t
	if index < 0 || index >= len(ls) {
		return fmt.Errorf("invalid task %d", index)
	}

	ls[index-1].CompletedAt = time.Now()
	ls[index-1].Done = true

	return nil
}

func (t *Todos) Remove(index int) error {
	ls := *t
	if index < 0 || index >= len(ls) {
		return fmt.Errorf("invalid task %d", index)
	}

	*t = append(ls[:index-1], ls[index:]...)

	return nil
}

func (t *Todos) List() {
	for i, todo := range *t {
		fmt.Printf("%d. %s\n", i+1, todo.Task)
	}
}

func (t *Todos) Load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, file)
	if err != nil {
		return err
	}

	err = gob.NewDecoder(buf).Decode(t)
	if err != nil {
		return err
	}

	return nil
}
