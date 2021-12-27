package main

import (
	"log"

	"github.com/Racinettee/tmenu"
	"github.com/rivo/tview"
)

func main() {
	fileMenu := tmenu.NewMenuItem("File")
	fileMenu.AddItem(tmenu.NewMenuItem("New File").SetOnClick(func(*tmenu.MenuItem) { log.Println("New file was clicked!") }))
	fileMenu.AddItem(tmenu.NewMenuItem("Open File"))
	fileMenu.AddItem(tmenu.NewMenuItem("Save File"))
	fileMenu.AddItem(tmenu.NewMenuItem("Close File"))
	fileMenu.AddItem(tmenu.NewMenuItem("Exit"))
	editMenu := tmenu.NewMenuItem("Edit")
	editMenu.AddItem(tmenu.NewMenuItem("Copy"))
	editMenu.AddItem(tmenu.NewMenuItem("Cut"))
	editMenu.AddItem(tmenu.NewMenuItem("Paste"))

	menuBar := tmenu.NewMenuBar().
		AddItem(fileMenu).
		AddItem(editMenu)

	menuBar.SetRect(0, 0, 100, 15)

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(menuBar, 1, 1, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Hello, world!"), 0, 1, true)

	app := tview.NewApplication().EnableMouse(true).SetRoot(flex, true).SetFocus(flex).SetAfterDrawFunc(menuBar.AfterDraw())

	if err := app.Run(); err != nil {
		panic(err)
	}
}
