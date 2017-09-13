# ImageView

* 集成lightgallery写的图片和视频浏览器
* 使用go-bindata打包html、css、js
* 使用ginweb服务器

![image](/WX20170913-145523.png)


bindata.go 是将静态文件打包成的*.go

index.go是程序入口

```go
导入静态文件
fs := assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
	}

	router := gin.Default()

	router.StaticFS("/static", &fs)
```

```go
这段代码将打包的html文件供gin渲染
r := multitemplate.New()
	bytes, err := Asset("temp/oldindex.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	t, err := template.New("index").Parse(string(bytes))
	fmt.Println(t, err)
	r.Add("index", t)
	router.HTMLRender = r
	router.GET("/image", func(c *gin.Context) {
		c.HTML(200, "index", gin.H{})
	})
```


[lightgallery](https://github.com/sachinchoolur/lightgallery.js)

[jteeuwen/go-bindata](https://github.com/jteeuwen/go-bindata)

[elazarl/go-bindata-assetfs](https://github.com/elazarl/go-bindata-assetfs)

[gin-gonic/gin](https://github.com/gin-gonic/gin)




