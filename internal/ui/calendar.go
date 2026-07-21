package ui

import (
	"TaskTracker/internal/calendar"
	"TaskTracker/internal/task"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (a *App) createCalendar() fyne.CanvasObject {
	return container.NewBorder(
		a.doTitle(),
		nil,
		nil,
		nil,
		a.doGrid(a.doDates(calendar.Fill(a.chosenDate.TaskDeadLine))))
}

func (a *App) doTitle() fyne.CanvasObject {
	title := widget.NewLabel(fmt.Sprintf("%s.%d", interpreateMonth(a.chosenDate.Month()), a.chosenDate.Year()))
	btn := widget.NewButton(title.Text, func() {
		a.PrintMonths()
	},
	)
	return btn
}

func (a *App) doGrid(buttons []fyne.CanvasObject) fyne.CanvasObject {
	return container.NewGridWithColumns(7, buttons...)
}

func (a *App) doDates(dates [6][7]task.TaskDeadLine) (res []fyne.CanvasObject) {
	weekdays := []string{"пн", "вт", "ср", "чт", "пт", "сб", "вс"}
	for _, day := range weekdays {
		label := widget.NewLabelWithStyle(day, fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
		res = append(res, label)
	}
	for i := range dates {
		for j := range dates[i] {
			current := dates[i][j]
			btn := widget.NewButton(fmt.Sprintf("%d", current.Day()), func() {})
			if current.Month() != a.chosenDate.Month() {
				btn.Importance = widget.LowImportance
			}
			btn.OnTapped = func() {
				a.chosenDate.SetDeadLine(
					current.Day(),
					current.Month(),
					current.Year(),
				)
				a.update()
			}
			if a.hasFavouriteOnDate(current) {
				btn.Icon = theme.InfoIcon()
			}
			res = append(res, btn)
		}
	}
	return
}
func (a *App) PrintMonths() {
	popUp := widget.NewPopUp(container.NewVBox(), a.window.Canvas())
	title := widget.NewButton(fmt.Sprintf("%d", a.chosenDate.Year()), func() {
		a.PrintYears()
	})
	var buttons []fyne.CanvasObject
	for month := 1; month <= 12; month++ {
		m := month
		btn := widget.NewButton(fmt.Sprintf("%d", m), func() {
			a.chosenDate.SetMonth(m)
			popUp.Hide()
			a.update()
		},
		)
		buttons = append(buttons, btn)
	}
	grid := container.NewGridWithColumns(4, buttons...)
	popUp.Content = container.NewBorder(
		title,
		nil,
		nil,
		nil,
		grid,
	)
	popUp.Resize(fyne.NewSize(200, 200))
	popUp.Show()
}

func (a *App) PrintYears() {
	popUp := widget.NewPopUp(container.NewVBox(), a.window.Canvas())
	years := []int{2025, 2026, 2027, 2028, 2029, 2030, 2031, 2032, 2033, 2034, 2035, 2036}
	var buttons []fyne.CanvasObject
	for _, year := range years {
		y := year
		btn := widget.NewButton(fmt.Sprintf("%d", y), func() {
			a.chosenDate.SetYear(y)
			popUp.Hide()
			a.update()
		})
		buttons = append(buttons, btn)
	}
	grid := container.NewGridWithColumns(4, buttons...)
	popUp.Content = grid
	popUp.Resize(fyne.NewSize(200, 200))
	popUp.Show()
}

func interpreateMonth(month int) string {
	switch month {
	case 1:
		return "01"
	case 2:
		return "02"
	case 3:
		return "03"
	case 4:
		return "04"
	case 5:
		return "05"
	case 6:
		return "06"
	case 7:
		return "07"
	case 8:
		return "08"
	case 9:
		return "09"
	case 10:
		return "10"
	case 11:
		return "11"
	case 12:
		return "12"
	default:
		return "00"
	}
}
