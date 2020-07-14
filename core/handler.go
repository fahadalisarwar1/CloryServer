package core

import (
	"fmt"
)



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
		userInput, err := GetUserInput()
		DisplayError(err)
		//fmt.Println(userInput)

		err = s.SendData(userInput)	// sending so client is at the same level
		DisplayError(err)

		switch {
		case userInput == "stop" || userInput == "99":			// exit condition
			fmt.Println("[-] Exiting")
			MainLoop = false
			break
		case userInput == "1":
			fmt.Println("Transfer encrypted file")		//
			s.SelectFile2upload()

		default:
			fmt.Println("[-] Invalid, you entered : ", userInput)
		}
	}
}
