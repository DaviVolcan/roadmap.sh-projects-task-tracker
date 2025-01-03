package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID          int       `json:"id"`          // Serializado como "id"
	Description string    `json:"description"` // Serializado como "description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func Dispatcher(commands []string, tasks *[]Task) error {
	numberOfArgs := len(commands)
	switch commands[0] {
	case "add":
		if numberOfArgs > 2 {
			return ErrTooManyArguments
		}
		if numberOfArgs == 1 {
			return ErrTooFewArguments
		}
		return addTask(commands[1], tasks)

	default:
		return ErrUnknownCommand
	}

}

func addTask(description string, tasks *[]Task) error {
	instantTIme := time.Now().UTC()
	task := Task{
		ID:          time.Now().Nanosecond(),
		Description: description,
		CreatedAt:   instantTIme,
		UpdatedAt:   instantTIme,
	}

	*tasks = append(*tasks, task)
	return nil
	// todo verify erros

}

func main() {

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Nenhum argumento foi passado.")
		return
	}

	err, tasks := loadTasks("tasks.json")
	if err != nil {
		//todo treat error
	}
	err = Dispatcher(args, &tasks)
	fmt.Println(err)

	err = saveTasks(tasks, "tasks.json")
	if err != nil {
		os.Exit(1)
	}

}

func loadTasks(filename string) (error, []Task) {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// Retorna uma lista vazia se o arquivo não existir
			return nil, []Task{}
		}
		return fmt.Errorf("erro ao ler o arquivo: %w", err), nil
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return fmt.Errorf("erro ao desserializar JSON: %w", err), nil
	}
	return nil, tasks
}

func saveTasks(tasks []Task, filename string) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("erro ao serializar as tasks: %w", err)
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("erro ao salvar o arquivo: %w", err)
	}
	return nil
}

var ErrUnknownCommand = errors.New("unknown command")
var ErrTooManyArguments = errors.New("too many arguments")
var ErrTooFewArguments = errors.New("too few arguments")
var ErrCantConvertStringToInt = errors.New("can't convert string to int")
var ErrFailedToOpenFile = errors.New("failed to open file")
