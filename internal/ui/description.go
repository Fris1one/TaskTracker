package ui

import (
	"TaskTracker/internal/task"

	"TaskTracker/internal/calendar"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (a *App) createDescription() fyne.CanvasObject {
	title := widget.NewLabel(fmt.Sprintf("%d.%s.%d", a.chosenDate.Day(), interpreateMonth(a.chosenDate.Month()), a.chosenDate.Year()))
	centeredTitle := container.NewCenter(title)
	var description []fyne.CanvasObject
	tasks := a.schedule.Dates()[a.chosenDate.TaskDeadLine]
	for i, t := range tasks {
		taskCopy := t
		num := widget.NewLabel(fmt.Sprintf("%d", i+1))
		nameData := binding.NewString()
		nameData.Set(string(t.Name()))
		name := widget.NewEntryWithData(nameData)
		name.OnSubmitted = func(text string) {
			taskCopy.SetName(text)
			a.update()
		}
		delim := widget.NewLabel("-")
		priorityData := binding.NewString()
		priorityData.Set(fmt.Sprintf("%d", t.Priority()))
		priority := widget.NewEntryWithData(priorityData)
		priority.OnSubmitted = func(text string) {
			value, err := strconv.Atoi(text)
			if err == nil {
				taskCopy.SetPriority(value)
				a.schedule.SortDay(a.chosenDate.TaskDeadLine)
			}
			a.update()
		}
		star := newStarButton(a, t)
		stage := newStageButton(a, t)
		changeDate := widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), func() {
			a.btnChangeDateAction(t)
			a.update()
		})
		row := container.NewBorder(
			nil,
			nil,
			num,
			container.NewHBox(
				delim,
				priority,
				star,
				stage,
				changeDate,
				widget.NewLabel(" "),
			),
			name,
		)
		description = append(description, row)
	}
	add := widget.NewButton("Add", func() {
		a.btnAddAction()
	})
	delete := widget.NewButton("Delete", func() {
		a.btnDeleteAction()
	})
	buttons := container.NewGridWithColumns(2, add, delete)
	grid := container.NewVBox(description...)
	scroll := container.NewVScroll(grid)
	scroll.SetMinSize(fyne.NewSize(500, 200))
	return container.NewBorder(
		centeredTitle,
		buttons,
		nil,
		nil,
		scroll,
	)
}

func (a *App) AddTaskPriority(name string, priority string) error {
	pr, err := strconv.Atoi(priority)
	if err != nil {
		return err
	}
	t, err := task.NewWithPriority(name, pr)
	if err != nil {
		return err
	}
	t.SetDeadLine(a.chosenDate.Day(), a.chosenDate.Month(), a.chosenDate.Year())
	a.schedule.AddTask(t)
	return nil
}

func (a *App) btnAddAction() {
	popUp := widget.NewPopUp(container.NewVBox(), a.window.Canvas())
	title := widget.NewLabel("Write action")
	name := widget.NewEntry()
	name.SetPlaceHolder("Write name")
	priority := widget.NewEntry()
	priority.SetPlaceHolder("Write priority from 1 to 5")
	inputs := container.NewGridWithColumns(3, name, widget.NewLabel("-"), priority)
	add := widget.NewButton("Add", func() {
		err := a.AddTaskPriority(name.Text, priority.Text)
		if err == nil {
			popUp.Hide()
			a.update()
		}
	},
	)
	popUp.Content = container.NewVBox(title, inputs, add)
	popUp.Resize(fyne.NewSize(300, 0))
	popUp.Show()
}

func (a *App) btnDeleteAction() {
	popUp := widget.NewPopUp(container.NewVBox(), a.window.Canvas())
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Write task number")
	deleteBtn := widget.NewButton("Delete", func() {
		index, err := strconv.Atoi(entry.Text)
		if err != nil {
			return
		}
		err = a.schedule.DeleteTask(a.chosenDate.TaskDeadLine, index)
		if err == nil {
			popUp.Hide()
			a.update()
		}
	},
	)
	popUp.Content = container.NewVBox(entry, deleteBtn)
	popUp.Resize(fyne.NewSize(300, 0))
	popUp.Show()
}

func (a *App) btnChangeDateAction(t *task.Task) {
	a.changeDay(t)
	a.update()
}

func (a *App) changeDay(t *task.Task) {
	popUp := widget.NewPopUp(container.NewVBox(), a.window.Canvas())
	month := container.NewCenter(widget.NewLabel(fmt.Sprintf("%s.%d", interpreateMonth(a.chosenDate.Month()), a.chosenDate.Year())))
	weekdays := []string{"пн", "вт", "ср", "чт", "пт", "сб", "вс"}
	var days []fyne.CanvasObject
	for _, day := range weekdays {
		label := widget.NewLabelWithStyle(day, fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
		days = append(days, label)
	}
	dates := calendar.Fill(a.chosenDate.TaskDeadLine)
	for i := range dates {
		for j := range dates[i] {
			current := dates[i][j]
			btn := widget.NewButton(fmt.Sprintf("%d", current.Day()), func() {})
			if current.Month() != a.chosenDate.Month() {
				btn.Importance = widget.LowImportance
			}
			btn.OnTapped = func() {
				a.schedule.MoveTask(t, current)
				popUp.Hide()
				a.update()
			}
			if a.hasFavouriteOnDate(current) {
				btn.Icon = theme.InfoIcon()
			}
			days = append(days, btn)
		}
	}
	gridMonths := container.NewGridWithColumns(7, days...)
	popUp.Content = container.NewBorder(
		month,
		nil,
		nil,
		nil,
		gridMonths,
	)
	popUp.Resize(fyne.NewSize(200, 200))
	popUp.Show()
}
func (a *App) hasFavouriteOnDate(date task.TaskDeadLine) bool {
	tasks := a.schedule.Dates()[date]
	for _, t := range tasks {
		if t.IsFavourite() {
			return true
		}
	}
	return false
}
