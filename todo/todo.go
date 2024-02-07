package todo

import (
	"encoding/json"
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
	fmt.Printf("%+v\n", *t)
}

func (t *Todos) Complete(index int) error {
	ls := *t
	if index < 0 || index > len(ls) {
		return fmt.Errorf("invalid task %d", index)
	}

	ls[index-1].CompletedAt = time.Now()
	ls[index-1].Done = true

	return nil
}

func (t *Todos) Remove(index int) error {
	ls := *t
	if index < 0 || index > len(ls) {
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

	var buf []byte
	buf, err = io.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(buf, t)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", *t)

	return nil
}

func (t *Todos) Store(filename string) error {

	buf, err := json.Marshal(t)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	file.Chmod(0644)
	n, err := file.Write(buf)
	if err != nil {
		fmt.Printf("err occured while writing: %+v\n", err)
		return err
	}

	fmt.Printf("wrote %d bytes\n", n)

	return nil
}
