package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"stagfoo.pennyblack/app/files"
)

func List(dbBooks []files.Book, selectedIndex int) *widget.List {
	var data []string
	for _, book := range dbBooks {
		data = append(data, book.Title)
	}
	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		},
	)
	// Set initial selection
	list.Select(selectedIndex)
	return list
}
