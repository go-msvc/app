package web

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/go-msvc/app"
	"github.com/go-msvc/errors"
	"github.com/go-msvc/logger"
)

var log = logger.New().WithLevel(logger.LevelDebug)

func Serve(addr string, app app.App) error {
	http.Handle("/", server{app: app})
	if err := http.ListenAndServe(addr, nil); err != nil {
		return errors.Wrapf(err, "failed to serve on address(%s)", addr)
	}
	return nil
}

type server struct {
	app app.App
	//sessions app.Sessions
}

type CtxSession struct{}

func (s server) ServeHTTP(httpRes http.ResponseWriter, httpReq *http.Request) {
	log.Debugf("HTTP %s %s", httpReq.Method, httpReq.URL.Path)

	//todo: see if should start a new session or use existing session

	//start new session for this app
	//s := sessions.New(s.app)

	//continue existing session:
	//sid := ...
	//s := sessions.Get(sid)
	//todo: verify that correct app is supported in this version
	//app := ...

	//render current page from the app
	// w := screen{
	// 	w: httpRes,
	// }
	ctx := context.Background()
	ctx = context.WithValue(ctx, CtxSession{}, s)

	//render page as web content
	//s.app.Render(w, ctx)
	d := s.app.(app.DisplayItem)
	httpRes.Header().Set("Content-Type", "text/html")
	httpRes.Write([]byte("<HTML><BODY><UL>"))
	render(httpRes, ctx, d)
	httpRes.Write([]byte("</UL></BODY></HTML>"))
}

func render(httpRes http.ResponseWriter, ctx context.Context, d app.DisplayItem) {
	switch t := d.(type) {
	case app.DisplayButton:
		s := t.Text()(ctx)
		httpRes.Write([]byte(fmt.Sprintf("<li><a href=\"/select/xxx\">%s</a></li>", s)))
	case app.DisplayText:
		s := t.Text()(ctx)
		httpRes.Write([]byte(fmt.Sprintf("<p>%s</p>", s)))
	case app.DisplaySection:
		s := t.Heading()(ctx)
		httpRes.Write([]byte(fmt.Sprintf("<h1>%s</h1>", s)))
		for _, i := range t.Items() {
			render(httpRes, ctx, i)
		}
	default:
		log.Errorf("cannot display %T", t)
	}
}

// implements app.Screen
type screen struct {
	w io.Writer
}

func (sc screen) P(s string) {

}
func (sc screen) H(level int, s string) {}
func (sc screen) B(s string)            {}
