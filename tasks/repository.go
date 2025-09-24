package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

const (
	StatusNew int = iota
	StatusInProgress
	StatusDone
	StatusArchived
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

func (repository *Repository) Get(id int) (*Task, error) {
	if task, ok := repository.tasks[id]; ok {
		return &task, nil
	}

	return nil, errors.New("Задачи не существует!")
}

func (repository *Repository) Add(name string) (int, error) {
	var max int = 0
	for _, v := range repository.tasks {
		if v.Id > max {
			max = v.Id
		}
	}
	id := max + 1
	task := Task{
		Id:     id,
		Name:   name,
		Status: 0,
	}
	repository.tasks[id] = task

	if err := repository.save(); err != nil {
		return 0, err
	}

	return id, nil
}

func (repository *Repository) Update(id int, name string) (int, error) {
	task, err := repository.Get(id)
	if err != nil {
		return 0, err
	}

	repository.tasks[id] = *task

	if err := repository.save(); err != nil {
		return 0, err
	}

	return id, nil
}

func (repository *Repository) UpdateStatus(id int, status int) (int, error) {
	task, err := repository.Get(id)
	if err != nil {
		return 0, err
	}
	task.Status = status

	repository.tasks[id] = *task

	if err := repository.save(); err != nil {
		return 0, err
	}

	return id, nil
}

func (repository *Repository) Delete(id int) (int, error) {
	_, err := repository.Get(id)
	if err != nil {
		return 0, err
	}

	delete(repository.tasks, id)

	if err := repository.save(); err != nil {
		return 0, err
	}

	return id, nil
}

func (repository *Repository) save() error {
	// пересоздаём файл и пишем обновлённый JSON
	file, err := os.Create(repository.filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(repository.tasks); err != nil {
		return err
	}

	return nil
}
