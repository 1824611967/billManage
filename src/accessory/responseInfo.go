package accessory

import (
	"encoding/json"
	"net/http"
)

type Response struct{
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}
func NewResponse() *Response{
	return &Response{
		Code: "default",
		Msg : "default",
	}
}

func (r *Response) Answer(w http.ResponseWriter) error{
	if r.Code == "default"{
		return nil
	}else {
		jsonStr, err := json.Marshal(r)
		w.Write(jsonStr)

		return err
	}
}