package CloryTools

import "os"

func GetExecutableName()(exe string){
	return os.Args[0]
}
