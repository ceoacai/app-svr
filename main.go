package main

import (
	_ "app-svr/docs"
	_ "app-svr/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}

	corsFunc := func(ctx *context.Context) {
		ctx.Output.Header("Access-Control-Allow-Methods", "GET,POST,PUT,OPTIONS")
		ctx.Output.Header("Access-Control-Allow-Origin", "*")
	}

	beego.InsertFilter("*", beego.BeforeRouter, corsFunc)

	beego.Run()
}
