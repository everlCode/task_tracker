package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

type Repository struct {
	filepath string
	tasks    []Task
}

type Task struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

func New(filepath string) *Repository {
	file, err := os.Open(filepath)
	if err != nil {
		// если файл не найден — создаём пустой репозиторий
		if os.IsNotExist(err) {
			return &Repository{filepath: filepath, tasks: []Task{}}
		}
		fmt.Println("Ошибка открытия файла:", err)
		return &Repository{filepath: filepath, tasks: []Task{}}
	}
	defer file.Close()

	tasks, err := getTasksFromFile(file)
	if err != nil && !errors.Is(err, io.EOF) {
		fmt.Println("Ошибка чтения файла:", err)
		return &Repository{filepath: filepath, tasks: []Task{}}
	}

	return &Repository{
		filepath: filepath,
		tasks:    tasks,
	}
}

func getTasksFromFile(file *os.File) ([]Task, error) {
	var tasks []Task
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (repository *Repository) Add(name string) {
	id := len(repository.tasks) + 1
	task := Task{
		Id:     id,
		Name:   name,
		Status: 0,
	}
	repository.tasks = append(repository.tasks, task)

	// пересоздаём файл и пишем обновлённый JSON
	file, err := os.Create(repository.filepath)
	if err != nil {
		fmt.Println("Ошибка записи файла:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(repository.tasks); err != nil {
		fmt.Println("Ошибка кодирования JSON:", err)
	}
}
