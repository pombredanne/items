package jsonify

import (
	"net/http"
	"fmt"
	"zonstfe_api/common/my_context"
	"zonstfe_api/common/utils/validator"
	"encoding/json"
	//"runtime"
	"runtime"
)

var Context *my_context.Context

type Response struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	Count  int       `json:"count,omitempty"`
	Total  int       `json:"total,omitempty"`
}

func Base(w http.ResponseWriter, data interface{}) {
	w.Header().Set("CONTENT-TYPE", "application/json")
	if data == nil {
		data = make([]map[string]string, 0)
	}
	r := &Response{}
	r.Status = 0
	r.Msg = "success"
	r.Data = data
	if b, err := json.Marshal(r); err == nil {
		w.Write(b)
	} else {
		Context.Logger.Println(err)
	}
}

func ErrorStatus(w http.ResponseWriter, status int, msg interface{}, error error) {
	w.Header().Set("CONTENT-TYPE", "application/json")
	w.WriteHeader(status)
	if error != nil {
		Context.Logger.Println(error)
	}
	r := &Response{}
	r.Status = status

	r.Data = make([]map[string]string, 0)
	r.Msg = fmt.Sprintf("%v", msg)
	if b, err := json.Marshal(r); err == nil {
		w.Write(b)
	} else {
		Context.Logger.Println(err)
	}

}

func Error(w http.ResponseWriter, msg interface{}, error error) {
	w.Header().Set("CONTENT-TYPE", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	if error != nil {
		pc, fn, line, _ := runtime.Caller(1)
		Context.Logger.Printf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, error)
	}
	r := &Response{}
	r.Status = -1
	r.Data = make([]map[string]string, 0)
	r.Msg = fmt.Sprintf("%v", msg)
	if b, err := json.Marshal(r); err == nil {
		w.Write(b)
	} else {
		Context.Logger.Println(err)
	}

}
func Page(w http.ResponseWriter, count, total int, data interface{}) {
	w.Header().Set("CONTENT-TYPE", "application/json")
	r := &Response{}
	r.Status = 0
	r.Msg = "success"
	r.Count = count
	r.Total = total
	r.Data = data
	if b, err := json.Marshal(r); err == nil {
		w.Write(b)
	} else {
		Context.Logger.Println(err)
	}

}

func ReadJsonObject(r *http.Request, obj interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(obj)
	if err != nil {
		return err
	}
	return nil
}

// Json赋值 并验证
func ReadJsonObjectValid(r *http.Request, obj interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(obj)
	if err != nil {
		return err
	}
	if err := validator.Validate(obj); err != nil {
		return err
	}
	return nil
}

func UnmarshalValid(data []byte, obj interface{}) error {
	if err := json.Unmarshal(data, obj); err != nil {
		return err
	}
	if err := validator.Validate(obj); err != nil {
		return err
	}
	return nil
}


