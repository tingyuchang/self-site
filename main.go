package main

import (
	"self-site/model"
	"self-site/pkg/util"
	"self-site/setting"
	"self-site/routers"
)

func init() {
	setting.Setup()
	model.Setup()
	util.Setup()
}

func main() {
	r := routers.SetupRouter()
	r.Run()
}
