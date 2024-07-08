package qtodo

import (
	"fmt"
	"time"
)

type App interface {
	StartTask(string) error
	StopTask(string)
	AddTask(string, string, time.Time, func(), bool) error
	DelTask(string) error
	GetTaskList() []Task
	GetActiveTaskList() []Task
	GetTask(string) (Task, error)
}

type Application struct {
	activeTasks map[string]struct{}
	db          Database
}

func (ap *Application) StartTask(name string) error {
	task, err := ap.GetTask(name)
	if err != nil {
		return err
	}

	ap.activeTasks[name] = struct{}{}

	go func() {
		time.Sleep(time.Until(task.GetAlarmTime()))
		if _, ok := ap.activeTasks[name]; ok {
			task.DoAction()
		}
	}()
	return nil
}

func (ap *Application) GetActiveTaskList() []Task {
	result := []Task{}
	allTasks := ap.GetTaskList()
	for _, task := range allTasks {
		if _, ok := ap.activeTasks[task.GetName()]; ok {
			result = append(result, task)
		}
	}
	return result
}

func (ap *Application) StopTask(name string) {
	delete(ap.activeTasks, name)
}

func (ap *Application) AddTask(name string, desc string, ti time.Time, action func(), deleteAfter bool) error {
	task, err := NewTask(action, ti, name, desc)
	if err != nil {
		return err
	}

	curaction := task.GetAction()
	if deleteAfter {
		err = task.SetAction(func() {
			curaction()
			err2 := ap.db.DelTask(task.GetName())
			if err2 != nil {
				fmt.Println(err2)
			}
		})
		if err != nil {
			return err
		}
	}

	return ap.db.SaveTask(task)
}

func (ap *Application) DelTask(name string) error {
	err := ap.db.DelTask(name)
	if err != nil {
		return err
	}

	ap.StopTask(name)
	return nil
}

func (ap *Application) GetTaskList() []Task {
	return ap.db.GetTaskList()
}

func (ap *Application) GetTask(name string) (Task, error) {
	return ap.db.GetTask(name)
}

func NewApp(db Database) *Application {
	return &Application{make(map[string]struct{}), db}
}
