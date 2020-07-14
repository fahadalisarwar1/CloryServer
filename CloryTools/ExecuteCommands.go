package CloryTools

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"

)

func ExecuteSystemCommand(command string, args ...string)(output string, err error){
	var out bytes.Buffer
	var stderr bytes.Buffer

	var cmd_instance *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd_instance = exec.Command(command, args...)
		//cmd_instance.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	} else {

		cmd_instance = exec.Command(command, args...)
	}
	cmd_instance.Stdout = &out
	cmd_instance.Stderr = &stderr
	err = cmd_instance.Run()
	if err != nil {
		fmt.Println()
		output = output + stderr.String()
		fmt.Println("[+] unable to execute command")
	} else {
		output = output + out.String()
	}
	return
}
