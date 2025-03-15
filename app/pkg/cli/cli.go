package cli

import (
	accountMenu "cli-banking-app/pkg/menus/account"
	mainMenu "cli-banking-app/pkg/menus/main"
	userMenu "cli-banking-app/pkg/menus/user"
	"fmt"
)

var mm *mainMenu.MainMenu
var um *userMenu.UserMenu
var am *accountMenu.AccountMenu

func init() {
	mm = &mainMenu.MainMenu{}
	um = &userMenu.UserMenu{}
	am = &accountMenu.AccountMenu{}

	mm.SubMenu = um
	um.SubMenu = am
}

func MainScreen() {
	// mm := &mainMenu.MainMenu{}
	// um := &userMenu.UserMenu{}
	// am := &accountMenu.AccountMenu{}
	
	// mm.SubMenu = um
	// um.SubMenu = am

	for {
		mm.ShowMenuOptions()

		if mm.Login() {
			fmt.Println("Login")

			// call customerLoginHandler to check 
			// if customer exists
			userScreen()
		}

		if mm.Register() {
			// call customerRegisterHandler to allow user
			// to register into the bank
			// create an "input" package for user input
			fmt.Println("Create Account")
		}

		if mm.Exit() {
			// exit the application
			fmt.Println("Exit")
			break
		}
	}
}

func userScreen() {
	for {
		mm.ShowSubMenu()
		// um.ShowMenuOptions()

		if um.CreateAccount() {
			fmt.Println("Create account")
		}

		if um.AccountOperations() {
			accountScreen()
		}

		if um.Logout() {
			fmt.Println("Logout")
			break
		}
	}
}

func accountScreen() {
	for {
		um.ShowSubMenu()

		if am.Deposit() {
			fmt.Println("Deposit")
		}

		if am.ReturnUserMenu() {
			fmt.Println("Return to user menu")
			break
		}
	}
}