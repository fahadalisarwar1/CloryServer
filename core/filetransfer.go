package core

import (
	"encoding/gob"
	"fmt"
)

type FileStruct struct {
	FileName           string
	FileSize           int
	FileContent        []byte
	EncryptionKey	   []byte
}

func (srv *Server)SelectFile2upload(){
	
	fileName := "random.txt"
	srv.UploadFile(fileName)
}

func (s *Server) UploadFile(fileName string){

	data, err := s.ReadFileContent(fileName)
	DisplayError(err)
	fs := &FileStruct{}

	key := "password"
	fs.EncryptionKey = []byte(key)
	fmt.Println("Before Encryption, ", data)

	encryptedData := Encrypt(data, key)
	fmt.Println("After encryption, ", encryptedData)
	fs.FileName = fileName
	fs.FileContent = encryptedData
	encoder := gob.NewEncoder(s.Conn)
	err = encoder.Encode(fs)
	DisplayError(err)

}