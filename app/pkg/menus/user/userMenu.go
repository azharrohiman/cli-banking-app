package user_menu

import (
	homeMenu "cli-banking-app/pkg/menus"

	"github.com/nexidian/gocliselect"
)

type UserMenu struct {
	optionId string
	SubMenu homeMenu.Menu
}

func menuOptions() *gocliselect.Menu {
	menu := gocliselect.NewMenu("How may we help you?")

	menu.AddItem("Create a savings account", "create_acc")
	menu.AddItem("Delete a savings account", "delete_acc")
	menu.AddItem("View account operations", "acc_op")
	menu.AddItem("Logout", "logout")

	return menu
}

func (m *UserMenu) ShowMenuOptions() {
	m.optionId = menuOptions().Display()
}

func (m *UserMenu) ShowSubMenu() {
	m.SubMenu.ShowMenuOptions()
}

func (m *UserMenu) CreateAccount() bool {
	return m.optionId == "create_acc"
}

func (m *UserMenu) DeleteAccount() bool {
	return m.optionId == "delete_acc"
}

func (m *UserMenu) AccountOperations() bool {
	return m.optionId == "acc_op"
}

func (m *UserMenu) Logout() bool {
	return m.optionId == "logout"
}