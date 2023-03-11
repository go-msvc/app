package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-msvc/app/service"
	"github.com/go-msvc/config"
	"github.com/go-msvc/logger"
	_ "github.com/go-msvc/nats-utils"
	"github.com/go-msvc/utils/ms"
	"github.com/google/uuid"
)

var log = logger.New().WithLevel(logger.LevelDebug)

func main() {
	msvc := ms.New("workflow",
		ms.WithOper("register", operRegister))
	config.AddSource("config.json", config.File("./config.json"))
	if err := config.Load(); err != nil {
		panic(fmt.Sprintf("config: %+v", err))
	}
	msvc.Serve()
}

func operRegister(ctx context.Context, req service.StartRequest) (res *service.StartResponse, err error) {
	log.Debugf("req: (%T)%+v", req, req)
	id := uuid.New().String()
	s := session{
		id:   id,
		app:  req.App,
		data: req.Data,
	}
	sessionsMutex.Lock()
	defer sessionsMutex.Unlock()
	sessions[id] = s
	res = &service.StartResponse{
		ID:   s.id,
		App:  s.app,
		Data: s.data,
	}
	return res, nil
}

var (
	sessionsMutex sync.Mutex
	sessions      = map[string]session{}
)

type session struct {
	id   string
	app  string
	data map[string]interface{}
}

type Validator interface {
	Validate() error
}
