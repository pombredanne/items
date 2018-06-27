package main

import (
	"net/http"
	"math/rand"
	"time"
	"encoding/json"
	"strings"
	"zonstfe_api/data_sync/my_context"
	"zonstfe_api/data_sync/handlers"
	"runtime"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net"
	"google.golang.org/grpc"
	"zonstfe_api/proto"
	"zonstfe_api/corm"
	"flag"
)

var config = &my_context.Config{}
var envModel string
var configMap = map[string]map[string]interface{}{
	"develop": {
		"addr":        ":1310",
		"dspRedisUrl": "redis://:123456@10.0.2.204:6379/0",
		"sspRedisUrl": "redis://:123456@10.0.2.204:6379/1",
		"pgUrl":       "postgresql://fe:qinglong@111.231.137.127:5432/fe?sslmode=disable",
	},
	"testing": {
		"addr":        ":1310",
		"dspRedisUrl": "redis://:123456@10.0.2.204:6379/0",
		"sspRedisUrl": "redis://:123456@10.0.2.204:6379/1",
		"pgUrl":       "postgresql://fe:qinglong@111.231.137.127:5432/fe?sslmode=disable",
	},
	"production": {
		"addr":        ":1310",
		"dspRedisUrl": "redis://:crs-54kipecw:zenist123@10.66.212.131:6379",
		"sspRedisUrl": "redis://:crs-31iprtde:zenist123@10.66.212.208:6379",
		"pgUrl":       "postgresql://fe:fe7S@QF4cLVwuLBR@127.0.0.1:5432/fe?sslmode=disable",
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
	myContext := my_context.NewContext(config)
	handler := handlers.NewHandler(myContext)
	data_sync := handlers.NewDataSync(myContext)
	corm.Db = myContext.Pgx

	go func() {
		lis, err := net.Listen("tcp", *config.Addr+"0")
		if err != nil {
			myContext.Logger.Println(err)
			panic(err)
		}
		ss := grpc.NewServer()
		proto.RegisterDataSyncServer(ss, data_sync)
		if err := ss.Serve(lis); err != nil {
			myContext.Logger.Println(err)
			panic(err)
		}
	}()
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.NotFoundHandler()
	r.Route("/v1", func(r chi.Router) {
		r.Route("/campaign", func(r chi.Router) {
			r.Get("/cache/all", handler.CampaignCacheAll)
		})
		r.Route("/app", func(r chi.Router) {
			r.Get("/cache/all", handler.AppCacheAll)
			r.Get("/cache/identifier", handler.AppIdentifier)
		})
		r.Route("/report", func(r chi.Router) {
			r.Get("/base/import", handler.ReportBaseImport)
			r.Get("/app/import", handler.ReportAppImport)
			r.Get("/geo/import", handler.ReportGeoImport)
			r.Get("/app/reward/import", handler.ReportAppRewardImport)
			r.Get("/app/slot/import", handler.ReportAppSlotImport)
			r.Get("/dev/profit/import", handler.ReportDevProfitImport)
		})
	})
	myContext.Logger.Printf("start server at %s\n", *config.Addr)
	myContext.Logger.Fatal(http.ListenAndServe(*config.Addr, r))

}
