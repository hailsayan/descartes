package qtodo

import (
	"errors"
	"time"
)

type Task interface {
	DoAction()
	GetAlarmTime() time.Time
	GetAction() func()
	GetName() string
	GetDescription() string
}

type TaskStruct struct {
	alarmAction func()
	alarmTime   time.Time
	name        string
	description string
}

func (ta *TaskStruct) DoAction() {
	ta.alarmAction()
}

func (ta *TaskStruct) SetAlarmTime(target time.Time) error {
	if time.Now().After(target) {
		return errors.New("time cannot be in the past")
	}

	ta.alarmTime = target

	return nil
}

func (ta *TaskStruct) SetAction(target func()) error {
	ta.alarmAction = target
	return nil
}

func (ta *TaskStruct) SetName(target string) error {
	if ta.GetName() == target || len(target) == 0 {
		return errors.New("invalid or duplicate")
	}

	ta.name = target
	return nil
}

func (ta *TaskStruct) SetDescription(target string) error {
	if ta.GetDescription() == target || len(target) == 0 {
		return errors.New("invalid or duplicate")
	}

	ta.description = target
	return nil
}

func (ta *TaskStruct) GetAction() func() {
	return ta.alarmAction
}

func (ta *TaskStruct) GetAlarmTime() time.Time {
	return ta.alarmTime
}

func (ta *TaskStruct) GetName() string {
	return ta.name
}

func (ta *TaskStruct) GetDescription() string {
	return ta.description
}

func NewTask(action func(), alarmTime time.Time, name, description string) (*TaskStruct, error) {
	if len(name) == 0 || len(description) == 0 {
		return &TaskStruct{}, errors.New("cannot have empty names or description")
	}

	if time.Until(alarmTime) <= 0 {
		return &TaskStruct{}, errors.New("alarn time has to be in the future")
	}

	task := &TaskStruct{action, alarmTime, name, description}
	return task, nil
}
