package core

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getUserInput()(userInput string, err error){
	reader := bufio.NewReader(os.Stdin)
	userInput, err = reader.ReadString('\n')
	userInput = strings.TrimSuffix(userInput, "\n")
	return userInput, err
}

func printUserOptions(){
	fmt.Println("\t\t[1] Send encrypted file")
	fmt.Println("\t\t[99] Exit")

}

func (s *Server) HandleConnection(){
	/*
		Method to handle connection

	*/

	fmt.Println("[+] Connection established with ", s.Conn.RemoteAddr().String())

	MainLoop := true	// for continously asking user for input
	for MainLoop {
		printUserOptions()
		fmt.Print("[+] What do you want to do: ")
		userInput, err := getUserInput()
		DisplayError(err)
		//fmt.Println(userInput)

		err = s.SendData(userInput)	// sending so client is at the same level
		DisplayError(err)

		switch {
		case userInput == "stop" || userInput == "99":
			fmt.Println("Exiting")
			MainLoop = false
			break
		case userInput == "1":
			fmt.Println("Transfer encrypted file")
			s.SelectFile2upload()

		default:
			fmt.Println("you entered : ", userInput)
		}
	}
}
