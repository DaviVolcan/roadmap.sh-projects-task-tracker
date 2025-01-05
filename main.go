package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
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
		addTask(commands[1], tasks)
		return nil
	case "list":
		if numberOfArgs > 1 {
			return listTasksByStatus(tasks, commands[1])
		}
		return listAllTasks(tasks)
	case "mark-in-progress":
		if numberOfArgs > 2 {
			return ErrTooManyArguments
		}
		if numberOfArgs == 1 {
			return ErrTooFewArguments
		}
		return markTaskAs(tasks, commands[1], "in-progress")
	case "mark-done":
		if numberOfArgs > 2 {
			return ErrTooManyArguments
		}
		if numberOfArgs == 1 {
			return ErrTooFewArguments
		}
		return markTaskAs(tasks, commands[1], "done")
	case "update":
		if numberOfArgs > 3 {
			return ErrTooManyArguments
		}
		if numberOfArgs <= 2 {
			return ErrTooFewArguments
		}
		return UpdateTaskDescription(tasks, commands[1], commands[2])
	case "delete":
		if numberOfArgs > 2 {
			return ErrTooManyArguments
		}
		if numberOfArgs == 1 {
			return ErrTooFewArguments
		}
		return DeleteTask(tasks, commands[1])
	default:
		return ErrUnknownCommand
	}
}

func addTask(description string, tasks *[]Task) {
	instantTIme := time.Now().UTC()
	task := Task{
		ID:          time.Now().Nanosecond(),
		Status:      "todo",
		Description: description,
		CreatedAt:   instantTIme,
		UpdatedAt:   instantTIme,
	}

	fmt.Printf("Task added successfully (ID:%d)", task.ID)
	*tasks = append(*tasks, task)
}

func listAllTasks(tasks *[]Task) error {
	if len(*tasks) == 0 {
		return ErrNoTasksFound
	}
	for _, task := range *tasks {
		printTask(task)
	}
	return nil
}

func listTasksByStatus(tasks *[]Task, status string) error {
	foundTask := false
	if len(*tasks) == 0 {
		return ErrNoTasksFound
	}
	for _, task := range *tasks {
		if task.Status == status {
			printTask(task)
			foundTask = true
		}
	}
	if !foundTask {
		return ErrNoTasksFound
	}
	return nil
}

func markTaskAs(tasks *[]Task, idParameter string, status string) error {
	foundTask := false
	if len(*tasks) == 0 {
		return ErrNoTasksFound
	}
	num, err := strconv.Atoi(idParameter)
	if err != nil {
		return ErrCantConvertStringToInt
	}
	for index, task := range *tasks {
		if task.ID == num {
			(*tasks)[index].Status = status
			(*tasks)[index].UpdatedAt = time.Now().UTC()
			foundTask = true
			break
		}
	}
	if !foundTask {
		return ErrNoTasksFound
	}
	return nil
}

func UpdateTaskDescription(tasks *[]Task, idParameter string, description string) error {
	foundTask := false
	if len(*tasks) == 0 {
		return ErrNoTasksFound
	}
	num, err := strconv.Atoi(idParameter)
	if err != nil {
		return ErrCantConvertStringToInt
	}
	for index, task := range *tasks {
		if task.ID == num {
			(*tasks)[index].Description = description
			(*tasks)[index].UpdatedAt = time.Now().UTC()
			foundTask = true
			break
		}
	}
	if !foundTask {
		return ErrNoTasksFound
	}
	return nil
}

func DeleteTask(tasks *[]Task, idParameter string) error {
	foundTask := false
	if len(*tasks) == 0 {
		return ErrNoTasksFound
	}
	num, err := strconv.Atoi(idParameter)
	if err != nil {
		return ErrCantConvertStringToInt
	}
	TaskIndex := 0
	for index, task := range *tasks {
		if task.ID == num {
			TaskIndex = index
			foundTask = true
			break
		}
	}
	if !foundTask {
		return ErrNoTasksFound
	}
	DeleteTaskByIndex(tasks, TaskIndex)
	return nil
}

func DeleteTaskByIndex(tasks *[]Task, index int) {
	*tasks = append((*tasks)[:index], (*tasks)[index+1:]...)
	fmt.Println(*tasks)
}

func printTask(task Task) {
	fmt.Printf("%+v\n", task)
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Nenhum argumento foi passado.")
		return
	}
	err, tasks := loadTasks("tasks.json")
	if err != nil {
		panic(err)
	}

	err = Dispatcher(args, &tasks)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = saveTasks(tasks, "tasks.json")
	if err != nil {
		panic(err)
	}

	fmt.Println("sucess!")

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
var ErrNoTasksFound = errors.New("no tasks found")
