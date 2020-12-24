package main

import (
	"net/http"

	"github.com/wepkg/wechat"
	"github.com/wepkg/wechat/message"
)

func main() {
	initWechat()
	http.HandleFunc("/wechat/notify", Handler)
	http.ListenAndServe(":8000", nil)
}

var mp *wechat.Offiaccount

func initWechat() {
	app := wechat.New(
		wechat.Appid("<appid>"),
		wechat.Secret("<secret>"),
		wechat.Token("<token>"),
		wechat.AESKey("<encodingAESKey>"), //EncdingAESKey
	)
	mp := app.Offiaccount()
	mp.SetReceivingHandler(func(msg message.Context) {
		msg.Event
	})
}

// Handler ..
func Handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "hello world")
	mp.Notify(w, r)
}
