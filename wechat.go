package wechat

import (
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"sort"

	"github.com/wepkg/wechat/message"
)

type App struct {
	*Core
}

// New a Core service
func New(opts ...Option) *App {
	core := &Core{}
	for _, opt := range opts {
		opt(core)
	}
	return &App{
		Core: core,
	}
}

// Validate ..
// func (a *App) Validate(h *handler) bool {
// 	signature := h.Query("signature")
// 	timestamp := h.Query("timestamp")
// 	nonce := h.Query("nonce")
// 	// echostr := h.Query("echostr")
// 	return signature == sign(h.Token, timestamp, nonce)
// }

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
	if r != nil {
		header := h.w.Header()
		if val := header["Content-Type"]; len(val) == 0 {
			header["Content-Type"] = r.ContentType()
		}
		h.w.Write(r.Body())
	}
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

// Notify ..
func (a *App) Notify(w http.ResponseWriter, r *http.Request) {
	srv := &service{
		w: w,
		r: r,
	}
	if !a.Validate(srv) {
		srv.Render(400, nil)
		return
	}
	echostr := srv.Query("echostr")
	if len(echostr) > 0 {
		srv.Render(200, plainRender(echostr))
		return
	}
}

// type httpHandler interface {
// 	Handler(w http.ResponseWriter, r *http.Request)
// }

// Offiaccount ..
type Offiaccount struct {
	*App
	// httpHandler
	Reply            (func())
	ReceivingHandler ReceivingHandler
}

// ReceivingHandler ..
type ReceivingHandler func(message.Context)

// SetReceivingHandler ..
func (o *Offiaccount) SetReceivingHandler(h ReceivingHandler) {
	o.ReceivingHandler = h
}
