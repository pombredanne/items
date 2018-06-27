package middlewares

import (
	"net/http"
	"github.com/go-chi/jwtauth"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func (mw *MiddleWares) Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			switch err {
			default:
				http.Error(w, http.StatusText(401), 401)
				return
			case jwtauth.ErrExpired:
				http.Error(w, "expired", 401)
				return
			case jwtauth.ErrUnauthorized:
				http.Error(w, http.StatusText(401), 401)
				return
			case nil:
				// no error
			}
		}

		if token == nil || !token.Valid {
			http.Error(w, http.StatusText(401), 401)
			return
		}
		client := mw.Rd.Get()
		defer client.Close()
		// 验证版本号(修改密码和注销将会更新当前账户的版本号如不同则需要重新登录)
		fmt.Sprintf("login_user:%v", claims["id"])
		version, _ := redis.String(client.Do("GET", fmt.Sprintf("login_user:%v", claims["id"])))
		if version == "" || version != claims["version"].(string) || claims["version"].(string) == "" {
			http.Error(w, http.StatusText(401), 401)
			return
		}
		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}