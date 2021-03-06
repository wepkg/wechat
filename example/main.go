package main

import (
	"net/http"

	"github.com/wepkg/wechat"
)

func main() {
	http.HandleFunc("/wechat/notify", Handler)
	http.ListenAndServe(":8000", nil)
}

// Handler ..
func Handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "hello world")
	app := wechat.New(
		wechat.Appid("<appid>"),
		wechat.Secret("<secret>"),
		wechat.Token("<token>"),
		wechat.AESKey("<encodingAESKey>"),
	)
	mp := app.Offiaccount()
	//PassiveUserReply_message
	mp.Reply(func() {

	})
	// mp.OnEvent(wechat.PassiveReply)
	mp.Handler(w, r)
}
