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
	tasks    map[int]Task
}

type Task struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

func New(filepath string) *Repository {
	file, err := os.Open(filepath)
	tasks, err := getTasksFromFile(file)
	if err != nil {
		// если файл не найден — создаём пустой репозиторий
		if os.IsNotExist(err) {
			return &Repository{filepath: filepath, tasks: tasks}
		}
		fmt.Println("Ошибка открытия файла:", err)
		return &Repository{filepath: filepath, tasks: tasks}
	}
	defer file.Close()
	
	if err != nil && !errors.Is(err, io.EOF) {
		fmt.Println("Ошибка чтения файла:", err)
		return &Repository{filepath: filepath, tasks: tasks}
	}

	return &Repository{
		filepath: filepath,
		tasks:    tasks,
	}
}

func getTasksFromFile(file *os.File) (map[int]Task, error) {
	tasks := make(map[int]Task)

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); !errors.Is(err, io.EOF) {
		return tasks, err
	}
	return tasks, nil
}

func (repository *Repository) Add(name string) (int, error) {
	id := len(repository.tasks) + 1
	task := Task{
		Id:     id,
		Name:   name,
		Status: 0,
	}
	repository.tasks[id] = task

	// пересоздаём файл и пишем обновлённый JSON
	file, err := os.Create(repository.filepath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(repository.tasks); err != nil {
		return 0, err
	}

	return id, nil
}
