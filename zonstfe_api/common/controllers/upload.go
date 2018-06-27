package controllers

import (
	"fmt"
	"encoding/json"
	"io"
	"os"
	"net/http"
	"zonstfe_api/common/my_context"
	"mime/multipart"
	"zonstfe_api/common/utils/jsonify"
	"bytes"
	"io/ioutil"
)

const UploadUrl = "http://upload.qinglong365.com/upload"

type FileController struct {
	*my_context.Context
}

type fileInfo struct {
	Path     string  `json:"path,omitempty"`
	Duration float64 `json:"duration,omitempty"`
	Height   int     `json:"height,omitempty"`
	Width    int     `json:"width,omitempty"`
	Size     int     `json:"size,omitempty"`
	FileName string  `json:"file_name,omitempty"`
}

func NewFileController(content *my_context.Context) *FileController {
	return &FileController{content}
}

func (c *FileController) FileUpload(w http.ResponseWriter, r *http.Request) {
	file_info := &fileInfo{}
	if err := forward(file_info, r); err != nil {
		c.JsonError(w, "上传失败", err)
		return
	}
	jsonify.Base(w, file_info)

}

func forward(res *fileInfo, req *http.Request) error {
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
	if err := json.Unmarshal(body, res); err != nil {
		return err
	}
	return nil
}

// 模拟客户端上传
func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("upload_file", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	//打开文件句柄操作
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}
