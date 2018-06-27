package main

import (
	"zonstfe_api/common/middlewares"
	"zonstfe_api/common/my_context"
	"zonstfe_api/developer_admin/handlers"
	"zonstfe_api/common"
	"runtime"
	"net/http"
	"math/rand"
	"encoding/json"
	"zonstfe_api/corm"
	"strings"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
	"time"
	"flag"
)

var config = &my_context.Config{}
var envModel string
var configMap = map[string]map[string]interface{}{
	"develop": {
		"addr":      ":1302",
		"redisUrl":  "redis://127.0.0.1:6379",
		"pgUrl":     "postgresql://fe:qinglong@127.0.0.1:5432/fe?sslmode=disable",
		"secretKey": "bestcaicai222222",
		"roleId":    2,
	},
	"testing": {
		"addr":      ":1402",
		"redisUrl":  "redis://127.0.0.1:6379/1",
		"pgUrl":     "postgresql://fe:fe7S@QF4cLVwuLBR@127.0.0.1:5432/test_fe?sslmode=disable",
		"secretKey": "bestcaicai222222",
		"roleId":    2,
	},
	"production": {
		"addr":      ":1302",
		"redisUrl":  "redis://127.0.0.1:6379",
		"pgUrl":     "postgresql://fe:fe7S@QF4cLVwuLBR@127.0.0.1:5432/fe?sslmode=disable",
		"secretKey": "bestcaicai222222",
		"roleId":    2,
	},
}

func LoadConfiguration(envModel string) *my_context.Config {
	configMap[envModel]["envModel"] = envModel
	configJson, _ := json.Marshal(configMap[envModel])
	jsonParser := json.NewDecoder(strings.NewReader(string(configJson)))
	jsonParser.Decode(config)
	return config
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().Unix())
	flag.StringVar(&envModel, "env_model", "develop", "开发环境")
	flag.Parse()
	LoadConfiguration(envModel)
}

func main() {
	context := my_context.NewContext(config)
	defer context.Defer()
	corm.Db = context.Pgx
	handler := handlers.NewHandler(context)
	mw := middlewares.NewMiddleWares(context)
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.NotFoundHandler()
	// 初始化公共模块
	r = common.InitRouter(r, context, mw)
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(context.TokenAuth))
		r.Use(mw.Authenticator)
		r.Route("/v1", func(r chi.Router) {
			r.Route("/account", func(r chi.Router) {
				r.Get("/list", handler.Accounts)
				r.Get("/{user_id:[0-9]+}", handler.AccountOne)
				r.Put("/review/{user_id:[0-9]+}", handler.AccountReview)
				r.Get("/payment/list", handler.AccountPaymentList)
				r.Get("/payment/{payment_id:[0-9]+}", handler.AccountPaymentOne)
				r.Put("/payment/review/{payment_id:[0-9]+}", handler.AccountPaymentReview)
				r.Put("/password/update", handler.AccountPassWordUpdate)
				r.Post("/", handler.AccountOpen)
				r.Post("/email/{email}", handler.AccountEmail)
			})
			r.Route("/app", func(r chi.Router) {
				r.Get("/list", handler.Apps)
				r.Get("/{app_id:[0-9]+}", handler.AppOne)
				r.Put("/review/{app_id:[0-9]+}", handler.AppReview)
			})
			r.Route("/options", func(r chi.Router) {
				r.Get("/", handler.Options)

			})
			r.Route("/option", func(r chi.Router) {
				r.Get("/user", handler.OptionUser)
			})
			r.Route("/action/list", func(r chi.Router) {
				r.Get("/", handler.Actions)
			})
		})
	})
	context.Logger.Printf("start server at %s\n", *config.Addr)
	if envModel == "production" {
		context.Logger.Fatal(http.ListenAndServe(*config.Addr, cors.New(cors.Options{
			AllowedMethods: []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowedOrigins: []string{"fe.qinglong365.com"},
			AllowedHeaders: []string{"*"},
		}).Handler(r)))

	} else {
		context.Logger.Fatal(http.ListenAndServe(*config.Addr, cors.AllowAll().Handler(r)))
	}
}
