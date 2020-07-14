package CloryTools

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func CheckOS() (OS string){
	os := runtime.GOOS
	return os
}

func SearchPattern(startPoint string, pattern string)(path []string){
	var matches []string
	err := filepath.Walk(startPoint,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.Contains(path, pattern){
				fmt.Println(path)
				matches = append(matches, path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return matches
}
