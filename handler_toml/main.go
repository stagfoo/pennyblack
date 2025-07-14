package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/BurntSushi/toml"
)

type Activity struct {
	Date    string `toml:"date"`
	Page    int    `toml:"page"`
	Minutes int    `toml:"minutes"`
}

type Book struct {
	Title string `toml:"title"`
	File  string `toml:"file"`
	// Cover    string   `toml:"cover"`
	// Author   string   `toml:"author"`
	Activity Activity `toml:"activity"`
}

type CacheToml struct {
	CreatedAt string `toml:"created_at"`
	UpdatedAt string `toml:"updated_at"`
	Version   int    `toml:"version"`
	Books     []Book `toml:"books"`
}

func ReadToml(path string) CacheToml {
	fmt.Print("Target Path", path)
	doc, readErr := os.ReadFile(path)
	if readErr != nil {
		fmt.Print("File Read Error", readErr)
		var empty CacheToml
		return empty
	}
	var db CacheToml
	err := toml.Unmarshal([]byte(doc), &db)
	if err != nil {
		fmt.Print("Toml Read Error", err)
		var empty CacheToml
		return empty
	}
	return db
}

func SaveToml(db CacheToml, path string) bool {
	b, err := toml.Marshal(db)
	fmt.Print("Got DB")

	if err != nil {
		fmt.Print("Got Error")

		panic(err)
	}
	err = os.WriteFile(path, b, 0644)
	if err != nil {
		fmt.Print("File Write Error", err)
		return false
	}
	fmt.Print("Got Done")

	return true
}

func ListDirectoryContents(currentDir string, shouldPrint bool) ([]string, error) {
	files, err := filepath.Glob(filepath.Join(currentDir, "*"))
	if err != nil {
		fmt.Printf("\n------------\n")
		fmt.Printf("ListDirectoryContents | Filepath.Glob Error | %v\n", err)
		return nil, fmt.Errorf("error using filepath.Glob: %v", err)
	}

	if shouldPrint {
		for _, file := range files {
			fmt.Println(file)
		}
		return nil, nil // Returning nil slice and nil error when printing
	} else {
		return files, nil
	}
}

func ConvertFilePathToBook(path string) Book {
	var u Book
	var a Activity
	u.Title = "Book"
	u.File = path
	u.Activity = a
	return u
}

func main() {

	var currentDir, err = os.Getwd()
	if err != nil {
		fmt.Print("failed to get " + currentDir)
		os.Exit(1)
	}
	fmt.Print("Looking for files")
	var files, fileErr = ListDirectoryContents(currentDir+"/mock/books", false)
	fmt.Print(fileErr)
	fmt.Print(files)

	if err == nil {
		var bookList []Book
		for _, file := range files {
			bookList = append(bookList, ConvertFilePathToBook(file))
		}
		var db = *&CacheToml{
			Books: bookList,
		}
		SaveToml(db, currentDir+"/mock/bin/books.toml")
	}

	go func() {
		ticker := time.NewTicker(2 * time.Second)
		for range ticker.C {

		}
	}()
}
