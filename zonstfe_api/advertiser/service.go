package main

import (
	"net/http"
	"runtime"
	"time"
	"math/rand"
	"encoding/json"
	"zonstfe_api/common/middlewares"
	"zonstfe_api/common/my_context"
	"zonstfe_api/common"
	"zonstfe_api/advertiser/handlers"
	"zonstfe_api/corm"
	"github.com/rs/cors"
	"strings"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"flag"
)

var config = &my_context.Config{}
var envModel string
var configMap = map[string]map[string]interface{}{
	"develop": {
		"addr":      ":1303",
		"redisUrl":  "redis://127.0.0.1:6379",
		"pgUrl":     "postgresql://fe:qinglong@127.0.0.1:5432/fe?sslmode=disable",
		"secretKey": "bestcaicai333333",
		"roleId":    3,
	},
	"testing": {
		"addr":      ":1403",
		"redisUrl":  "redis://127.0.0.1:6379/1",
		"pgUrl":     "postgresql://fe:fe7S@QF4cLVwuLBR@127.0.0.1:5432/test_fe?sslmode=disable",
		"secretKey": "bestcaicai333333",
		"roleId":    3,
	},
	"production": {
		"addr":      ":1303",
		"redisUrl":  "redis://127.0.0.1:6379",
		"pgUrl":     "postgresql://fe:fe7S@QF4cLVwuLBR@127.0.0.1:5432/fe?sslmode=disable",
		"secretKey": "bestcaicai333333",
		"roleId":    3,
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
				r.Put("/", handler.AccountUpdate)
				r.Put("/password", handler.AccountPassWordUpdate)
				r.Get("/recharge/list", handler.AccountRecharges)
				r.Post("/recharge", handler.AccountRechargesCreate)
				r.Get("/balance", handler.AccountBalance)
			})
			r.Route("/campaign", func(r chi.Router) {
				r.Post("/", handler.CampaignCreate)
				r.Get("/list", handler.Campaigns)
				r.Put("/{campaign_id:[0-9]+}", handler.CampaignUpdate)
				r.Put("/switch/{campaign_id:[0-9]+}", handler.CampaignSwitch)
				r.Get("/{campaign_id:[0-9]+}", handler.CampaignOne)
				r.Get("/ads", handler.CampaignAds)
				r.Post("/ad/{creative_id:[0-9]+}", handler.AdCreate)
				r.Put("/ad/{ad_id:[0-9]+}", handler.AdUpdate)
				r.Get("/ad/{ad_id:[0-9]+}", handler.AdOne)
				r.Post("/creative/upload", handler.CreativeUpload)
				r.Get("/segment/list", handler.Segments)
				r.Get("/creative/list", handler.CreativeList)
				r.Post("/segment", handler.SegmentCreate)

			})
			r.Route("/report", func(r chi.Router) {
				r.Get("/base", handler.ReportBase)
				r.Get("/base/hour", handler.ReportBaseHour)
				r.Get("/geo", handler.ReportGeo)
				//r.Get("/geo/country", handler.ReportGeoCountry)
				//r.Get("/geo/province", handler.ReportGeoProvince)
				//r.Get("/geo/city", handler.ReportGeoCity)
				r.Get("/app", handler.ReportApp)

			})
			r.Route("/options", func(r chi.Router) {
				r.Get("/", handler.Options)
			})
			r.Route("/option", func(r chi.Router) {
				r.Get("/campaign", handler.OptionCampaign)
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
