package tmenu

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type MenuItem struct {
	*tview.Box
	Title    string
	SubItems []*MenuItem
}

func NewMenuItem(title string) *MenuItem {
	return &MenuItem{
		Box:      tview.NewBox(),
		Title:    title,
		SubItems: make([]*MenuItem, 0),
	}
}

func (menuItem *MenuItem) AddItem(item *MenuItem) *MenuItem {
	menuItem.SubItems = append(menuItem.SubItems, item)
	return menuItem
}

func (menuItem *MenuItem) Draw(screen tcell.Screen) {
	menuItem.Box.DrawForSubclass(screen, menuItem)
	x, y, _, _ := menuItem.GetInnerRect()
	tview.PrintSimple(screen, menuItem.Title, x, y)
}

type SubMenu struct {
	*tview.Box
	Items            []*MenuItem
	offsetX, offsetY int
}

func NewSubMenu(items []*MenuItem) *SubMenu {
	subMenu := &SubMenu{
		Box:   tview.NewBox(),
		Items: items,
	}
	subMenu.SetBorder(true)
	return subMenu
}

func (subMenu *SubMenu) Draw(screen tcell.Screen) {
	log.Println("Printing the sub menu")
	subMenu.Box.DrawForSubclass(screen, subMenu)
	_, y, _, _ := subMenu.GetInnerRect()

	for _, item := range subMenu.Items {
		tview.PrintSimple(screen, item.Title, 1, 0+y) //x, y+i)
	}
}

type MenuBar struct {
	*tview.Box
	MenuItems     []*MenuItem
	subMenu       *SubMenu // sub menu if not nil will be drawn
	currentOption int
}

func NewMenuBar() *MenuBar {
	return &MenuBar{
		Box:       tview.NewBox(),
		MenuItems: make([]*MenuItem, 0),
	}
}

func (menuBar *MenuBar) AfterDraw() func(tcell.Screen) {
	return func(screen tcell.Screen) {
		if menuBar.subMenu != nil {
			menuBar.subMenu.Draw(screen)
		}
	}
}

func (menuBar *MenuBar) AddItem(item *MenuItem) *MenuBar {
	menuBar.MenuItems = append(menuBar.MenuItems, item)
	return menuBar
}

func (menuBar *MenuBar) Draw(screen tcell.Screen) {
	menuBar.Box.DrawForSubclass(screen, menuBar)

	x, y, width, _ := menuBar.GetInnerRect()

	for i := 0; i < width; i += 1 {
		screen.SetContent(x+i, y, tcell.RuneBlock, nil, tcell.StyleDefault)
	}

	menuItemOffset := 1
	for _, mi := range menuBar.MenuItems {
		itemLen := len([]rune(mi.Title))
		mi.SetRect(menuItemOffset, y, itemLen, 1)
		mi.Draw(screen)
		menuItemOffset += itemLen + 1
	}
	if menuBar.subMenu != nil {
		//	menuBar.subMenu.Draw(screen)
	}
}

func (menuBar *MenuBar) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	return menuBar.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
		switch event.Key() {
		case tcell.KeyLeft:
			menuBar.currentOption--
			if menuBar.currentOption < 0 {
				menuBar.currentOption = -1
			}
		case tcell.KeyRight:
			menuBar.currentOption++
			if menuBar.currentOption >= len(menuBar.MenuItems) {
				menuBar.currentOption = len(menuBar.MenuItems) - 1
			}
		}
	})
}

func (p *MenuBar) MouseHandler() func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive) {
	//log.Println("Mouse handler called")
	return p.WrapMouseHandler(func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive) {
		if !p.InRect(event.Position()) {
			return false, nil
		}
		//log.Println("Mouse captured in the menu bar")
		// Pass mouse events down.
		for _, item := range p.MenuItems {
			consumed, capture = item.MouseHandler()(action, event, setFocus)
			if consumed {
				//log.Println("Mouse captured by a menu item")
				p.subMenu = NewSubMenu(item.SubItems)
				x, y, _, _ := item.GetRect()
				p.subMenu.Box.SetRect(x+1, y+1, 15, 10)
				return
			}
		}

		// ...handle mouse events not directed to the child primitive...
		return true, nil
	})
}

func (menuBar *MenuBar) Focus(delegate func(p tview.Primitive)) {
	if menuBar.subMenu != nil {
		delegate(menuBar.subMenu)
	} else {
		menuBar.Box.Focus(delegate)
		menuBar.subMenu = nil
	}
}
