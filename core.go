package wechat

import "net/http"

//Core ..
type Core struct {
	Appid  string
	Secret string
	Token  string
	AESKey string
	client *http.Client
}

// Option ..
type Option func(*Core)

// Appid set Appid
func Appid(v string) Option {
	return func(c *Core) {
		c.Appid = v
	}
}

// Secret set Secret
func Secret(v string) Option {
	return func(c *Core) {
		c.Secret = v
	}
}

// Token set Token
func Token(v string) Option {
	return func(c *Core) {
		c.Token = v
	}
}

// AESKey set AESKey
func AESKey(v string) Option {
	return func(c *Core) {
		c.AESKey = v
	}
}
