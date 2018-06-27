package middlewares

import (
	"net/http"
	"zonstfe_api/common/my_context"
	"encoding/json"
	"encoding/base64"
	"github.com/garyburd/redigo/redis"
	"strings"
	"zonstfe_api/common/utils"
	"zonstfe_api/common/utils/jsonify"
	"zonstfe_api/common/models"
	"github.com/gorilla/context"
	"log"
	"runtime/debug"
	"time"
)

type MiddleWares struct {
	*my_context.Context
}

func NewMiddleWares(content *my_context.Context) *MiddleWares {
	return &MiddleWares{content}
}

func (mw *MiddleWares) CheckCookieToken(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		user := &models.LoginResult{}
		cookie, err := r.Cookie("session")
		if err != nil {
			jsonify.ErrorStatus(w, 401, "Please Login", nil)
			return
		}
		decode_str, err := base64.RawStdEncoding.DecodeString(cookie.Value)
		if err != nil {
			jsonify.ErrorStatus(w, 401, "Please Login", nil)
			return
		}
		value, err := utils.NewCBCDecrypter([]byte(decode_str), []byte(mw.SecretKey))
		if err != nil {
			jsonify.ErrorStatus(w, 401, "Please Login", nil)
			return
		}
		if !strings.Contains(string(value), "$") {
			jsonify.ErrorStatus(w, 401, "Please Login", nil)
			return
		}
		list := strings.Split(string(value), "$")
		// 获取用户信息
		client := mw.Rd.Get()
		obj, _ := redis.String(client.Do("GET", "login_user:"+list[0]))
		defer client.Close()
		if obj == "" {
			jsonify.ErrorStatus(w, 401, "Please Login", nil)
			return
		}
		// 传递用户信息
		if err := json.Unmarshal([]byte(obj), user); err != nil {
			jsonify.ErrorStatus(w, 401, "Please Login", nil)
			return
		}
		if *user.Status == 0 {
			jsonify.ErrorStatus(w, 401, "Please Login", nil)
			return
		}
		//ctx := context.WithValue(r.Context(), "current_user", user)
		//next.ServeHTTP(w, r.WithContext(ctx))
		context.Set(r, "current_user", user)
		next.ServeHTTP(w, r)
		log.Printf(
			"%s\t%s\t%s\t%d\t%s",
			r.Method,
			r.RequestURI,
			mw.GetIPAdress(r),
			*user.Id,
			time.Since(start),
		)
	}
	return http.HandlerFunc(fn)

}

func (mw *MiddleWares) RecoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {
				debug.PrintStack()
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()
		next.ServeHTTP(w, r)
		context.Clear(r)

	}
	return http.HandlerFunc(fn)
}

func (mw *MiddleWares) TestUser(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		user := &models.LoginResult{}
		user.Id = func(i int) *int { return &i }(1)
		user.AppKey = func(i string) *string { return &i }("zonst")
		user.DealType = func(i string) *string { return &i }("share")
		context.Set(r, "current_user", user)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}


