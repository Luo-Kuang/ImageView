package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

type ImageItem struct {
	ImageThum  string
	ImageNomal string
}

func main() {

	fs := assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
	}

	router := gin.Default()
	router.StaticFS("/static", &fs)
	r := multitemplate.New()
	bytes, err := Asset("temp/oldindex.html") // 根据地址获取对应内容
	if err != nil {
		fmt.Println(err)
		return
	}
	t, err := template.New("index").Parse(string(bytes)) // 比如用于模板处理
	fmt.Println(t, err)
	r.Add("index", t)
	router.HTMLRender = r
	router.GET("/image", func(c *gin.Context) {
		c.HTML(200, "index", gin.H{})
	})

	var imageArr []ImageItem
	for index := 0; index < 3; index++ {
		ii := &ImageItem{"static/img/thumb-2.jpg", "static/img/1-1600.jpg"}
		imageArr = append(imageArr, *ii)
	}
	for index := 0; index < 3; index++ {
		i5 := &ImageItem{"static/img/thumb-4.jpg", "static/video/video.mp4"}
		imageArr = append(imageArr, *i5)
	}

	router.GET("/getimagelist", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": func() string {
				jsonStr, _ := json.Marshal(imageArr)
				return string(jsonStr)
			}(),
		})
	})

	http.ListenAndServe(":8888", router)
}
