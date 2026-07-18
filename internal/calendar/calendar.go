package calendar

import (
	"TaskTracker/internal/task"
	"time"
)

func NewDate(day, month, year int) task.TaskDeadLine {
	var new task.TaskDeadLine
	new.SetDeadLine(day, month, year)
	return new
}

func Fill(date task.TaskDeadLine) [6][7]task.TaskDeadLine {
	var monthDays [6][7]task.TaskDeadLine
	first := time.Date(date.Year(), time.Month(date.Month()), 1, 0, 0, 0, 0, time.UTC).Weekday()
	if first == 0 {
		first = 7
	}
	daysThis := daysAmount(date)
	var monthPrev task.TaskDeadLine
	if date.Month() > 1 {
		monthPrev.SetDeadLine(1, date.Month()-1, date.Year())
	} else {
		monthPrev.SetDeadLine(1, 12, date.Year()-1)
	}
	daysPrev := daysAmount(monthPrev)
	var monthNext task.TaskDeadLine
	if date.Month() == 12 {
		monthNext.SetDeadLine(1, 1, date.Year()+1)
	} else {
		monthNext.SetDeadLine(1, date.Month()+1, date.Year())
	}
	for i := 0; i < int(first)-1; i++ {
		monthDays[0][i].SetDeadLine(daysPrev-int(first)+i+2, monthPrev.Month(), monthPrev.Year())
	}
	cnt := 1
	indj := 0
	for i := 0; i < 6; i++ {
		if i == 0 {
			indj = int(first) - 1
		} else {
			indj = 0
		}
		for j := indj; j < 7; j++ {
			monthDays[i][j].SetDeadLine(cnt, date.Month(), date.Year())
			if cnt > daysThis {
				monthDays[i][j].SetDeadLine(cnt-daysThis, monthNext.Month(), monthNext.Year())
			}
			cnt = cnt + 1
		}
	}
	return monthDays
}
func daysAmount(date task.TaskDeadLine) (amount int) {
	month := date.Month()
	year := date.Year()
	if month == -1 {
		month = 12
		year = year - 1
	}
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		amount = 31
	case 2:
		if isLeapYear(date) {
			amount = 29
		} else {
			amount = 28
		}
	default:
		amount = 30
	}
	return
}

func isLeapYear(date task.TaskDeadLine) (res bool) {
	switch {
	case date.Year()%4 == 0:
		res = true
	case date.Year()%100 == 0:
		res = false
	case date.Year()%400 == 0:
		res = true
	default:
		res = false
	}
	return
}
