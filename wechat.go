package wechat

import (
	"github.com/wepkg/wechat/auth"
	"github.com/wepkg/wechat/util"
)

const WechatURL = "https://api.weixin.qq.com"

//Config ..
type Config struct {
	Appid  string
	Secret string
	Token  string
	AESKey string
}

// Option ..
type Option func(*Config)

// Appid set Appid
func Appid(v string) Option {
	return func(c *Config) {
		c.Appid = v
	}
}

// Secret set Secret
func Secret(v string) Option {
	return func(c *Config) {
		c.Secret = v
	}
}

// Token set Token
func Token(v string) Option {
	return func(c *Config) {
		c.Token = v
	}
}

// AESKey set AESKey
func AESKey(v string) Option {
	return func(c *Config) {
		c.AESKey = v
	}
}

// New a app service
func New(opts ...Option) *App {
	conf := &Config{}
	// conf.WechatURL = WechatURL
	for _, opt := range opts {
		opt(conf)
	}
	client, _ := util.NewClient(
		util.WithEndpointBase(WechatURL),
	)
	return &App{
		Config: conf,
		Client: client,
	}
}

// App ..
type App struct {
	Config *Config
	Auth   auth.Auth
	Client *util.Client
}

// Offiaccount 公众号
func (a *App) Offiaccount() *Offiaccount {
	return &Offiaccount{App: a}
}

// Miniprogram 小程序
func (a *App) Miniprogram() *Miniprogram {
	return &Miniprogram{App: a}
}

// Result ..
type Result struct {
	Errcode int    `json:"errcode"` //错误码
	Errmsg  string `json:"Errmsg"`  //错误信息
}

// -1	系统繁忙，此时请开发者稍候再试
// 0	请求成功
// 40029	code 无效
// 45011	频率限制，每个用户每分钟100次

// IsError ..
func (r *Result) IsError() bool {
	if r.Errcode != 0 {
		return true
	}
	return false
}

/*
// Validate ..
func (a *App) Validate(h *handler) bool {
	signature := h.Query("signature")
	timestamp := h.Query("timestamp")
	nonce := h.Query("nonce")
	// echostr := h.Query("echostr")
	return signature == sign(h.Token, timestamp, nonce)
}
func sign(params ...string) string {
	sort.Strings(params)
	h := sha1.New()
	for _, s := range params {
		_, _ = io.WriteString(h, s)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Handler ..
type service struct {
	w http.ResponseWriter
	r *http.Request
}

//Query ..
func (h *service) Query(k string) string {
	return h.r.URL.Query().Get(k)
}

func (h *service) Render(code int, r render) {
	h.w.WriteHeader(code)
	header := h.w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = r.ContentType()
	}

	h.w.Write(r.Body())
}

type render interface {
	ContentType() []string
	Body() []byte
}

type plainRender []byte

func (r plainRender) ContentType() []string {
	return []string{"text/plain; charset=utf-8"}
}
func (r plainRender) Body() []byte {
	return r
}

type xmlRender []byte

func (r xmlRender) ContentType() []string {
	return []string{"application/xml; charset=utf-8"}
}
func (r xmlRender) Body() []byte {
	return r
}

type httpHandler interface {
	Handler(w http.ResponseWriter, r *http.Request)
}



// Handler ..
func (a *Offiaccount) Handler(w http.ResponseWriter, r *http.Request) {
	srv := &service{
		w: w,
		r: r,
	}
	if !a.Validate(srv) {
		return
	}
	echostr := srv.Query("echostr")
	if len(echostr) > 0 {
		srv.Render(200, plainRender(echostr))
		return
	}
}

// Notify 微信公众号(mp)
func (a *App) Notify() {

}
*/
