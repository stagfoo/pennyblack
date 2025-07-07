// ui-generator/main.go
package main

import (
	"time"

	"stagfoo.pennyblack/cmd"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var data = []string{"Book 1", "Book 2", "Book 3", "Book 4", "Book 5", "Book 6", "Book 7", "Book 8", "Book 9", "Book 10", "Book 11", "Book 12", "Book 13", "Book 14", "Book 15", "Book 16", "Book 17", "Book 18", "Book 19", "Book 20"}

var ROUTE = "list"
var selectedBook string

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("E-ink UI")
	myWindow.Resize(fyne.NewSize(600, 400))

	// View "List"
	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})

	// Add click handler to the list
	list.OnSelected = func(id widget.ListItemID) {
		selectedBook = data[id]
		ROUTE = "book"
		// Update the window content to show the book view
		updateWindowContent(myWindow, list)
	}

	// View "Book"

	// Initial content setup
	updateWindowContent(myWindow, list)

	// Screenshot capture routine
	go func() {
		ticker := time.NewTicker(2 * time.Second)
		for range ticker.C {
			cmd.ScreenRefresh()
		}
	}()

	myWindow.ShowAndRun()
}

func updateWindowContent(myWindow fyne.Window, list *widget.List) {
	switch ROUTE {
	case "list":
		myWindow.SetContent(list)
	case "book":
		text := widget.NewLabel("Selected: " + selectedBook)
		listButton := widget.NewButton("To List", func() {
			ROUTE = "list"
			updateWindowContent(myWindow, list)
		})
		bookView := container.NewVBox(text, listButton)
		myWindow.SetContent(bookView)
	}
}
