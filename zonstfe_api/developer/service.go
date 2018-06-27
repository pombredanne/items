package main

import (
	"zonstfe_api/common/my_context"
	"zonstfe_api/common/middlewares"
	"zonstfe_api/corm"
	"zonstfe_api/common"
	"runtime"
	"math/rand"
	"time"
	"net/http"
	"encoding/json"
	"strings"
	"github.com/go-chi/chi/middleware"
	"zonstfe_api/developer/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/rs/cors"
	"flag"
)

var config = &my_context.Config{}
var envModel string

var configMap = map[string]map[string]interface{}{
	"develop": {
		"addr":      ":1301",
		"redisUrl":  "redis://127.0.0.1:6379",
		"pgUrl":     "postgresql://fe:qinglong@127.0.0.1:5432/fe?sslmode=disable",
		"secretKey": "bestcaicai111111",
		"roleId":    1,
	},
	"testing": {
		"addr":      ":1401",
		"redisUrl":  "redis://127.0.0.1:6379/1",
		"pgUrl":     "postgresql://fe:fe7S@QF4cLVwuLBR@127.0.0.1:5432/test_fe?sslmode=disable",
		"secretKey": "bestcaicai111111",
		"roleId":    1,
	},
	"production": {
		"addr":      ":1301",
		"redisUrl":  "redis://127.0.0.1:6379",
		"pgUrl":     "postgresql://fe:fe7S@QF4cLVwuLBR@127.0.0.1:5432/fe?sslmode=disable",
		"secretKey": "bestcaicai111111",
		"roleId":    1,
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
	r.NotFoundHandler()
	// 初始化公共模块
	r = common.InitRouter(r, context, mw)
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(context.TokenAuth))
		r.Use(mw.Authenticator)
		r.Route("/v1", func(r chi.Router) {
			r.Route("/account", func(r chi.Router) {
				r.Get("/", handler.Account)
				r.Put("/", handler.AccountEdit)
				r.Get("/balance", handler.AccountBalance)
				r.Put("/password", handler.AccountPassWordEdit)
				r.Get("/finance", handler.Finance)
				r.Put("/finance/{financeId:[0-9]+}", handler.FinanceEdit)
				r.Post("/payment", handler.PaymentCreate)
				r.Get("/payment/list", handler.Payments)
			})
			r.Route("/app", func(r chi.Router) {
				r.Get("/list", handler.Apps)
				r.Post("/", handler.AppCreate)
				r.Put("/{appId:[0-9]+}", handler.AppEdit)
				r.Get("/{appId:[0-9]+}", handler.AppOne)
			})
			r.Route("/report", func(r chi.Router) {
				r.Get("/app/reward", handler.ReportAppReward)
				r.Get("/app/slot", handler.ReportAppSlot)
			})
			r.Route("/options", func(r chi.Router) {
				r.Get("/", handler.Options)

			})
			r.Route("/option", func(r chi.Router) {
				r.Get("/app", handler.Options)

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
