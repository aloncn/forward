package service

import (
	"fmt"
	"github.com/idoubi/goz"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

const FengAideHost = "https://orderapi.phone580.com"

// 从url参数中提取实际要访问的地址
func getUrl(s string) string {
	return strings.Replace(s, "/?url=", "", 1)
}

// GetDo 直接请求把结果返回
func GetDo(w http.ResponseWriter, r *http.Request) {
	fullUrl := r.RequestURI
	if fullUrl == "/favicon.ico" {
		return
	}
	log.Println("request", fullUrl)
	fullUrl = getUrl(fullUrl)

	// 方法 1 直接请求把结果返回
	resp, err := goz.Get(fullUrl)
	if err != nil {
		fmt.Println("err", err)
	}
	body, _ := resp.GetBody()
	log.Println(body.String())
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

// ProxyDo 使用代理进行请求转发
func ProxyDo(w http.ResponseWriter, r *http.Request) {
	fullUrl := r.RequestURI
	if fullUrl == "/favicon.ico" {
		return
	}
	log.Println("request", fullUrl)
	fullUrl = getUrl(fullUrl)
	// 不是完整 url 需要判断内置服务地址
	if strings.HasPrefix(fullUrl, "http") == false {
		if strings.HasPrefix(fullUrl, "/fzs-open-api") {
			fullUrl = FengAideHost + fullUrl
		}
	}

	u, err := url.Parse(fullUrl)
	log.Println("urlParse", u, err)
	if err != nil {
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w, r)
}
