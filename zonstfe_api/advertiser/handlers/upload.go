package handlers

import (
	"fmt"
	"path/filepath"
	"net/http"
	"errors"
	_ "image/jpeg"
	_ "image/png"
	_ "image/gif"
	"zonstfe_api/common/utils/jsonify"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strconv"
)

const UploadUrl = "http://upload.qinglong365.com/ad/upload"

var (
	ErrFileType        = errors.New("文件格式不支持")
	ErrGraphicFileSize = errors.New("图像文件太大 必须小于1mb")
	ErrVideoFileSize   = errors.New("视频文件太大 必须小于30mb")
	ErrVideoDuration   = errors.New("视频时长太长")
	ErrBad             = errors.New("上传失败")
	ErrDimension       = errors.New("文件尺寸不支持")
)

type fileInfo struct {
	Path     string  `json:"path"`
	Duration float64 `json:"duration"`
	Height   int     `json:"height"`
	Width    int     `json:"width"`
	Size     int     `json:"size"`
	FileName string  `json:"file_name"`
}

var extMap = map[string]string{
	".mp4":  "video",
	".mov":  "video",
	".avi":  "video",
	".gif":  "graphic",
	".png":  "graphic",
	".jpg":  "graphic",
	".jpeg": "graphic",
}

// 广告上传
func (c *Handler) CreativeUpload(w http.ResponseWriter, r *http.Request) {
	file_info := &fileInfo{}
	if err := forward(file_info, r); err != nil || file_info.Path == "" {
		c.JsonError(w, "上传失败", err)
		return
	}
	//campaign_ad := &models.CampaignAd{}
	//if err := models.GetCampaignAd(ad_id, campaign_ad); err != nil {
	//	c.JsonError(w, my_context.ErrDefault, err)
	//	return
	//}
	r.ParseMultipartForm(1000 << 20)
	width := r.FormValue("width")
	height := r.FormValue("height")
	file_ext := filepath.Ext(file_info.Path)
	// 验证实际尺寸和标识尺寸是否相同
	if strconv.Itoa(file_info.Width) != width || strconv.Itoa(file_info.Height) != height {
		c.JsonError(w, fmt.Sprintf("%s(%s)",
			ErrDimension, file_info.FileName), nil)
		return
	}
	if value, ok := extMap[file_ext]; ok {
		switch value {
		case "video":
			// 验证大小
			if file_info.Size > 30*1024000 {
				c.JsonError(w, fmt.Sprintf("%s(%s)", ErrVideoFileSize, file_info.FileName), nil)
				return
			}
			// 验证时长
			if file_info.Duration > 30 || file_info.Duration < 5 {
				c.JsonError(w, fmt.Sprintf("%v %s(%s)", file_info.Duration, ErrVideoDuration, file_info.FileName), nil)
				return
			}
		case "graphic":
			// 验证大小
			if file_info.Size > 1*1024000 {
				c.JsonError(w, fmt.Sprintf("%s(%s)", ErrGraphicFileSize, file_info.FileName), nil)
				return
			}
		}

	} else {
		c.JsonError(w, ErrFileType, nil)
		return
	}
	data := map[string]interface{}{
		"width":    file_info.Width,
		"height":   file_info.Height,
		"path":     file_info.Path,
		"duration": file_info.Duration,
	}
	jsonify.Base(w, data)

}

// 请求转发
func forward(file_info *fileInfo, req *http.Request) error {
	httpClient := &http.Client{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	req.Body = ioutil.NopCloser(bytes.NewReader(body))
	proxyReq, err := http.NewRequest(req.Method, UploadUrl, bytes.NewReader(body))
	proxyReq.Header = make(http.Header)
	for h, val := range req.Header {
		proxyReq.Header[h] = val
	}
	resp, err := httpClient.Do(proxyReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		return err1
	}
	if err := json.Unmarshal(body, file_info); err != nil {
		return err
	}
	return nil
}
