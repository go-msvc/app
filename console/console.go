package console

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"

	"github.com/go-msvc/app"
	"github.com/go-msvc/app/service"
	"github.com/go-msvc/errors"
	"github.com/go-msvc/logger"
)

var log = logger.New().WithLevel(logger.LevelDebug)

func call(path string, req interface{}, resTmpl interface{}) (res interface{}, err error) {
	jsonReq, _ := json.Marshal(req)
	body := bytes.NewReader(jsonReq)
	httpReq, _ := http.NewRequest(http.MethodPost, "http://localhost:11111/"+path, body)
	httpRes, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		panic(err)
	}
	for n, v := range httpRes.Header {
		log.Debugf("Header: %v: %v", n, v)
	}
	resPtrValue := reflect.New(reflect.TypeOf(resTmpl))
	if err := json.NewDecoder(httpRes.Body).Decode(resPtrValue.Interface()); err != nil {
		return nil, errors.Wrapf(err, "failed to decode JSON response body")
	}
	return resPtrValue.Elem().Interface(), nil
}

func Run(a app.App) error {
	//start new session in server
	respValue, err := call("start", service.StartRequest{
		App: "app1", //todo: must be registered with the server as an app
		//todo: add optional data values
	}, service.StartResponse{})
	if err != nil {
		return errors.Wrapf(err, "failed to start")
	}
	res := respValue.(service.StartResponse)
	log.Debugf("res: (%T)%+v", res, res)

	//todo: exec until need user interaction...
	//call services
	//sleep with scheduled resume
	//

	//display the app

	// ctx := context.Background()
	// d := a.(app.DisplayItem)
	// render(os.Stdout, ctx, d)
	return errors.Errorf("NYI")
} //Run()

func render(w io.Writer, ctx context.Context, d app.DisplayItem) {
	switch t := d.(type) {
	case app.DisplayButton:
		s := t.Text()(ctx)
		w.Write([]byte(fmt.Sprintf("#) %s\n", s)))
	case app.DisplayText:
		s := t.Text()(ctx)
		w.Write([]byte(fmt.Sprintf("%s\n", s)))
	case app.DisplaySection:
		s := t.Heading()(ctx)
		w.Write([]byte(fmt.Sprintf("** %s **\n", s)))
		for _, i := range t.Items() {
			render(w, ctx, i)
		}
	default:
		log.Errorf("cannot display %T", t)
	}
}
