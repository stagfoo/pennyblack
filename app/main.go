// ui-generator/main.go
package main

import (
	"image/color"
	"time"

	"stagfoo.pennyblack/cmd"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type customTheme struct {
	font fyne.Resource
}

func (t *customTheme) Font(style fyne.TextStyle) fyne.Resource {
	return t.font
}

func (t *customTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, variant)
}

func (t *customTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t *customTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

var ppmondwestRegular []byte

var data = []string{"Book 1", "Book 2", "Book 3", "Book 4", "Book 5", "Book 6", "Book 7", "Book 8", "Book 9", "Book 10", "Book 11", "Book 12", "Book 13", "Book 14", "Book 15", "Book 16", "Book 17", "Book 18", "Book 19", "Book 20"}

var ROUTE = "list"
var selectedBook string
var selectedIndex int

func main() {
	myApp := app.New()
	// Set custom font as default
	fontResource := fyne.NewStaticResource("ppmondwest-regular.ttf", ppmondwestRegular)
	myApp.Settings().SetTheme(&customTheme{
		font: fontResource,
	})
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
		},
	)
	// Set initial selection
	list.Select(selectedIndex)

	// Add click handler to the list
	list.OnSelected = func(id widget.ListItemID) {
		updateWindowContent(myWindow, list)
	}
	// Add keyboard shortcuts
	myWindow.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
		handleKeyPress(key, myWindow, list)
	})

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

func handleKeyPress(key *fyne.KeyEvent, myWindow fyne.Window, list *widget.List) {
	switch ROUTE {
	case "list":
		switch key.Name {
		case fyne.KeyUp:
			if selectedIndex > 0 {
				selectedIndex--
				list.Select(selectedIndex)
				list.ScrollTo(selectedIndex)
			}
		case fyne.KeyDown:
			if selectedIndex < len(data)-1 {
				selectedIndex++
				list.Select(selectedIndex)
				list.ScrollTo(selectedIndex)
			}
		case fyne.KeyReturn, fyne.KeyEnter:
			selectedBook = data[selectedIndex]
			ROUTE = "book"
			updateWindowContent(myWindow, list)
		}
	case "book":
		switch key.Name {
		case fyne.KeyEscape, fyne.KeyBackspace:
			ROUTE = "list"
			updateWindowContent(myWindow, list)
		}
	}
}
