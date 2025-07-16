// ui-generator/main.go
package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/google/uuid"
	files "stagfoo.pennyblack/app/files"
	"stagfoo.pennyblack/app/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var DB files.DB
var ROUTE = "list"
var selectedBook files.Book
var selectedIndex int

func updateScreen(myWindow fyne.Window) {
	// Capture the current canvas content
	id := uuid.New()
	img := myWindow.Canvas().Capture()
	// Check if capture returned nil (window not ready)
	if img == nil {
		fmt.Println("Canvas capture returned nil - window not ready")
		return
	}
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		return
	}
	file, err := os.Create(cwd + "/mock/screenshots/" + id.String() + ".png")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()
	png.Encode(file, img)
}

func LoadDB() files.DB {
	cwd, _ := os.Getwd()
	return files.ReadToml(cwd + "/mock/books/books.toml")
}

func GetListFromDB(db_books []files.Book) []string {
	var list []string
	for _, book := range db_books {
		list = append(list, book.Title)
	}
	return list
}

func main() {
	myApp := app.New()
	// Set custom font as default
	SetTheme(myApp)
	myWindow := myApp.NewWindow("E-ink UI")
	myWindow.Resize(fyne.NewSize(600, 400))
	DB = LoadDB()
	listView := ui.List(DB.Books, selectedIndex)
	// Add click handler to the list
	listView.OnSelected = func(id widget.ListItemID) {
		updateWindowContent(myWindow, listView)
	}
	// Add keyboard shortcuts
	myWindow.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
		handleKeyPress(key, myWindow, listView)
	})

	// Initial content setup
	updateWindowContent(myWindow, listView)

	myWindow.ShowAndRun()
}

func updateWindowContent(myWindow fyne.Window, listView *widget.List) {
	switch ROUTE {
	case "list":
		myWindow.SetContent(listView)
	case "book":
		var chapters, book = files.ReadEPUB(selectedBook.File)
		for _, chapter := range chapters {
			files.ReadItem(*chapter.Item)
		}
		defer book.Close()
		text := widget.NewLabel("Selected: " + selectedBook.Title)
		content, err := files.XhtmlToRichText(string(files.ReadItem(*chapters[1].Item)))
		if err != nil {
			fmt.Println("Error converting XHTML to RichText:", err)
			return
		}
		listButton := widget.NewButton("To List", func() {
			ROUTE = "list"
			updateWindowContent(myWindow, listView)
		})

		// Create the content container first
		bookView := container.NewVBox(text, content, listButton)

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
			if selectedIndex < len(DB.Books)-1 {
				selectedIndex++
				list.Select(selectedIndex)
				list.ScrollTo(selectedIndex)
			}
		case fyne.KeyReturn, fyne.KeyEnter:
			selectedBook = DB.Books[selectedIndex]
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
	updateScreen(myWindow)
}
