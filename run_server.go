package main

import (
	"github.com/fahadalisarwar1/CloryServer/core"
)

func main(){
	sc := core.Server{}
	sc.IP="192.168.0.13"
	sc.Port="8080"
	err := sc.CreateConnection()
	core.DisplayError(err)

	err = sc.Connect()
	core.DisplayError(err)

	sc.HandleConnection()

}

