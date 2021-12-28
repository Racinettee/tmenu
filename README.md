# tmenu 
tmenu is tview based library that provides a main menu bar widget

To use tmenu follow these steps per cmd/main.go for example

1. Create the MenuItems that will be your menu bars main items
2. For each main item add sub MenuItems
3. Create the MenuBar
4. Add each main item to the menu bar in the desired order
5. Add the menu bar to the desired container, such as a flex - don't allow its size to be recomputed
6. Install the MenuBars AfterDraw() function as part of your AfterDraw routine, or as the main one to the application
7. Run the application

[example](cmd/main.go)