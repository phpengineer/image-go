package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"os"
	"net/http"
	"log"
	"encoding/json"
	"strings"
)

type Size interface {
	Size() int64
}

type FileInfo interface {
	FileInfo() (os.FileInfo, error)
}

type Handler struct {

}

func NewHandler() *Handler{
	return &Handler{}
}

//静态文件服务
func (handle *Handler) Index(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	fs := http.FileServer(http.Dir("view"))
	fs.ServeHTTP(w, req)
}

func (handle *Handler) Upload(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	formField := conf.GetString("upload.form_field")
	allowTypeSlice := conf.GetString("upload.allow_type")
	rootDir := conf.GetString("upload.root_dir")
	filenameLen := conf.GetInt("upload.filename_len")
	dirNameLen := conf.GetInt("upload.dirname_len")
	maxSize := conf.GetInt("upload.max_size")
	thumnails := conf.GetString("upload.thumbnails")
	server := conf.GetString("listen.server")

	imageUri := "http://" + server + "/images/"

	req.ParseMultipartForm(5 * 1024)
	//获取表单文件上传数据
	file, fileHeader, err := req.FormFile(formField)

	if err != nil {
		log.Println("upload field error" , err.Error())
		handle.jsonError(w,"Upload field error", nil)
		return
	}

	defer file.Close()

	filename := fileHeader.Filename
	ext := filename[strings.LastIndex(filename, "."):]
	isAllow := false
	for _, allowType := range allowTypeSlice {
		if strings.ToLower(allowType) == strings.ToLower(ext) {
			isAllow = true
			break
		}
	}

	if isAllow == false {
		log.Println("Forbidden file format: ",  ext)
		handle.jsonError(w, "Forbidden file format!", nil)
		return
	}

	if size ,ok := file.(Size); ok {
		
	}



}


func (handle *Handler) jsonError(w http.ResponseWriter, message string, data interface{}) {
	handle.jsonMessage(w, 0, message, data)
}

func (handle *Handler) jsonMessage(w http.ResponseWriter, code int, message, data interface{}) {
	type Result struct {
		Code    int         `json:"code"`
		Message interface{} `json:"message"`
		Data    interface{} `json:"data"`
	}
	result := Result {
		Code:    code,
		Message: message,
		Data:    data,
	}
	resultByte, err := json.Marshal(result)
	if err != nil {
		result.Code = 0
		result.Message = err.Error()
		resultByte, _ = json.Marshal(result)
		fmt.Fprint(w, string(resultByte))
		return
	}
	fmt.Fprint(w, string(resultByte))
}
