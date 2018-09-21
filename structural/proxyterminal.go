package main

import (
	"os"
	"fmt"
	"time"
)

// For example:
// we must a execute some command
// so before that we must to create new terminal session
// and provide our user name and command
func main() {
	// Create new instance of Proxy terminal
	fmt.Println("Enter the User : ")
	var user string
	fmt.Scanf("%s",&user)
	t, err := NewTerminal(user)
	if err != nil {
		// panic: User cant be empty
		// Or
		// panic: You (badUser) are not allowed to use terminal and execute commands
		panic(err.Error())
	}

	// Execute user command
	fmt.Println("Enter the Command to Execute : ")
	var res string
	fmt.Scanf("%s",&res)
	excResp, excErr := t.Execute(res) // Proxy prints to STDOUT -> PROXY: Intercepted execution of user (user), asked command (say_hi)
	if excErr != nil {
		fmt.Printf("ERROR: %s\n", excErr.Error()) // Prints: ERROR: I know only how to execute commands: say_hi, man
	}

	// Show execution response
	fmt.Println(excResp) // Prints: dilip@go_term$: Hi user
}

/*
	From that it's can be different terminals 
	realizations with different methods, propertys, yda yda...
*/

/*
 ITerminal is interface, it's a public method whose implemented in 
Terminal(Proxy) and user Terminal
*/
type ITerminal interface {
	Execute(cmd string) (resp string, err error)
}

/*
 UserTerminal for example:
*/
type UserTerminal struct {
	// user is a current authorized user
	User string
}

// Execute just runs known commands for current authorized user
func (gt *UserTerminal) Execute(cmd string) (resp string, err error) {
	// Set "terminal" prefix for output
	prefix := fmt.Sprintf("%s@admin$:", gt.User) 

	// Execute some asked commands if we know them
	switch cmd {
	case "man":
		resp = fmt.Sprintf("%s Visit 'https://golang.org/doc/' for Golang documentation", prefix)
	case "date":
		current_time := time.Now().Local()
		resp = fmt.Sprintf("%s today date %s ",prefix,current_time.Format("2006-01-02"))
	case "time":
		t := time.Now()
		resp=fmt.Sprintf("Current Time is : %s",t.Format("15:04:05"))
	default:
		err = fmt.Errorf("%s Unknown command", prefix)
	}

	return
}

/*
	And now we will create owr proxy to deliver user and commands to specific objects
*/

// Terminal is a implementation of Proxy, it's  validates and sends data to UserTerminal
// As example before send commands, user must be authorized
type Terminal struct {
	currentUser    string
	terminal *UserTerminal
}

// NewTerminal creates new instance of terminal
func NewTerminal(user string) (t *Terminal, err error) {
	// Check user if given correctly
	if user == "" {
		err = fmt.Errorf("User cant be empty")
		return
	}

	// Before we execute user commands, we validate current user, if he have rights to do it
	if authErr := authorizeUser(user); authErr != nil {
		err = fmt.Errorf("You (%s) are not allowed to use terminal and execute commands", user)
		return
	}

	// Create new instance of terminal and set valid user
	t = &Terminal{currentUser: user}

	return
}

// Execute intercepts execution of command, implements authorizing user, validates it and
// poxing command to real terminal (terminal) method
func (t *Terminal) Execute(command string) (resp string, err error) {
	// If user allowed to execute send commands then, for example we can decide which terminal can be used, remote or local etc..
	// but for example we just creating new instance of terminal,
	// set current user and send user command to execution in terminal
	t.terminal = &UserTerminal{User: t.currentUser}

	// For example our proxy can log or output intercepted execution... etc
	fmt.Printf("PROXY: Intercepted execution of user (%s), asked command (%s)\n", t.currentUser, command)

	// Transfer data to original object and execute command
	if resp, err = t.terminal.Execute(command); err != nil {
		err = fmt.Errorf("I know only how to execute commands: man, date, time")
		return
	}

	return
}

// authorize validates user right to execute commands
func authorizeUser(user string) (err error) {
	// As we use terminal like proxy, then
	// we will intercept user name to validate if it's allowed to execute commands
	if user == "dilip" || user =="DILIP" {// Do some logs, notifications etc...
		return
	}else if user =="VIJAY" || user == "vijay"{
		return
	}else{
		fmt.Println("Insufficent Privilege ")
		os.Exit(1)
	}

	return
}