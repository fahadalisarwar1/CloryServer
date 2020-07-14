package main

import (
	"fmt"
	"github.com/fahadalisarwar1/CloryServer/CloryTools"
	"github.com/fahadalisarwar1/CloryServer/core"
)

func main(){
	result, err := CloryTools.ExecuteSystemCommand("ifconfig")
	core.DisplayError(err)
	fmt.Println(result)
}
