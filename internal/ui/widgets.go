package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (a *App) buildMainContent() fyne.CanvasObject {
	left := container.NewVBox(
		a.createCalendar(),
	)
	right := container.NewVBox(
		a.createDescription(),
	)
	title := container.NewCenter(widget.NewLabel("This is your schedule!"))
	main := a.doCard(
		left,
		right,
		"",
	)
	res := container.NewVBox(title, main)
	return res
}

func (a *App) doCard(left fyne.CanvasObject, right fyne.CanvasObject, title string) fyne.CanvasObject {
	split := container.NewHSplit(left, right)
	split.Offset = 0.33
	return widget.NewCard(
		title,
		"",
		split,
	)
}
