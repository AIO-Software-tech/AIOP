package main

//Done
//? Imports
import (
    "fmt"
    "os"
    "syscall"
	"time"
    "golang.org/x/term"
)

//Done
//@ Who It Was Made By And The Current Version
func main() {

	fmt.Print("\n")
	fmt.Print("All-In-One By Imre Kiss And Oliver Boucher \n")
	fmt.Print("Version 0.0.1 REV-1 2022 \n")
	fmt.Print("\n")

	LoginInSystem()
}

//Done
//! Username and Password System:
func LoginInSystem() {
	fmt.Print("Username: ")
	var Username string
	fmt.Scanln(&Username)
    fmt.Print("Password: ")
    bytepw, err := term.ReadPassword(int(syscall.Stdin))
    if err != nil {
        os.Exit(1)
    }
    Password := string(bytepw)
	fmt.Scanln(&Password)

	if Username == "Ollie" && Password == "123"{
		MainMenu(Username)
	}
}

//ToDo
//? Main Menu
//This Prints The Main Menu And Takes The Selection Input
//MMS Stands For Main Menu Selection
func MainMenu(Username string) {
	fmt.Print("Hello ", Username, "\n")
	fmt.Print("\n")
	fmt.Print("Selection Menu: Please Select One Of The Options:\n")
	fmt.Print("▣───────────────────────────────────────────────▣\n")
	fmt.Print("│                                               │\n")
	fmt.Print("│   1. = Personal                               │\n")
	fmt.Print("│                                               │\n")
	fmt.Print("│   2. = Calculations                           │\n")
	fmt.Print("│                                               │\n")
	fmt.Print("│   3. = Games                                  │\n")
	fmt.Print("│                                               │\n")
	fmt.Print("│   4. = Error Codes                            │\n")
	fmt.Print("│                                               │\n")
	fmt.Print("▣───────────────────────────────────────────────▣\n")
	fmt.Print("\n")
	fmt.Print("Please Select An Option From 1 To 4 \n")
	fmt.Print("Option: \n")
	var MMS int
	fmt.Scanln(&MMS)

	if MMS >= 6 || MMS <= 0{
		fmt.Print("Please Try Again \n")
		fmt.Print("\n")
		MainMenu(Username)
	} else if MMS == 1{
		PersonalMenu()
	} else if MMS == 2{
		CalculationsMenu()
	} else if MMS == 3{
		GamesMenu()
	} else if MMS == 4{
		ErrorCodes(Username)
	} else{
		fmt.Print("AIO Error Code 3 \n")
	}
	
}

//ToDo
//? Personal Menu
func PersonalMenu() {

}

//ToDo
//? Calculations Menu
func CalculationsMenu() {

}

//ToDo
//? Games Menu
func GamesMenu() {

}

//ToDo
//? Error Codes
func ErrorCodes(Username string) {
	fmt.Print("▣──────────────────────────────────▣\n")
	fmt.Print("│                                   │\n")
	fmt.Print("│   1. = Invalid option/chose       │\n")
	fmt.Print("│                                   │\n")
	fmt.Print("│   2. = Invalid Username or Pass   │\n")
	fmt.Print("│                                   │\n")
	fmt.Print("│   3. = Invalid input              │\n")
	fmt.Print("│                                   │\n")
	fmt.Print("▣──────────────────────────────────▣\n")
	fmt.Print("\n")
	fmt.Print("Please Wait\n")
	fmt.Print("\n")
	time.Sleep(2 * time.Second)
	MainMenu(Username)
}
