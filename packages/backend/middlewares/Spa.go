package middlewares

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"taskism/utils"

	"github.com/gin-gonic/gin"
)

func Spa() gin.HandlerFunc {
	ignoreFunc := func(c *gin.Context) {
		c.Next()
	}
	if utils.StaticFS == nil {
		return ignoreFunc
	}

	// 读取index.html
	file, err := utils.StaticFS.Open("index.html")
	if err != nil {
		fmt.Println(err)
		return ignoreFunc
	}

	fileContentBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return ignoreFunc
	}
	fileContent := string(fileContentBytes)

	fileServer := http.FileServer(utils.StaticFS)

	return func(c *gin.Context) {
		path := c.Request.URL.Path

		// API 跳过
		if strings.HasPrefix(path, "/api") || strings.HasPrefix(path, "/custom") || strings.HasPrefix(path, "/dav") || path == "/manifest.json" {
			c.Next()
			return
		}

		// 不存在的路径和index.html均返回index.html
		if (path == "/index.html") || (path == "/") || !utils.StaticFS.Exists("/", path) {
			// 读取、替换站点设置
			/*
				options := model.GetSettingByNames("siteName", "siteKeywords", "siteScript",
					"pwa_small_icon")
				finalHTML := util.Replace(map[string]string{
					"{siteName}":       options["siteName"],
					"{siteDes}":        options["siteDes"],
					"{siteScript}":     options["siteScript"],
					"{pwa_small_icon}": options["pwa_small_icon"],
				}, fileContent)
			*/

			c.Header("Content-Type", "text/html")
			c.String(200 /*finalHTML*/, fileContent)
			c.Abort()
			return
		}

		if path == "/service-worker.js" {
			c.Header("Cache-Control", "public, no-cache")
		}

		// 存在的静态文件
		fileServer.ServeHTTP(c.Writer, c.Request)
		c.Abort()
	}
}
