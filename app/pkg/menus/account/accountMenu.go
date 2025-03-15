package account_menu

import "github.com/nexidian/gocliselect"

type AccountMenu struct {
	optionId string
}

func menuOptions() *gocliselect.Menu {
	menu := gocliselect.NewMenu("Choose an operation to perform on your accounts:")
	
	menu.AddItem("Deposit money into an account", "deposit")
	menu.AddItem("Withdraw money from an account", "withdraw")
	menu.AddItem("Transfer money to another account", "transfer")
	menu.AddItem("Transfer money to another user's account", "transfer_user")
	menu.AddItem("Return to user menu", "user_menu")

	return menu
}

func (m *AccountMenu) ShowMenuOptions() {
	m.optionId = menuOptions().Display()
}

func (m *AccountMenu) ShowSubMenu() {

}

func (m *AccountMenu) Deposit() bool {
	return m.optionId == "deposit"
}

func (m *AccountMenu) Withdraw() bool {
	return m.optionId == "withdraw"
}

func (m *AccountMenu) TransferOwn() bool {
	return m.optionId == "transfer"
}

func (m *AccountMenu) Transfer() bool {
	return m.optionId == "transfer_user"
}

func (m *AccountMenu) ReturnUserMenu() bool {
	return m.optionId == "user_menu"
}