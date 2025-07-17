package main

import (
	"fmt"
	"github.com/ncruces/zenity"
	"log"
	"os"
	"regexp"
)


func main() {
	fmt.Println("Music FX - A simple music effects application")
	tmatch := regexp.MustCompile(`[^\/]+\.mp3$`)
	menu := &Menu{
		Prompt:    "Select an option:",
		CursorPos: 0,
		MenuItems: make([]*MenuItem, 0),
	}
	
	filename, err := zenity.SelectFile(
		zenity.Title("Select an MP3 file"),
		zenity.FileFilter{
			Name:     "MP3 Audio",
			Patterns: []string{"*.mp3"},
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Read %d bytes from %s\n", len(data), tmatch.FindString(filename))
	menu.AddItem("Play Music", "1")
	menu.AddItem("Pause Music", "2")
	menu.AddItem("Stop Music", "3")
	menu.AddItem("Import Music", "4")
	menu.AddItem("About", "5")
	menu.AddItem("Exit", "")
	for {
		res := menu.Display()
		if res == "" {
			break;
		} 
	}
}

