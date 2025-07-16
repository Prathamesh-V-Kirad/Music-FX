package main

import (
	"fmt"
	"github.com/pkg/term"
	"github.com/buger/goterm"
)

var up byte = 65
var down byte = 66
var esc byte = 27
var enter byte = 13

var keys = map[byte]bool{
	up:    true,
	down:  true,
}

type MenuItem struct {
	Title string
	ID string
} 

type Menu struct {
	Prompt string
	CursorPos int
	MenuItems []*MenuItem
}

func main() {
	fmt.Println("Music FX - A simple music effects application")
	menu := &Menu{
		Prompt:    "Select an option:",
		CursorPos: 0,
		MenuItems: make([]*MenuItem, 0),
	}

	menu.AddItem("Play Music", "1")
	menu.AddItem("Pause Music", "2")
	menu.AddItem("Stop Music", "3")
	menu.AddItem("Import Music", "4")
	menu.AddItem("About", "5")
	menu.AddItem("Exit", "6")
	menu.Display()
}


func (m *Menu) AddItem(title string, id string) *Menu{
	item := &MenuItem{
		Title: title,
		ID:    id,
	}
	m.MenuItems = append(m.MenuItems, item)
	return m
}

func getInput() byte{
	inputBytes := make([]byte, 3)
	t , _ := term.Open("/dev/tty")
	err := term.RawMode(t)
	if err != nil {
		fmt.Println("Error opening terminal:", err)
		
	}
	var inputLen int
	inputLen , err = t.Read(inputBytes)
	t.Restore()
	t.Close()
	if err != nil {		
		fmt.Println("Error reading input:", err)
	}
	if inputLen == 3 {
		if (keys[inputBytes[2]]){
			return inputBytes[2]
		}
	}else{
		return inputBytes[0]
	}
	return 0;	
}


func (m *Menu) renderMenuItems(redraw bool) {
	if redraw {
		fmt.Printf("\033[%dA", len(m.MenuItems)-1)
	}

	for index, menuItem := range m.MenuItems {
		var newline = "\n"
		if index == len(m.MenuItems)-1 {
			newline = ""
		}

		menuItemText := menuItem.Title
		cursor := "  "
		if index == m.CursorPos {
			cursor = goterm.Color("> ", goterm.YELLOW)
			menuItemText = goterm.Color(menuItemText, goterm.YELLOW)
		}

		fmt.Printf("\r%s %s%s", cursor, menuItemText, newline)
	}
}

func (m *Menu) Display() string {
	defer func() {
		fmt.Printf("\033[?25h")
	}()

	fmt.Printf("%s\n", goterm.Color(goterm.Bold(m.Prompt)+":", goterm.CYAN))

	m.renderMenuItems(false)

	fmt.Printf("\033[?25l")

	for {
		keyCode := getInput()
		if keyCode == esc {
			return ""
		} else if keyCode == enter {
			menuItem := m.MenuItems[m.CursorPos]
			fmt.Println("\r")
			return menuItem.ID
		} else if keyCode == up {
			m.CursorPos = (m.CursorPos + len(m.MenuItems) - 1) % len(m.MenuItems)
			m.renderMenuItems(true)
		} else if keyCode == down  {
			m.CursorPos = (m.CursorPos + 1) % len(m.MenuItems)
			m.renderMenuItems(true)
		}
	}
}
