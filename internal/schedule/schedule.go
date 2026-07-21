package schedule

import (
	"TaskTracker/internal/task"
	"fmt"
	"sort"
)

type Dates []task.TaskDeadLine
type Tasks map[task.TaskNumber]*task.Task
type PriorityList map[task.TaskPriority][]*task.Task
type DayTaskList map[task.TaskDeadLine][]*task.Task

type Schedule struct {
	dates map[task.TaskDeadLine][]*task.Task
}

func New() *Schedule {
	return &Schedule{
		dates: make(map[task.TaskDeadLine][]*task.Task),
	}
}

func (s *Schedule) Dates() map[task.TaskDeadLine][]*task.Task {
	return s.dates
}

func (s *Schedule) AddTask(t *task.Task) {
	s.SortPriority(t)
}

func (s *Schedule) SortPriority(t *task.Task) {
	key := t.DeadLine()
	s.dates[key] = append(s.dates[key], t)
	s.SortDay(key)
}
func (s *Schedule) SortDay(td task.TaskDeadLine) {
	sort.Slice(s.dates[td], func(i, j int) bool {
		return s.dates[td][i].Priority() > s.dates[td][j].Priority()
	})
}
func (s *Schedule) DeleteTask(t task.TaskDeadLine, i int) error {
	if i <= 0 || i > len(s.dates[t]) {
		return fmt.Errorf("incorrect number")
	}
	s.dates[t] = append(s.dates[t][:i-1], s.dates[t][i:]...)
	return nil
}

func (s *Schedule) MoveTask(t *task.Task, newDate task.TaskDeadLine) {
	oldDate := t.DeadLine()
	tasks := s.dates[oldDate]
	for i, tt := range tasks {
		if tt == t {
			s.dates[oldDate] = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	t.SetDeadLine(newDate.Day(), newDate.Month(), newDate.Year())
	s.dates[newDate] = append(s.dates[newDate], t)
	s.SortDay(newDate)
}
