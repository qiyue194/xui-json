package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Config 创建json信息结构体
type Config struct {
	XuiServer []string `json:"xui_server"`
	SSL       struct {
		Key  string `json:"key"`
		Cert string `json:"cert"`
	} `json:"ssl"`
}

func xuiHttp(w http.ResponseWriter, r *http.Request) {
	//获取token参数
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "token不存在", http.StatusBadRequest)
		return
	}
	//获取json信息
	configData, err := ioutil.ReadFile("./config.json")
	if err != nil {
		http.Error(w, "读取配置文件失败", http.StatusInternalServerError)
		return
	}
	//解析json信息
	var config Config
	err = json.Unmarshal(configData, &config)
	if err != nil {
		http.Error(w, "解析配置文件失败", http.StatusInternalServerError)
		return
	}
	var xuiServerJsonData string
	//遍历xui_server
	for _, xuiServer := range config.XuiServer {
		//发起http请求
		resp, err := http.Get(xuiServer + token)
		if err != nil {
			//http.Error(w, "请求失败", http.StatusInternalServerError)
			//return
			continue
		}
		defer resp.Body.Close()
		//读取响应
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "读取响应失败", http.StatusInternalServerError)
			return
		}
		//进行base64解码
		decodeData, err := base64.StdEncoding.DecodeString(string(data))
		if err != nil {
			http.Error(w, "解码失败", http.StatusInternalServerError)
			return
		}
		//字符串拼接
		xuiServerJsonData += string(decodeData)
	}
	//fmt.Println(xuiServerJsonData)
	//将字符串金额base64编码返回
	w.Write([]byte(base64.StdEncoding.EncodeToString([]byte(xuiServerJsonData))))
}

// 获取公钥证书路径地址
func getSsLCert() string {
	//获取json信息
	configData, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println("读取配置文件失败:", err)
	}
	//解析json信息
	var config Config
	err = json.Unmarshal(configData, &config)
	if err != nil {
		fmt.Println("解析配置文件失败:", err)
	}
	return config.SSL.Cert
}

// 获取私钥证书路径地址
func getSslKey() string {
	//获取json信息
	configData, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println("读取配置文件失败:", err)
	}
	//解析json信息
	var config Config
	err = json.Unmarshal(configData, &config)
	if err != nil {
		fmt.Println("解析配置文件失败:", err)
	}
	return config.SSL.Key
}
func main() {
	http.HandleFunc("/xuijson", xuiHttp)
	//启动http服务器

	sslKey := getSslKey()
	sslCert := getSsLCert()
	if sslKey != "" && sslCert != "" {
		fmt.Println("服务器启动，监听端口: 30303 -- 启动ssl")
		err := http.ListenAndServeTLS(":30303", sslCert, sslKey, nil)
		if err != nil {
			fmt.Println("服务器启动失败:", err)
		}
	} else {
		fmt.Println("服务器启动，监听端口: 30303")
		err := http.ListenAndServe(":30303", nil)
		if err != nil {
			fmt.Println("服务器启动失败:", err)
		}
	}
}
