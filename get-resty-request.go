package core

import "github.com/go-resty/resty/v2"

func GetRestyRequest() *resty.Request {
	req := resty.New().R()
	req.SetHeader("User-Agent", "maschine/1.0")
	req.SetHeader("Content-Type", "application/json")
	req.SetHeader("Accept", "application/json")
	return req
}
