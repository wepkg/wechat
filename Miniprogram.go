package wechat

import (
	"context"
	"encoding/json"
	"net/url"
)

// Miniprogram ..
type Miniprogram struct {
	*App
}

// Session ..
type Session struct {
	*Result
	Openid     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` //会话密钥
	Unionid    string `json:"unionid"`     // 用户在开放平台的唯一标识符，在满足 UnionID 下发条件的情况下会返回，详见 UnionID 机制说明。
}

// Code2Session ...
func (a *Miniprogram) Code2Session(ctx context.Context, code string) (*Session, error) {
	const Code2SessionURL = "/sns/jscode2session"
	u := url.Values{}
	u.Add("grant_type", "authorization_code")
	u.Add("appid", a.Config.Appid)
	u.Add("secret", a.Config.Secret)
	u.Add("js_code", code)
	resp, err := a.Client.Get(ctx, Code2SessionURL, u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data := &Session{}
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, err
	}
	return data, nil
}
