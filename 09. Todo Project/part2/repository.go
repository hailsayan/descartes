package qtodo

import (
	"errors"
)

type Database interface {
	GetTaskList() []Task
	GetTask(string) (Task, error)
	SaveTask(Task) error
	DelTask(string) error
}

type InMemoryDatabase struct {
	Tasks map[string]Task
}

func NewDatabase() *InMemoryDatabase {
	return &InMemoryDatabase{map[string]Task{}}
}

func (in *InMemoryDatabase) GetTaskList() []Task {
	result := []Task{}
	for _, v := range in.Tasks {
		result = append(result, v)
	}
	return result
}

func (in *InMemoryDatabase) GetTask(name string) (Task, error) {
	if t, ok := in.Tasks[name]; !ok {
		return nil, errors.New("no such task")
	} else {
		return t, nil
	}
}

func (in *InMemoryDatabase) SaveTask(task Task) error {
	_, err := in.GetTask(task.GetName())
	if err == nil {
		return errors.New("duplicate")
	}

	in.Tasks[task.GetName()] = task
	return nil
}

func (in *InMemoryDatabase) DelTask(name string) error {
	_, err := in.GetTask(name)
	if err != nil {
		return err
	}

	delete(in.Tasks, name)
	return nil
}
