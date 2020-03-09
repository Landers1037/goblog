/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: cloudp
*/
package cmd

import (
	"os/exec"
)
func cmdMem() string {
	cmdout := exec.Command("")
	op,err:= cmdout.Output()
	if err!=nil{
		return "getmem failed"
	}
	return string(op)
}
