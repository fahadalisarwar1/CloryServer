package core
/*
AUTHOR:		 Fahad Ali SARWAR
DATE: 		 7/7/2020
VERSION:	 1.0-dev


*/

import (
	"bufio"
	"fmt"

	"net"
	"os"
)




type Server struct {
	/*
	This struct provides a basic structure for connecting with the client
	IP : Local IP address for listening for incoming connections
	Port: Local port for listening
	Conn: Connection object used for communicating and data transfer
	Listener: Same listener for accepting connections
	*/
	IP string
	Port string
	Conn net.Conn
	Listener net.Listener
}

func (s *Server) CreateConnection()(err error){
	/*
		method attached to server struct type
		used in conjunction with the server struct
		listens for incoming connection on specified port defined in the struct.
	*/
	LocalAddressPort := s.IP + ":" + s.Port
	fmt.Println("[+] Listening on  ", LocalAddressPort)
	s.Listener, err = net.Listen("tcp", LocalAddressPort)
	return err
}

func (s *Server) Connect() (err error) {
	/*
		method attached to server struct type
		used in conjunction with the server struct
		listens for incoming connection on specified port defined in the struct.
		Accepts the incoming connection
	*/
	s.Conn, err = s.Listener.Accept()

	return err
}

func (s *Server) SendData(data string)(err error){
	/*
		Sends the given string to the connected client. string must not include the \n character
	*/
	_, err = s.Conn.Write([]byte(data+"\n"))
	return err
}

func (s *Server) ReceiveData()(data string, err error){
	/*
		Receives a string from the connected client
		returns the string received without the "\n" character
	*/
	data, err = bufio.NewReader(s.Conn).ReadString('\n')
	return data, err
}

func (s *Server) ReadFileContent(fileName string)( []byte, error){
	/*
		Given a filename as string, reads the contents of the file in buffer and returns them as array of bytes
	*/
	file, err := os.Open(fileName)
	DisplayError(err)
	defer file.Close()
	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil, statsErr
	}
	var size int64 = stats.Size()
	bytes := make([]byte, size)

	bufr := bufio.NewReader(file)
	_,err = bufr.Read(bytes)
	return bytes, err
}

func (s *Server) GetEncryptedContent(fileName string)(data []byte, key []byte, err error){
	/*
		Method to get encrypted content of a given file
	*/
	return

}