package wechat

// Offiaccount ..
type Offiaccount struct {
	*App
	Reply (func())
}

func (a *App) GetToken() {
	a.Auth.GetToken()
}

//Client
