package files

import (
	"fmt"
	"os"

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

type DB struct {
	CreatedAt string `toml:"created_at"`
	UpdatedAt string `toml:"updated_at"`
	Version   int    `toml:"version"`
	Books     []Book `toml:"books"`
}

func ReadToml(path string) DB {
	fmt.Print("Target Path", path)
	doc, readErr := os.ReadFile(path)
	if readErr != nil {
		fmt.Print("File Read Error", readErr)
		var empty DB
		return empty
	}
	var db DB
	err := toml.Unmarshal([]byte(doc), &db)
	if err != nil {
		fmt.Print("Toml Read Error", err)
		var empty DB
		return empty
	}
	return db
}

func SaveToml(db DB, path string) bool {
	b, err := toml.Marshal(db)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(path, b, 0644)
	if err != nil {
		fmt.Print("File Write Error", err)
		return false
	}
	return true
}
