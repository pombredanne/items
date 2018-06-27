package my_context

import (
	"github.com/garyburd/redigo/redis"
	"github.com/jmoiron/sqlx"
	"github.com/json-iterator/go"
	"net/http"
	"strings"
	"fmt"
	"strconv"
	"net"
	"bytes"
	"log"
	"os"
	"errors"
	"zonstfe_api/common/mydb"
	"runtime"
	"github.com/go-playground/form"
	"gopkg.in/go-playground/validator.v9"
	"zonstfe_api/common/options"
	"reflect"
	"github.com/go-chi/jwtauth"
	"github.com/dlclark/regexp2"
	"time"
	"github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"zonstfe_api/proto"
	"golang.org/x/net/context"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	DeveloperRoleId        = 1
	DeveloperManageRoleId  = 2
	AdvertiserRoleId       = 3
	AdvertiserManageRoleId = 4
)

var (
	ErrDefault   = errors.New("服务错误")
	ErrBadFormat = errors.New("数据格式错误")
	ErrBadValid  = errors.New("数据验证错误")
	ErrBadSelect = errors.New("数据查询错误 请联系管理员")
	ErrBadExec   = errors.New("当前操作执行失败 请联系管理员")
	ErrObject    = errors.New("操作对象不存在")
	ActionModule = actionModule{
		Account:     "账户",
		Campaign:    "推广活动",
		App:         "媒体",
		Ad:          "广告",
		TarGet:      "定向",
		CreativeSet: "创意",
		Finance:     "财务",
		Segment:     "人群包",
	}
)

type actionModule struct {
	Account     string
	Campaign    string
	Ad          string
	TarGet      string
	CreativeSet string
	Finance     string
	Segment     string
	App         string
}

type Context struct {
	Pgx          *sqlx.DB
	Rd           *redis.Pool
	SecretKey    string
	AppSecretKey string
	RoleID       int
	Logger       *log.Logger
	TokenAuth    *jwtauth.JWTAuth
	DataSync     *grpc.ClientConn
	EnvModel     string
}

type Config struct {
	Addr      *string `json:"addr"`
	RedisUrl  *string `json:"redisUrl"`
	PgUrl     *string `json:"pgUrl"`
	SecretKey *string `json:"secretKey"`
	RoleId    *int    `json:"roleId"`
	EnvModel  *string `json:"envModel"`
}

func NewContext(config *Config) *Context {
	return &Context{
		Pgx:          mydb.GetPgx(*config.PgUrl),
		Rd:           mydb.GetRedis(*config.RedisUrl),
		SecretKey:    *config.SecretKey,
		AppSecretKey: "bestcaicai999999",
		RoleID:       *config.RoleId,
		EnvModel:     *config.EnvModel,
		Logger:       log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile),
		TokenAuth:    jwtauth.New("HS256", []byte(*config.SecretKey), nil),
		DataSync:     GetDataSyncConn("localhost:13100"),
	}
}

func GetDataSyncConn(url string) *grpc.ClientConn {
	data_sync, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return data_sync
}

func (c *Context) Defer() {
	c.Pgx.Close()
	c.Rd.Close()
	c.DataSync.Close()
}

var decoder = form.NewDecoder()
var validate = validator.New()

// 判断正则
func myRegexp(fl validator.FieldLevel) bool {
	// 为空不做验证
	if fl.Field().String() == "" {
		return true
	}
	if val, ok := options.RegexpMap[fl.Param()]; ok {
		re := regexp2.MustCompile(val, 0)
		if isMatch, _ := re.MatchString(fl.Field().String()); isMatch {
			return true
		}
		//if m, err := regexp.MatchString(val, fl.Field().String()); m {
		//	fmt.Println(err,"----")
		//	return true
		//}
	}
	return false
}

// 判断是否包含在 map 中
func myOption(fl validator.FieldLevel) bool {
	if _, ok := options.All[fl.Param()][fl.Field().String()]; ok {
		return true
	}
	return false
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Func, reflect.Map, reflect.Slice:
		return v.IsNil()
	case reflect.Array:
		z := true
		for i := 0; i < v.Len(); i++ {
			z = z && isZero(v.Index(i))
		}
		return z
	case reflect.Struct:
		z := true
		for i := 0; i < v.NumField(); i++ {
			z = z && isZero(v.Field(i))
		}
		return z
	}
	// Compare other types directly:
	z := reflect.Zero(v.Type())
	return v.Interface() == z.Interface()
}

func init() {
	validate.RegisterValidation("regexp", myRegexp)
	validate.RegisterValidation("option", myOption)
}

func (c *Context) LogAction(user_id, action_user_id, action_module_id interface{}, action_module, action_type string, r *http.Request) {
	//sqls := strings.Join(sql_list, ";")
	query := `insert into log_actions (user_id,action_user_id,action_sql,request_path,
	request_method,platform_id,ip_address,
	action_module,action_type,action_id) values ($1, $2, $3,$4,$5,$6,$7,$8,$9,$10,$11)`
	args := []interface{}{user_id, action_user_id, "", r.RequestURI, r.Method, c.RoleID, c.GetIPAdress(r),
		action_module, action_type, action_module_id}
	if _, err := c.Pgx.Exec(query, args...); err != nil {
		c.Logger.Println(err)
	}
}

func (c *Context) ParseSql(format string, p []interface{}) string {
	args, i := make([]string, len(p)*2), 0
	for k, v := range p {
		args[i] = "$" + strconv.Itoa(k+1)
		args[i+1] = fmt.Sprint(v)
		i += 2
	}
	return strings.NewReplacer(args...).Replace(format)

}

func (c *Context) LogEventEnd(event_id, err string, status int, end_time int64) {
	if _, err := c.Pgx.Exec(`update log_event set status=$1,error_msg=$2,end_time=$3 where event_id=$4`,
		status, err, end_time, event_id); err != nil {
		c.Logger.Println(err)
	}
}

func (c *Context) AccountMsg(user_id interface{}, user_role int, title, content, group_name string) error {
	if _, err := c.Pgx.Exec(`insert into account_message(user_id,user_role,
		title,content,group_name) values($1,$2,$3,$4,$5)`, user_id, user_role, title, content, group_name); err != nil {
		return err
	}
	return nil
}

func (c *Context) LogEventStart(event_obj interface{}, name, request_url string) {
	start_time := time.Now().Unix()
	event_id := uuid.NewV4().String()
	if _, err := c.Pgx.Exec(`insert into log_event(name,event_id,event_obj,start_time) values($1,$2,$3,$4)`,
		name, event_id, event_obj, start_time); err != nil {
		c.Logger.Println(err)
	}
	//
	resp, err := http.Get(fmt.Sprintf("%s/%s/%s", request_url, event_obj, event_id))
	if err != nil {
		c.Logger.Println(err)
	}
	defer resp.Body.Close()

}

func (c *Context) CampaignCache(campaign_id interface{}, name string) {
	if c.EnvModel == "production" {
		start_time := time.Now().Unix()
		event_id := uuid.NewV4().String()
		if _, err := c.Pgx.Exec(`insert into log_event(name,event_id,event_obj,start_time) values($1,$2,$3,$4)`,
			name, event_id, campaign_id, start_time); err != nil {
			c.Logger.Println(err)
		}
		client := proto.NewDataSyncClient(c.DataSync)
		_, err := client.CampaignCache(context.Background(), &proto.CampaignCacheRequest{
			CampaignId: fmt.Sprintf("%v", campaign_id),
			EventId:    event_id,
		})

		if err != nil {
			c.Logger.Println(err)
		}
	}

}

func (c *Context) AppCache(app_id interface{}, name string) {
	if c.EnvModel == "production" {
		start_time := time.Now().Unix()
		event_id := uuid.NewV4().String()
		if _, err := c.Pgx.Exec(`insert into log_event(name,event_id,event_obj,start_time) values($1,$2,$3,$4)`,
			name, event_id, app_id, start_time); err != nil {
			c.Logger.Println(err)
		}
		client := proto.NewDataSyncClient(c.DataSync)
		_, err := client.AppCache(context.Background(), &proto.AppCacheRequest{
			AppId:   fmt.Sprintf("%v", app_id),
			EventId: event_id,
		})
		if err != nil {
			c.Logger.Println(err)
		}
	}

}

func isPrivateSubnet(ipAddress net.IP) bool {
	// my use case is only concerned with ipv4 atm
	if ipCheck := ipAddress.To4(); ipCheck != nil {
		// iterate over all our ranges
		for _, r := range privateRanges {
			// check if this ip is in a private range
			if inRange(r, ipAddress) {
				return true
			}
		}
	}
	return false
}

type ipRange struct {
	start net.IP
	end   net.IP
}

// inRange - check to see if a given ip address is within a range given
func inRange(r ipRange, ipAddress net.IP) bool {
	// strcmp type byte comparison
	if bytes.Compare(ipAddress, r.start) >= 0 && bytes.Compare(ipAddress, r.end) < 0 {
		return true
	}
	return false
}

var privateRanges = []ipRange{
	ipRange{
		start: net.ParseIP("10.0.0.0"),
		end:   net.ParseIP("10.255.255.255"),
	},
	ipRange{
		start: net.ParseIP("100.64.0.0"),
		end:   net.ParseIP("100.127.255.255"),
	},
	ipRange{
		start: net.ParseIP("172.16.0.0"),
		end:   net.ParseIP("172.31.255.255"),
	},
	ipRange{
		start: net.ParseIP("192.0.0.0"),
		end:   net.ParseIP("192.0.0.255"),
	},
	ipRange{
		start: net.ParseIP("192.168.0.0"),
		end:   net.ParseIP("192.168.255.255"),
	},
	ipRange{
		start: net.ParseIP("198.18.0.0"),
		end:   net.ParseIP("198.19.255.255"),
	},
}

func (c *Context) GetIPAdress(r *http.Request) string {
	for _, h := range []string{"X-Forwarded-For", "X-Real-Ip"} {
		addresses := strings.Split(r.Header.Get(h), ",")
		// march from right to left until we get a public address
		// that will be the address right before our proxy.
		for i := len(addresses) - 1; i >= 0; i-- {
			ip := strings.TrimSpace(addresses[i])
			// header can contain spaces too, strip those out.
			realIP := net.ParseIP(ip)
			if !realIP.IsGlobalUnicast() || isPrivateSubnet(realIP) {
				// bad address, go to next
				continue
			}
			return ip
		}
	}
	return "127.0.0.1"
}

func (c *Context) GetUser(r *http.Request) map[string]interface{} {
	_, claims, _ := jwtauth.FromContext(r.Context())
	return claims
}

type Response struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	Sum    interface{} `json:"sum,omitempty"`
	Count  int         `json:"count,omitempty"`
	Total  int64       `json:"total,omitempty"`
}

func (c *Context) JsonBase(w http.ResponseWriter, data interface{}) {
	w.Header().Set("CONTENT-TYPE", "application/json; charset=utf-8")
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
func (c *Context) JsonSumPage(w http.ResponseWriter, count int, total int64, data interface{}, sum interface{}) {
	w.Header().Set("CONTENT-TYPE", "application/json; charset=utf-8")
	r := &Response{}
	r.Status = 0
	r.Msg = "success"
	r.Count = count
	r.Total = total
	r.Data = data
	r.Sum = sum
	if b, err := json.Marshal(r); err == nil {
		w.Write(b)
	} else {
		c.Logger.Println(err)
	}

}
func (c *Context) JsonSum(w http.ResponseWriter, data interface{}, sum interface{}) {
	w.Header().Set("CONTENT-TYPE", "application/json; charset=utf-8")
	if data == nil {
		data = make([]map[string]string, 0)
	}
	r := &Response{}
	r.Status = 0
	r.Msg = "success"
	r.Data = data
	r.Sum = sum
	if b, err := json.Marshal(r); err == nil {
		w.Write(b)
	} else {
		c.Logger.Println(err)
	}
}

func (c *Context) JsonErrorStatus(w http.ResponseWriter, status int, msg interface{}, error error) {
	w.Header().Set("CONTENT-TYPE", "application/json; charset=utf-8")
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
	w.Header().Set("CONTENT-TYPE", "application/json; charset=utf-8")
	r := &Response{}
	w.WriteHeader(http.StatusBadRequest)
	if error != nil {
		pc, fn, line, _ := runtime.Caller(1)
		c.Logger.Printf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, error)
		r.Msg = fmt.Sprintf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, error)
	}
	if c.EnvModel == "production" {
		r.Msg = fmt.Sprintf("%v", msg)
	} else if c.EnvModel != "production" && error == nil {
		r.Msg = fmt.Sprintf("%v", msg)
	}
	r.Status = -1
	r.Data = make([]map[string]string, 0)
	if b, err := json.Marshal(r); err == nil {
		w.Write(b)
	} else {
		c.Logger.Println(err)
	}
	return

}
func (c *Context) JsonPage(w http.ResponseWriter, count int, total int64, data interface{}) {
	w.Header().Set("CONTENT-TYPE", "application/json; charset=utf-8")
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

func (c *Context) Bind(r *http.Request, obj interface{}) error {
	if r.Method == "GET" {
		err := decoder.Decode(obj, r.URL.Query())
		if err != nil {
			return err
		}
	} else {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(obj)
		if err != nil {
			return err
		}
		if err := validate.Struct(obj); err != nil {
			return err
		}
	}
	return nil
}

func (c *Context) BindJson(data string, obj interface{}) error {
	if err := json.Unmarshal([]byte(data), obj); err != nil {
		return err
	}
	if err := validate.Struct(obj); err != nil {
		return err
	}
	return nil
}
