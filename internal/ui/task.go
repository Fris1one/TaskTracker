package ui

import (
	"TaskTracker/internal/task"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
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
		row := container.NewBorder(
			nil,
			nil,
			num,
			container.NewHBox(
				delim,
				priority,
				star,
				stage,
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
		err :=
			a.AddTaskPriority(name.Text, priority.Text)
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

func (a *App) hasFavouriteOnDate(date task.TaskDeadLine) bool {
	tasks := a.schedule.Dates()[date]
	for _, t := range tasks {
		if t.IsFavourite() {
			return true
		}
	}
	return false
}
