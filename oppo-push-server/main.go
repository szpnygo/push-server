package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/jasonlvhit/gocron"
	"github.com/szpnygo/go-tools"
	"neobaran.com/server/push/oppo/auth"
	"os"
)

type PushCenter struct {
}

func main() {
	os.Mkdir("logs", os.ModePerm)

	go func() {
		gocron.Every(12).Hours().Do(auth.RequestAuthToken)
	}()
	auth.RequestAuthToken()

	logs.SetLogger(logs.AdapterFile, `{
		"filename":"logs/oppo-push-center.log",
		"maxdays": 10,
		"daily": true,
		"color": true,
		"level": 7
	}`)
	logs.Info("order center server version:" + beego.AppConfig.String("version"))
	go func() {
		neo.InitJsonRpc(new(PushCenter), ":"+beego.AppConfig.String("rpcport"))
	}()
	beego.Run()
}
