package sysApi

import (
	"../global"
	"fmt"
)

type SysRouter struct {
}

func (s *SysRouter) Login() {

}
func (s *SysRouter) AddSys() {
	global.Global["sys"] = "addSys"
	fmt.Println("sys中的打印：", global.Global["name"])
	global.Global["name"] = "sysName"
}
