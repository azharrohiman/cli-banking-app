package main_menu

import (
	homeMenu "cli-banking-app/pkg/menus"

	"github.com/nexidian/gocliselect"
)

type MainMenu struct {
	optionId string
	SubMenu homeMenu.Menu
}

func menu() *gocliselect.Menu {
	menu := gocliselect.NewMenu("Choose an option")

	menu.AddItem("Login with a user account", "login")
	menu.AddItem("Create a user account", "register")
	menu.AddItem("Exit", "exit")

	return menu
}

func (m *MainMenu) ShowMenuOptions() {
	m.optionId = menu().Display()
}

func (m *MainMenu) ShowSubMenu() {
	m.SubMenu.ShowMenuOptions()
}

func (m *MainMenu) Login() bool {
	return m.optionId == "login"
}

func (m *MainMenu) Register() bool {
	return m.optionId == "register"
}

func (m *MainMenu) Exit() bool {
	return m.optionId == "exit"
}