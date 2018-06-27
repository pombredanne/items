package handlers

import (
	"path/filepath"
	"strconv"
	"bytes"
	"github.com/gin-gonic/gin"
	"errors"
	"api-libs/rsp"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"adv-admin-api/models"
)

const UploadUrl = "http://10.105.105.217/upload/ad"

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
	Size     int64   `json:"size"`
	FileName string  `json:"file_name"`
	Msg      string  `json:"msg"`
	Status   int     `json:"status"`
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
func CreativeUpload(c *gin.Context) {
	file_info := &fileInfo{}
	if err := forward(file_info, c.Request); err != nil || file_info.Path == "" {
		c.JSON(rsp.Error("上传失败", err))
		return
	}
	c.Request.ParseMultipartForm(1000 << 20)
	width := c.Request.FormValue("width")
	height := c.Request.FormValue("height")
	file_ext := filepath.Ext(file_info.Path)
	// 验证实际尺寸和标识尺寸是否相同
	if strconv.Itoa(file_info.Width) != width || strconv.Itoa(file_info.Height) != height {
		c.JSON(rsp.Error(fmt.Sprintf("%s(%s)",
			ErrDimension, file_info.FileName)))

		return
	}
	if value, ok := extMap[file_ext]; ok {
		switch value {
		case "video":
			// 验证大小
			if file_info.Size > 30*1024000 {
				c.JSON(rsp.Error(fmt.Sprintf("%s(%s)", ErrVideoFileSize, file_info.FileName)))
				return
			}
			// 验证时长
			if file_info.Duration > 30 || file_info.Duration < 5 {
				c.JSON(rsp.Error(fmt.Sprintf("%v %s(%s)", file_info.Duration, ErrVideoDuration, file_info.FileName)))
				return
			}
		case "graphic":
			// 验证大小
			if file_info.Size > 1*1024000 {
				c.JSON(rsp.Error(fmt.Sprintf("%s(%s)", ErrGraphicFileSize, file_info.FileName)))
				return
			}
		}

	} else {
		c.JSON(rsp.Error(ErrFileType))
		return
	}
	data := map[string]interface{}{
		"width":    file_info.Width,
		"height":   file_info.Height,
		"path":     file_info.Path,
		"duration": file_info.Duration,
	}
	c.JSON(rsp.Base(data))
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

// 创意列表
func CreativeList(c *gin.Context) {
	ad_creative := &models.AdCreativeList{}
	if err := models.GetCreatives(ad_creative); err != nil {
		c.JSON(rsp.Error(rsp.ErrBadSelect, err))
		return
	}
	c.JSON(rsp.Base(ad_creative))
}
