package main

import (
	"flag"
	"corm"
	"github.com/gin-gonic/gin"
	"adv-admin-api/config"
	"adv-admin-api/handlers"
	"runtime"
	"api-libs/middleware/token"
	"api-libs/middleware"
	"api-libs/rsp"
	"log"
	"adv-admin-api/client"
	"github.com/rs/cors"
	"net/http"
	_ "api-libs/validation"
)

var (
	envModel string
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.StringVar(&envModel, "env_model", "develop", "开发环境")
	flag.Parse()
	// 配置加载
	config.Init(envModel)
	// 设置 response envModel
	rsp.SetEnvModel(envModel)
	// db
	config.Conf.DB.NewDb()
	corm.Db = config.Conf.DB.GetDb()
	// redis
	config.Conf.Redis.NewRedis()
	//
	client.InitGrpcClient("localhost:13100")
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	defer config.Conf.DB.CloseDb()
	r.POST("/login", handlers.Login)
	v1 := r.Group("/v1")
	v1.Use(token.JwtAuth(config.Conf.SigningKey, "user"), middleware.LoginVersion(config.Conf.Redis.GetPool()))
	{
		v1.POST("/logout", handlers.Logout)
		v1.POST("/upload", handlers.Upload)
		account := v1.Group("/account")
		{
			// 账户
			account.GET("/list", handlers.Accounts)
			account.GET("/one/:userId", handlers.AccountOne)
			account.PUT("/review/:userId", handlers.AccountReview)
			account.POST("/open", handlers.AccountOpen)
			account.POST("/email/:email", handlers.AccountEmail)
			account.PUT("/edit/:userId", handlers.AccountUpdate)
			account.PUT("/password", handlers.AccountPassWordUpdate)
			//account.PUT("/tax/:taxId", handlers.AccountTaxUpdate)
			// 税务
			account.PUT("/tax/:userId", handlers.AccountTaxUpdate)
			account.GET("/tax/:userId", handlers.GetAccountTax)
			// 财务
			account.PUT("/finance/:userId", handlers.AccountFinanceUpdate)
			account.GET("/finance/:userId", handlers.GetFinance)
			// 联系人
			account.PUT("/deliver/:userId", handlers.AccountDeliverUpdate)
			account.GET("/deliver/:userId", handlers.GetAccountDeliver)
			// 资质
			account.PUT("/qualification/:userId", handlers.QualificationUpdate)
			account.GET("/qualification/:userId", handlers.GetQualification)
			// 行业
			account.POST("/industry", handlers.IndustryCreate)
			account.GET("/industry/list", handlers.GetIndustry)
			account.PUT("/industry/:industryId", handlers.IndustryUpdate)
			// 充值
			account.POST("/recharge", handlers.AccountRechargeCreate)
			account.GET("/recharge/list", handlers.AccountRecharges)
			account.GET("/recharge/one/:rechargeId", handlers.AccountRechargeOne)
			account.GET("/recharge/cost/list", handlers.AccountRechargeCost)
			account.PUT("/recharge/review/:rechargeId", handlers.AccountRechargesReview)
			// 合同
			account.POST("/contract", handlers.AccountContractCreate)
			account.POST("/contract/list", handlers.GetAccountContract)
			account.GET("/cost/list", handlers.GetAccountCost)
		}
		campaign := v1.Group("/campaign")
		{
			campaign.POST("/add", handlers.CampaignCreate)
			campaign.GET("/list", handlers.CampaignList)
			campaign.PUT("edit/:campaignId", handlers.CampaignUpdate)
			campaign.PUT("/switch/:campaignId", handlers.CampaignSwitch)
			campaign.GET("/one/:campaignId", handlers.CampaignOne)
		}

		ad := v1.Group("/ad")
		{
			ad.GET("/list", handlers.AdList)
			ad.POST("add/:creativeId", handlers.AdCreate)
			ad.PUT("edit/:adId", handlers.AdUpdate)
			ad.PUT("/switch/:adId", handlers.AdSwitch)
			ad.PUT("/review_batch", handlers.CampaignAdReviewBatch)
			ad.PUT("/review/:adId", handlers.CampaignAdReview)
			ad.GET("/one/:adId", handlers.AdOne)
			ad.GET("/cost/list", handlers.AdCost)
			ad.GET("/costDetail", handlers.AdCostDetail)
			ad.POST("/costImport", handlers.AdCostImport)
		}

		creative := v1.Group("/creative")
		{
			creative.GET("/list", handlers.CreativeList)
			creative.POST("/upload", handlers.CreativeUpload)
		}

		segment := v1.Group("/segment")
		{
			segment.GET("/list", handlers.Segments)
		}

		report := v1.Group("/report")
		{
			report.GET("/base", handlers.ReportBase)
			report.GET("/base/hour", handlers.ReportBaseHour)
			report.GET("/geo", handlers.ReportGeo)
			report.GET("/app", handlers.ReportApp)
			report.GET("/app/export", handlers.ReportAppExport)
			report.GET("/geo/export", handlers.ReportGeoExport)
			report.GET("/base/export", handlers.ReportBaseExport)
			report.GET("/base/hour/export", handlers.ReportBaseHourExport)
		}

		option := v1.Group("/option")
		{
			option.GET("/user", handlers.GetOptionUser)
			option.GET("/campaign", handlers.GetOptionCampaign)
			option.GET("/segment", handlers.GetOptionSegment)
			option.GET("/list", handlers.Options)
		}

		action := v1.Group("/action")
		{
			action.GET("/list", handlers.Actions)
		}

		admin := v1.Group("/admin")
		{
			admin.POST("/signup", handlers.AdminSignup)
			admin.PUT("/password", handlers.AdminPassWordUpdate)
		}

	}
	if envModel == "production" {
		log.Fatal(http.ListenAndServe(config.Conf.Addr, cors.New(cors.Options{
			AllowedMethods: []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowedOrigins: []string{"http://admin.qinglong365.com"},
			AllowedHeaders: []string{"*"},
		}).Handler(r)))

	} else {
		log.Fatal(http.ListenAndServe(config.Conf.Addr, cors.AllowAll().Handler(r)))
	}
}
