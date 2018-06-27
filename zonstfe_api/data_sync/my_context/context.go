package my_context

import (
	"github.com/jmoiron/sqlx"
	"zonstfe_api/common/mydb"
	"github.com/garyburd/redigo/redis"
	"net/http"
	"log"
	"os"
	"fmt"
	"runtime"
	"github.com/json-iterator/go"
	"github.com/go-playground/form"
	"gopkg.in/go-playground/validator.v9"
	"zonstfe_api/common/utils/myemail"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary
var decoder = form.NewDecoder()
var validate = validator.New()

type Response struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	Count  int         `json:"count,omitempty"`
	Total  int         `json:"total,omitempty"`
}

type Context struct {
	Pgx    *sqlx.DB
	SspRd  *redis.Pool
	DspRd  *redis.Pool
	Logger *log.Logger
	EnvModel     string
}

type Config struct {
	Addr        *string `json:"addr"`
	DspRedisUrl *string `json:"dspRedisUrl"`
	SspRedisUrl *string `json:"sspRedisUrl"`
	PgUrl       *string `json:"pgUrl"`
	EnvModel  *string `json:"envModel"`
}

func NewContext(config *Config) *Context {
	return &Context{
		Pgx:    mydb.GetPgx(*config.PgUrl),
		DspRd:  mydb.GetRedis(*config.DspRedisUrl),
		SspRd:  mydb.GetRedis(*config.SspRedisUrl),
		EnvModel:     *config.EnvModel,
		Logger: log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (c *Context) LogEventEnd(event_id, err string, status int, end_time int64) {
	//发送邮件
	if status != 1 {
		myemail.SendEmail([]string{"1020300659@qq.com"}, []string{"1020300659@qq.com"}, "报表导入出错", err)
	}
	if _, err := c.Pgx.Exec(`update log_event set status=$1,error_msg=$2,end_time=$3 where event_id=$4`,
		status, err, end_time, event_id); err != nil {
		c.Logger.Println(err)
	}
}

func (c *Context) LogEventStart(name, event_id, event_obj string, start_time int64) {
	if _, err := c.Pgx.Exec(`insert into log_event(name,event_id,event_obj,start_time) values($1,$2,$3,$4)`,
		name, event_id, event_obj, start_time); err != nil {
		c.Logger.Println(err)
	}
}

func (c *Context) JsonBase(w http.ResponseWriter, data interface{}) {
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
		c.Logger.Println(err)
	}
}

func (c *Context) JsonErrorStatus(w http.ResponseWriter, status int, msg interface{}, error error) {
	w.Header().Set("CONTENT-TYPE", "application/json")
	w.WriteHeader(status)
	if error != nil {
		c.Logger.Println(error)
	}
	r := &Response{}
	r.Status = status

	r.Data = make([]map[string]string, 0)
	r.Msg = fmt.Sprintf("%v", msg)
	if b, err := json.Marshal(r); err == nil {
		w.Write(b)
	} else {
		c.Logger.Println(err)
	}

}

func (c *Context) JsonError(w http.ResponseWriter, msg interface{}, error error) {
	w.Header().Set("CONTENT-TYPE", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	if error != nil {
		pc, fn, line, _ := runtime.Caller(1)
		c.Logger.Printf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, error)
	}
	r := &Response{}
	r.Status = -1
	r.Data = make([]map[string]string, 0)
	r.Msg = fmt.Sprintf("%v", msg)
	if b, err := json.Marshal(r); err == nil {
		w.Write(b)
	} else {
		c.Logger.Println(err)
	}
	return

}
func (c *Context) JsonPage(w http.ResponseWriter, count, total int, data interface{}) {
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
		c.Logger.Println(err)
	}

}

// Json赋值 并验证
func (c *Context) BindJson(r *http.Request, obj interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(obj)
	if err != nil {
		return err
	}
	if err := validate.Struct(obj); err != nil {
		return err
	}
	return nil
}

func (c *Context) BindStringJson(data string, obj interface{}) error {
	if err := json.Unmarshal([]byte(data), obj); err != nil {
		return err
	}
	if err := validate.Struct(obj); err != nil {
		return err
	}
	return nil
}

// query参数绑定 不做验证
func (c *Context) BindQuery(r *http.Request, obj interface{}) error {
	err := decoder.Decode(obj, r.URL.Query())
	if err != nil {
		return err
	}
	return nil
}
