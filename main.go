package main

import (
	_ "EftServer/boot"
	_ "EftServer/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
