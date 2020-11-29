package wechat

import (
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"sort"
)

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
	for _, opt := range opts {
		opt(conf)
	}
	return &App{
		Config: conf,
	}
}

// App ..
type App struct {
	*Config
}

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

// Offiaccount 公众号
func (a *App) Offiaccount() *Offiaccount {
	return &Offiaccount{
		App: a,
	}
}

type httpHandler interface {
	Handler(w http.ResponseWriter, r *http.Request)
}

// Offiaccount ..
type Offiaccount struct {
	*App
	httpHandler
	Reply (func())
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
