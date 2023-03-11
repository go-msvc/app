package service

import "github.com/go-msvc/errors"

type StartRequest struct {
	App  string                 `json:"app" doc:"Identifies the application"`
	Data map[string]interface{} `json:"data" doc:"Data to start the session with (optional)"`
}

func (req StartRequest) Validate() error {
	if req.App == "" {
		return errors.Errorf("missing app")
	}
	return nil
}

type StartResponse struct {
	ID   string                 `json:"id"`
	App  string                 `json:"app" doc:"Identifies the application (echoed from req)"`
	Data map[string]interface{} `json:"data" doc:"Data to start the session with (optional, echoed from req)"`
}
