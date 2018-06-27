package main

import (
	"net/http"
	"crypto/md5"
	"io"
	"encoding/hex"
	"os"
	"bytes"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"fmt"
	"io/ioutil"
	"os/exec"
	"encoding/json"
	"path/filepath"
	"github.com/rs/cors"
	"runtime"
	"log"

	"strconv"
)

type Response struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	Count  int         `json:"count,omitempty"`
	Total  int         `json:"total,omitempty"`
}

type fileInfo struct {
	Path     string `json:"path"`
	Duration string `json:"duration"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
	Size     int    `json:"size"`
	FileName string `json:"file_name"`
}

// 通用静态文件上传 返回
const (
	domain     = "https://static.qinglong365.com/"
	adPath     = "/home/tonnn/ad/"
	publicPath = "/home/tonnn/public/"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Post("/upload", simpleUpload)
	r.Post("/ad/upload", adUpload)
	r.Post("/public/upload", publicUpload)
	http.ListenAndServe(":20001", cors.Default().Handler(r))

}

type Sizer interface {
	Size() int64
}

func simpleUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		Error(w, err, err)
		return
	}

	file_ext := filepath.Ext(handler.Filename)
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		Error(w, err, err)
		return
	}
	// 生成Md5 文件名
	file_name := hashByteMd5(buf.Bytes()) + file_ext
	// 先保存文件
	f, err := os.OpenFile(publicPath+file_name, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		Error(w, err, err)
		return
	}
	defer f.Close()
	_, err2 := io.Copy(f, buf)
	if err2 != nil {
		Error(w, err2, err2)
		return
	}
	data_map := make(map[string]interface{}, 0)
	data_map["path"] = domain + "p/" + file_name
	data_map["file_name"] = handler.Filename
	data_map["size"] = file.(Sizer).Size()
	if b, err := json.Marshal(data_map); err == nil {
		w.Write(b)
	} else {
		log.Println(err)
	}

}

// 公共上传
func publicUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		Error(w, err, err)
		return
	}

	file_ext := filepath.Ext(handler.Filename)
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		Error(w, err, err)
		return
	}
	// 生成Md5 文件名
	file_name := hashByteMd5(buf.Bytes()) + file_ext
	// 先保存文件
	f, err := os.OpenFile(publicPath+file_name, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		Error(w, err, err)
		return
	}
	defer f.Close()
	_, err2 := io.Copy(f, buf)
	if err2 != nil {
		Error(w, err2, err2)
		return
	}
	cmd := exec.Command("ffprobe", "-i", publicPath+file_name, "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams", "-hide_banner")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		//fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		Error(w, err, err)
		return
	}

	video_info := make(map[string]interface{}, 0)
	if err := json.Unmarshal([]byte(out.String()), &video_info); err != nil {
		Error(w, err, err)
		return
	}
	streams, data_map := video_info["streams"].([]interface{})[0].(map[string]interface{}), make(map[string]interface{}, 0)
	format := video_info["format"].(map[string]interface{})
	if streams["width"] == nil {
		streams["width"] = 0
	}
	if streams["height"] == nil {
		streams["height"] = 0
	}
	if format["duration"] == nil {
		format["duration"] = "0"
	}

	size, err := strconv.Atoi(format["size"].(string))
	if err != nil {
		Error(w, err, err)
		return
	}
	duration, err := strconv.ParseFloat(format["duration"].(string), 64)
	if err != nil {
		Error(w, err, err)
		return
	}

	data_map["path"] = domain + "p/" + file_name
	data_map["size"] = size
	data_map["duration"] = duration
	data_map["width"] = streams["width"]
	data_map["height"] = streams["height"]
	data_map["file_name"] = handler.Filename
	if b, err := json.Marshal(data_map); err == nil {
		w.Write(b)
	} else {
		log.Println(err)
	}
}

// ad 上传
func adUpload(w http.ResponseWriter, r *http.Request) () {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		Error(w, err, err)
		return
	}
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		Error(w, err, err)
		return
	}

	file_ext := filepath.Ext(handler.Filename)
	// 生成Md5 文件名
	file_name := hashByteMd5(buf.Bytes()) + file_ext
	// 先保存文件
	f, err := os.OpenFile(adPath+file_name, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		Error(w, err, err)
		return
	}
	defer f.Close()
	_, err2 := io.Copy(f, buf)
	if err2 != nil {
		Error(w, err2, err2)
		return
	}
	cmd := exec.Command("ffprobe", "-i", adPath+file_name, "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams", "-hide_banner")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		Error(w, err, err)
		return
	}
	video_info := make(map[string]interface{}, 0)
	if err := json.Unmarshal([]byte(out.String()), &video_info); err != nil {
		Error(w, err, err)
		return
	}
	fmt.Println(video_info)
	streams, data_map := video_info["streams"].([]interface{})[0].(map[string]interface{}), make(map[string]interface{}, 0)
	format := video_info["format"].(map[string]interface{})
	if streams["width"] == nil {
		streams["width"] = 0
	}
	if streams["height"] == nil {
		streams["height"] = 0
	}
	if format["duration"] == nil {
		format["duration"] = "0"
	}
	size, err := strconv.Atoi(format["size"].(string))
	if err != nil {
		Error(w, err, err)
		return
	}
	duration, err := strconv.ParseFloat(format["duration"].(string), 64)
	if err != nil {
		Error(w, err, err)
		return
	}
	data_map["path"] = domain + file_name
	data_map["size"] = size
	data_map["duration"] = duration
	data_map["width"] = streams["width"]
	data_map["height"] = streams["height"]
	data_map["file_name"] = handler.Filename
	if b, err := json.Marshal(data_map); err == nil {
		w.Write(b)
	} else {
		log.Println(err)
	}

}

// 请求转发
func forwardHandler(w http.ResponseWriter, req *http.Request) {
	httpClient := &http.Client{}
	// we need to buffer the body if we want to read it here and send it
	// in the request.
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// you can reassign the body if you need to parse it as multipart
	req.Body = ioutil.NopCloser(bytes.NewReader(body))

	//// create a new url from the raw RequestURI sent by the client
	//url := fmt.Sprintf("%s://%s%s", proxyScheme, proxyHost, req.RequestURI)

	proxyReq, err := http.NewRequest(req.Method, "http://localhost:20001/public/upload", bytes.NewReader(body))

	// We may want to filter some headers, otherwise we could just use a shallow copy
	// proxyReq.Header = req.Header
	proxyReq.Header = make(http.Header)
	for h, val := range req.Header {
		proxyReq.Header[h] = val
	}

	resp, err := httpClient.Do(proxyReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	body, err1 := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err1.Error(), http.StatusBadGateway)
		return
	}
	w.Header().Set("CONTENT-TYPE", "application/json")
	w.Write(body)

	// legacy code
}

func hashByteMd5(file []byte) string {
	hash := md5.New()
	hash.Write(file)
	return hex.EncodeToString(hash.Sum(nil))
}

func hashFileMd5(filePath string) (string, error) {
	var returnMD5String string
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil
}

func Base(w http.ResponseWriter, data interface{}) {
	w.Header().Set("CONTENT-TYPE", "application/json")
	if data == nil {
		data = make([]map[string]string, 0)
	}
	r := &Response{}
	r.Status = 0
	r.Msg = "success"
	r.Data = data
	if b, err := json.Marshal(r); err == nil {
		fmt.Println(string(b))
		w.Write(b)
	} else {
		log.Println(err)
	}
}
func Error(w http.ResponseWriter, msg interface{}, error error) {
	w.Header().Set("CONTENT-TYPE", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	if error != nil {
		pc, fn, line, _ := runtime.Caller(1)
		log.Printf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, error)
	}
	r := &Response{}
	r.Status = -1
	r.Data = make([]map[string]string, 0)
	r.Msg = fmt.Sprintf("%v", msg)
	if b, err := json.Marshal(r); err == nil {
		w.Write(b)
	} else {
		log.Println(err)
	}

}
