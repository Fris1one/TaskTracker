package task

import (
	"fmt"
	"time"
)

var cnt = 0

const (
	LowestPriority = iota + 1
	LowPriority
	MediumPriority
	HighPriority
	HighestPriority
)
const (
	Start    = "Started"
	Progress = "In progress"
	Finish   = "Completed"
	Delete   = "Deleted"
)

type TaskName string
type TaskPriority int
type TaskDeadLine struct {
	day   int
	month int
	year  int
}

func NewDeadLine() *TaskDeadLine {
	return &TaskDeadLine{
		day:   time.Now().Day(),
		month: int(time.Now().Month()),
		year:  time.Now().Year(),
	}
}
func (td TaskDeadLine) Day() int {
	return td.day
}
func (td TaskDeadLine) Month() int {
	return td.month
}
func (td TaskDeadLine) Year() int {
	return td.year
}
func (td *TaskDeadLine) SetDeadLine(day, month, year int) {
	td.day = day
	td.month = month
	td.year = year
}
func (td *TaskDeadLine) SetDay(day int) {
	td.day = day
}
func (td *TaskDeadLine) SetMonth(month int) {
	td.month = month
}
func (td *TaskDeadLine) SetYear(year int) {
	td.year = year
}

type TaskPossibility bool
type TaskNumber uint

type Task struct {
	name      TaskName
	priority  TaskPriority
	deadline  TaskDeadLine
	favourite bool
	stage     int
	id        int
}
type dailyTaskTime time.Time
type dailyTaskDuration time.Duration

type DailyTask struct {
	task Task
	time dailyTaskTime
}

func New(name string) *Task {
	var t = TaskDeadLine{
		day:   15,
		month: 6,
		year:  2036,
	}
	cnt++
	return &Task{
		name:      TaskName(name),
		priority:  LowestPriority,
		deadline:  t,
		stage:     0,
		favourite: false,
		id:        cnt,
	}
}
func NewWithPriority(name string, priority int) (*Task, error) {
	t := New(name)
	err := t.SetPriority(priority)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (t *Task) Name() TaskName {
	return t.name
}
func (t *Task) Priority() TaskPriority {
	return t.priority
}
func (t *Task) DeadLine() TaskDeadLine {
	return t.deadline
}

func (t *Task) Stage() int {
	return t.stage
}
func (t *Task) IsFavourite() bool {
	return t.favourite
}
func (t *Task) ID() int {
	return t.id
}
func (t *Task) SetName(newName string) {
	t.name = TaskName(newName)
}
func isValidPriority(pr int) bool {
	switch pr {
	case LowestPriority, LowPriority, MediumPriority, HighPriority, HighestPriority:
		return true
	default:
		return false
	}
}
func (t *Task) SetPriority(pr int) error {
	if isValidPriority(pr) {
		t.priority = TaskPriority(pr)
		return nil
	}
	return fmt.Errorf(`Do correct level of priority`)
}
func (t *Task) SetDeadLine(day, month, year int) {
	t.deadline.SetDeadLine(day, month, year)
}

func (t *Task) SetStage(stage int) error {
	switch stage {
	case 0, 1, 2:
		t.stage = stage
	default:
		return fmt.Errorf(`Do correct stage for this task`)
	}
	return nil
}

func (t *Task) SetFavourite(is bool) {
	t.favourite = is
}
