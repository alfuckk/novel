package reutrns

import "encoding/json"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (resp *Response) WithMsg(msg string) *Response {
	return &Response{
		Code: resp.Code,
		Msg:  resp.Msg,
		Data: resp.Data,
	}
}

func (resp *Response) WithData(data interface{}) *Response {
	return &Response{
		Code: resp.Code,
		Msg:  resp.Msg,
		Data: data,
	}
}

func response(code int, msg string) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

func (resp *Response) ToString() string {
	err := &struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{
		Code: resp.Code,
		Msg:  resp.Msg,
		Data: resp.Data,
	}
	raw, _ := json.Marshal(err)
	return string(raw)
}
