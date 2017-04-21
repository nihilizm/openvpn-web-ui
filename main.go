package main

import (
	"github.com/nihilizm/openvpn-web-ui/lib"
	_ "github.com/nihilizm/openvpn-web-ui/routers"
	"github.com/astaxie/beego"
)

func main() {
	lib.AddFuncMaps()
	beego.Run()
}
