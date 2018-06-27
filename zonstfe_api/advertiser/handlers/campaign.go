package handlers

import (
	"zonstfe_api/common/my_context"
	"zonstfe_api/common/utils"
	"zonstfe_api/advertiser/models"
	"net/http"
	"fmt"
	"time"
	"io/ioutil"
	"os"
	"strings"
	"archive/zip"
	"path/filepath"
	"bufio"
	"gopkg.in/fatih/set.v0"
	"github.com/go-chi/chi"
	"strconv"
	"path"
	"zonstfe_api/common/utils/myfile"
	"zonstfe_api/proto"
	"golang.org/x/net/context"
)

// 广告活动列表
func (c *Handler) Campaigns(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	var req models.CampaignReq
	if err := c.Bind(r, &req); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	campaigns := &models.Campaigns{}
	var total int64
	if err := models.GetCampaigns(user["id"], &total, &req, campaigns); err != nil {
		c.JsonError(w, my_context.ErrBadSelect, err)
		return
	}
	c.JsonPage(w, len(*campaigns), total, campaigns)

}

//func validTargeting(targeting interface{}) error {
//	// 判断定向是否合法
//	sv := reflect.ValueOf(targeting)
//	if sv.Kind() == reflect.Ptr && !sv.IsNil() {
//		return validTargeting(sv.Elem().Interface())
//	}
//	//st := reflect.TypeOf(targeting)
//	size := sv.NumField()
//	for i := 0; i < size; i++ {
//		format := &models.TargetingFormat{}
//		f := sv.Field(i)
//		value := reflect.ValueOf(f.Interface())
//		if value.Kind() == reflect.Ptr && !value.IsNil() {
//			value = reflect.ValueOf(value.Elem().Interface())
//		}
//		fmt.Println(value.String(),"formatformat")
//		if err := c.BindJson([]byte(value.String()), format); err != nil {
//			return err
//		}
//
//		if *format.Open == true {
//			if len(*format.List) <= 0 {
//				return errors.New("定向数据格式错误")
//			}
//		} else {
//			*format.List = []interface{}{}
//		}
//
//	}
//	return nil
//
//}
// 查询单个广告活动
func (c *Handler) CampaignOne(w http.ResponseWriter, r *http.Request) {
	current_user := c.GetUser(r)
	campaign_id := chi.URLParam(r, "campaign_id")
	campaign := &models.Campaign{}
	if err := models.GetCampaign(campaign_id, current_user["id"], campaign); err != nil {
		c.JsonError(w, my_context.ErrBadSelect, err)
		return
	}
	c.JsonBase(w, campaign)
}

// 添加广告活动
func (c *Handler) CampaignCreate(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	campaign, targeting, url, freq := &models.Campaign{}, &models.Targeting{}, &models.Url{}, &models.CampaignFreq{}
	if err := c.Bind(r, campaign); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	if err := c.BindJson(string(*campaign.Freq), freq); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	if err := c.BindJson(string(*campaign.Targeting), targeting); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	if err := c.BindJson(string(*campaign.Url), url); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}

	//if err := validTargeting(targeting); err != nil {
	//	c.JsonError(w, content.ErrBadValid, err)
	//	return
	//}
	var campaignId int64
	if err := models.AddCampaign(user["id"], &campaignId, campaign); err != nil {
		c.JsonError(w, "创建失败", err)
		return
	}
	data := map[string]interface{}{"campaign": campaignId,}
	c.JsonBase(w, data)
	c.Context.LogAction(user["id"], user["id"], campaignId, my_context.ActionModule.Campaign, "创建", r)

}

// 开启广告活动
func (c *Handler) CampaignSwitch(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	campaignId := chi.URLParam(r, "campaign_id")
	campaign, balance := &models.Campaign{}, &models.AccountBalance{}
	if err := models.GetCampaign(campaignId, user["id"], campaign); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	if *campaign.Status == 1 {
		//开启状态则关闭
		if err := models.UpdateCampaignStatus(user["id"], campaignId, 0); err != nil {
			c.JsonError(w, my_context.ErrDefault, err)
			return
		}
	} else {
		//关闭状态则开启
		//查询账户余额
		if err := models.GetAccountBalance(user["user_id"],c.RoleID, balance); err != nil {
			c.JsonError(w, my_context.ErrDefault, err)
			return
		}
		if *balance.Balance < 100 {
			c.JsonError(w, "当前余额不足100,请充值", nil)
			return
		}
		if err := models.UpdateCampaignStatus(user["id"], campaignId, 1); err != nil {
			c.JsonError(w, my_context.ErrDefault, err)
			return
		}

	}
	c.JsonBase(w, nil)
	c.CampaignCache(campaignId, "Campaign修改缓存更新")
	c.Context.LogAction(user["id"], user["id"], campaignId, my_context.ActionModule.Campaign, "状态更改", r)
}

// 广告活动修改
func (c *Handler) CampaignUpdate(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	campaignId := chi.URLParam(r, "campaign_id")
	campaign, targeting, url := &models.Campaign{}, &models.Targeting{}, &models.Url{}
	if err := c.Bind(r, campaign); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}

	if err := c.BindJson(string(*campaign.Targeting), targeting); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	if err := c.BindJson(string(*campaign.Url), url); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	//if err := validTargeting(targeting); err != nil {
	//	c.JsonError(w, content.ErrBadValid, err)
	//	return
	//}
	if err := models.UpdateCampaign(user["id"], campaignId, campaign); err != nil {
		c.JsonError(w, my_context.ErrBadExec, err)
		return
	}
	c.JsonBase(w, nil)
	c.CampaignCache(campaignId, "Campaign修改缓存更新")
	c.Context.LogAction(user["id"], user["id"], campaignId, my_context.ActionModule.Campaign, "修改", r)
}

// 通过广告活动ID查询广告列表
func (c *Handler) CampaignAds(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	var req models.AdReq
	if err := c.Bind(r, &req); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	ads := &models.CampaignAds{}
	var total int64
	if err := models.GetAds(user["id"], &total, &req, ads); err != nil {
		c.JsonError(w, my_context.ErrBadSelect, err)
		return
	}
	c.JsonPage(w, len(*ads), total, ads)
}

// 创意列表
func (c *Handler) CreativeList(w http.ResponseWriter, r *http.Request) {
	ad_creative := &models.AdCreativeList{}
	if err := models.GetCreatives(ad_creative); err != nil {
		c.JsonError(w, my_context.ErrBadSelect, err)
		return
	}
	c.JsonBase(w, ad_creative)
}

// 广告创建
func (c *Handler) AdCreate(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	creative_id := chi.URLParam(r, "creative_id")
	adCreative := &models.AdCreative{}
	if err := models.GetCreative(creative_id, adCreative); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	var adId int64
	data := make(map[string]interface{}, 0)
	switch *adCreative.AdType {
	case "video":
		cretive_obj := &models.VideoCreative{}
		if err := c.Bind(r, cretive_obj); err != nil {
			c.JsonError(w, my_context.ErrBadFormat, err)
			return
		}
		if err := models.AddVideoAd(user["id"], &adId, adCreative, cretive_obj); err != nil {
			c.JsonError(w, my_context.ErrDefault, err)
			return
		}
		data = map[string]interface{}{
			"campaign_id": *cretive_obj.CampaignId,
			"ad_id":       adId,
		}
	case "graphic":
		cretive_obj := &models.ImageCreative{}
		if err := c.Bind(r, cretive_obj); err != nil {
			c.JsonError(w, my_context.ErrBadFormat, err)
			return
		}
		if err := models.AddImageAd(user["id"], &adId, adCreative, cretive_obj); err != nil {
			c.JsonError(w, my_context.ErrDefault, err)
			return
		}
		data = map[string]interface{}{
			"campaign_id": *cretive_obj.CampaignId,
			"ad_id":       adId,
		}

	}

	c.JsonBase(w, data)
	c.Context.LogAction(user["id"], user["id"], adId, my_context.ActionModule.Ad,
		*adCreative.Name+"创建", r)

}

// 广告修改
func (c *Handler) AdUpdate(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	ad_id := chi.URLParam(r, "ad_id")
	old_campaign_ad, campaign_ad := &models.CampaignAd{}, &models.CampaignAd{}
	if err := models.GetCampaignAd(ad_id, user["id"], campaign_ad); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	if err := c.Bind(r, campaign_ad); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	// 判断素材是否有变更
	if string(*old_campaign_ad.Creative) != string(*campaign_ad.Creative) {
		if err := models.UpdateAdWithStatus(user["id"], ad_id, 0, campaign_ad); err != nil {
			c.JsonError(w, my_context.ErrDefault, err)
			return
		}
	} else {
		if err := models.UpdateAd(user["id"], ad_id, campaign_ad); err != nil {
			c.JsonError(w, my_context.ErrDefault, err)
			return
		}
	}
	data := map[string]interface{}{
		"campaign_id": *old_campaign_ad.CampaignId,
		"ad_id":       ad_id,
	}
	c.JsonBase(w, data)
	c.CampaignCache(*old_campaign_ad.CampaignId, "Campaign AD 修改缓存更新")
	c.Context.LogAction(user["id"], user["id"], ad_id, my_context.ActionModule.Ad,
		string(*campaign_ad.Name)+"修改", r)

}

// 查询单个广告
func (c *Handler) AdOne(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	ad_id := chi.URLParam(r, "ad_id")
	campaign_ad := &models.CampaignAd{}
	if err := models.GetCampaignAd(ad_id, user["id"], campaign_ad); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	c.JsonBase(w, campaign_ad)
}

// 人群包列表
func (c *Handler) Segments(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	var req models.SegmentReq
	if err := c.Bind(r, &req); err != nil {
		c.JsonError(w, my_context.ErrBadValid, err)
	}
	segments := &models.Segments{}
	var total int64
	if err := models.GetSegments(user["id"], &total, &req, segments); err != nil {
		c.JsonError(w, my_context.ErrBadSelect, err)
		return
	}
	c.JsonPage(w, len(*segments), total, segments)

}

// 人群包创建
func (c *Handler) SegmentCreate(w http.ResponseWriter, r *http.Request) {
	user := c.GetUser(r)
	root_path := "/home/tonnn/segments/"
	segment, old_segment := &models.Segment{}, &models.Segment{}
	if err := c.Bind(r, segment); err != nil {
		c.JsonError(w, my_context.ErrBadFormat, err)
		return
	}
	filename := strconv.FormatInt(time.Now().Unix(), 10) + "." + path.Base(*segment.PkgPath)

	// 存储文件
	if err := myfile.SaveFile(*segment.PkgPath, root_path+filename); err != nil {
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}

	reader, err := zip.OpenReader(root_path + filename)
	if err != nil {
		c.JsonError(w, "提交失败", err)
		return
	}
	defer reader.Close()
	var lines, old_lines []string
	var segment_id int64
	for _, f := range reader.File {
		file_ext := filepath.Ext(f.Name)
		if !utils.StringInSlice(file_ext, ".txt") {
			continue
		}
		rc, err := f.Open()
		defer rc.Close()
		if err != nil {
			c.JsonError(w, "提交失败", err)
			return
		}
		txt_reader := bufio.NewReader(rc)
		content, err := ioutil.ReadAll(txt_reader)
		if err != nil {
			c.JsonError(w, "提交失败", err)
			return
		}
		txt_line := strings.Split(string(content), "\n")
		if len(txt_line) > 0 && txt_line[len(txt_line)-1] == "" {
			txt_line = txt_line[:len(txt_line)-1]
		}
		lines = append(lines, txt_line...)
	}
	if len(lines) <= 0 {
		c.JsonError(w, "人群包不包含txt文件或者txt文件为空", nil)
		return
	}
	for _, checkLine := range lines {
		if len(strings.Split(checkLine, "\t")) != 2 {
			c.JsonError(w, "数据格式不正确", nil)
			return
		}
	}
	uv := len(lines)
	// 判断当前 name 是否存在
	models.GetSegment(user["id"],*segment.Name,old_segment)

	if old_segment.Id != nil && *old_segment.Id != 0 {
		segment_id = int64(*old_segment.Id)
		segment_filename := root_path + fmt.Sprintf("segment:%v:%v.txt", user["id"], *old_segment.Id)
		file, err := os.OpenFile(segment_filename, os.O_CREATE|os.O_RDONLY, 0660)
		defer file.Close()
		txt_reader := bufio.NewReader(file)
		content, err := ioutil.ReadAll(txt_reader)
		if err != nil {
			c.JsonError(w, "提交失败", err)
			return
		}
		old_lines = strings.Split(string(content), "\n")
		//增量模式
		if *segment.Type == 1 {
			lines_s := set.New(utils.SliceInterface(lines)...)
			old_lines_s := set.New(utils.SliceInterface(old_lines)...)
			uv = len(set.StringSlice(set.Intersection(lines_s, old_lines_s)))
			// 加快速度 判断当前上传的和老的差集
			lines = set.StringSlice(set.Difference(lines_s, old_lines_s))
		} else {
			uv = len(lines)
		}

	}

	if err := models.AddSegment(user["id"], &segment_id, uv, segment); err != nil {
		c.JsonError(w, "提交失败", err)
		return
	}
	c.JsonBase(w, nil)
	//删除临时文件
	os.Remove(root_path + fmt.Sprintf("tmp:segment:%v:%v", user["id"], segment_id))
	os.Remove(*segment.PkgPath)
	//重新写入并覆盖静态文件
	segment_filename := root_path + fmt.Sprintf("segment:%v:%v.txt", user["id"], segment_id)
	tmp_segment_filename := root_path + fmt.Sprintf("tmp:segment:%v:%v.txt", user["id"], segment_id)

	data := map[string][]string{
		"sadd": lines,
		"srem": {},
	}
	if *segment.Type == 2 && len(old_lines) > 0 {
		data["srem"] = old_lines
	}
	client := proto.NewDataSyncClient(c.DataSync)
	_, err3 := client.SegmentCache(context.Background(), &proto.SegmentCacheRequest{
		SegmentId: segment_id,
		Sadd:      lines,
		Srem:      data["srem"],
	})
	if err3 != nil {
		c.Logger.Println(err3)
	}
	var ch chan int = make(chan int, 1)
	go func() {
		if err := ioutil.WriteFile(tmp_segment_filename, []byte(strings.Join(lines, "\n")), 0644); err != nil {
			c.Logger.Println(err)
		}
		if err := os.Rename(tmp_segment_filename, segment_filename); err != nil {
			c.Logger.Println(err)
		}
		ch <- 1
	}()
	<-ch
}
