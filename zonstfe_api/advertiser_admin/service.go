package main

import (
	"net/http"
	"runtime"
	"time"
	"math/rand"
	"zonstfe_api/common/middlewares"
	"zonstfe_api/common/my_context"
	"zonstfe_api/common"
	"zonstfe_api/advertiser_admin/handlers"
	"zonstfe_api/corm"
	"encoding/json"
	"strings"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
	"flag"
)

var config = &my_context.Config{}

var envModel string
var configMap = map[string]map[string]interface{}{
	"develop": {
		"addr":      ":1304",
		"redisUrl":  "redis://127.0.0.1:6379",
		"pgUrl":     "postgresql://fe:qinglong@127.0.0.1:5432/fe?sslmode=disable",
		"secretKey": "bestcaicai444444",
		"roleId":    4,
	},
	"testing": {
		"addr":      ":1404",
		"redisUrl":  "redis://127.0.0.1:6379/1",
		"pgUrl":     "postgresql://fe:fe7S@QF4cLVwuLBR@127.0.0.1:5432/test_fe?sslmode=disable",
		"secretKey": "bestcaicai444444",
		"roleId":    4,
	},
	"production": {
		"addr":      ":1304",
		"redisUrl":  "redis://127.0.0.1:6379",
		"pgUrl":     "postgresql://fe:fe7S@QF4cLVwuLBR@127.0.0.1:5432/fe?sslmode=disable",
		"secretKey": "bestcaicai444444",
		"roleId":    4,
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
				r.Get("/list", handler.Accounts)
				r.Get("/{user_id:[0-9]+}", handler.AccountOne)
				r.Put("/review/{user_id:[0-9]+}", handler.AccountReview)
				r.Get("/tax/{user_id:[0-9]+}", handler.AccountTaxOne)
				r.Get("/deliver/{user_id:[0-9]+}", handler.AccountDeliverOne)
				r.Post("/", handler.AccountOpen)
				r.Post("/email/{email}", handler.AccountEmail)
				r.Put("/{user_id:[0-9]+}", handler.AccountUpdate)
				r.Put("/tax/{tax_id:[0-9]+}", handler.AccountTaxUpdate)
				r.Put("/deliver/{deliver_id:[0-9]+}", handler.AccountDeliverUpdate)
				r.Put("/password/update", handler.AccountPassWordUpdate)
				r.Post("/recharge", handler.AccountRechargeCreate)
				r.Get("/recharge/list", handler.AccountRecharges)
				r.Get("/recharge/{recharge_id:[0-9]+}", handler.AccountRechargeOne)
				r.Put("/recharge/review/{recharge_id:[0-9]+}", handler.AccountRechargesReview)

			})
			r.Route("/campaign", func(r chi.Router) {
				r.Get("/list", handler.CampaignList)
				r.Get("/ad/list", handler.AdList)
				r.Get("/ad/review/batch", handler.CampaignAdReviewBatch)
				r.Put("/ad/review/{ad_id:[0-9]+}", handler.CampaignAdReview)
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
