package main

import (
	"log"

	"github.com/Racinettee/tmenu"
	"github.com/rivo/tview"
)

func clickedMessageFn(msg string) func(*tmenu.MenuItem) {
	return func(*tmenu.MenuItem) { log.Printf("%v clicked\n", msg) }
}

func main() {
	fileMenu := tmenu.NewMenuItem("File")
	fileMenu.AddItem(tmenu.NewMenuItem("New File").SetOnClick(clickedMessageFn("New File")))
	fileMenu.AddItem(tmenu.NewMenuItem("Open File").SetOnClick(clickedMessageFn("Open File")))
	fileMenu.AddItem(tmenu.NewMenuItem("Save File").SetOnClick(clickedMessageFn("Save File")))
	fileMenu.AddItem(tmenu.NewMenuItem("Close File").SetOnClick(clickedMessageFn("Close File")))
	fileMenu.AddItem(tmenu.NewMenuItem("Exit").SetOnClick(clickedMessageFn("Exit")))
	editMenu := tmenu.NewMenuItem("Edit")
	editMenu.AddItem(tmenu.NewMenuItem("Copy").SetOnClick(clickedMessageFn("Copy")))
	editMenu.AddItem(tmenu.NewMenuItem("Cut").SetOnClick(clickedMessageFn("Cut")))
	editMenu.AddItem(tmenu.NewMenuItem("Paste").SetOnClick(clickedMessageFn("Paste")))

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
