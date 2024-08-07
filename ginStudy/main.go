//package main
//
//import (
//	"net/http"
//
//	"github.com/gin-gonic/gin"
//)
//
//var db = make(map[string]string)
//
//func setupRouter() *gin.Engine {
//	// Disable Console Color
//	// gin.DisableConsoleColor()
//	r := gin.Default()
//
//	// Ping test
//	r.GET("/ping", func(c *gin.Context) {
//		c.String(http.StatusOK, "pong")
//	})
//
//	db["zhangsan"] = "zhangsanValue"
//
//	// Get user value
//	r.GET("/user/:name", func(c *gin.Context) {
//		user := c.Params.ByName("name")
//		value, ok := db[user]
//		if ok {
//			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
//		} else {
//			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
//		}
//	})
//
//	// Authorized group (uses gin.BasicAuth() middleware)
//	// Same than:
//	// authorized := r.Group("/")
//	// authorized.Use(gin.BasicAuth(gin.Credentials{
//	//	  "foo":  "bar",
//	//	  "manu": "123",
//	//}))
//	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
//		"foo":  "bar", // user:foo password:bar
//		"manu": "123", // user:manu password:123
//	}))
//
//	/* example curl for /admin with basicauth header
//	   Zm9vOmJhcg== is base64("foo:bar")
//
//		curl -X POST \
//	  	http://localhost:8080/admin \
//	  	-H 'authorization: Basic Zm9vOmJhcg==' \
//	  	-H 'content-type: application/json' \
//	  	-d '{"value":"bar"}'
//	*/
//	authorized.POST("admin", func(c *gin.Context) {
//		user := c.MustGet(gin.AuthUserKey).(string)
//
//		// Parse JSON
//		var json struct {
//			Value string `json:"value" binding:"required"`
//		}
//
//		if c.Bind(&json) == nil {
//			db[user] = json.Value
//			c.JSON(http.StatusOK, gin.H{"status": "ok"})
//		}
//	})
//
//	return r
//}
//
//const (
//	DevOpsBaseURL   = "https://devopstest.pocketcity.com"
//	DevOpsLoginPath = "/api/api-auth/"
//)
//
//func loginDevops() {
//
//}
//
//func main() {
//	r := setupRouter()
//	// Listen and Server in 0.0.0.0:8080
//	r.Run(":8080")
//}

package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type UserStruct struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func User(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "hello %s", name)
}

func loginJson(c *gin.Context) {
	var user UserStruct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.Name != "ops" || user.Password != "opsValue" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "login failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

func login(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "%s login success", name)
}
func read(c *gin.Context) {
	name := c.DefaultQuery("name", "zhangsan")
	c.String(http.StatusOK, fmt.Sprintf("read user %s", name))
}

func addUser(c *gin.Context) {
	c.String(http.StatusOK, "add user success")
}

func uploadSingleFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	// upload file
	c.SaveUploadedFile(file, "./"+file.Filename)
	c.String(http.StatusOK, fmt.Sprintf("upload file %s success", file.Filename))
}

/*
	curl -X POST http://localhost:8080/v1/uploadFiles \
	 -F "files=@./test.txt" \
	 -F "files=@./test1.txt" \
	 -H "Content-Type: multipart/form-data"
*/
func uploadFiles(c *gin.Context) {
	//if form, err := c.MultipartForm(); err != nil {
	//	c.String(http.StatusBadRequest, err.Error())
	//}
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	files := form.File["files"]
	if len(files) == 0 {
		c.String(http.StatusBadRequest, "no files found")
		return
	}

	for _, file := range files {
		fmt.Println(file.Filename)
		if err := c.SaveUploadedFile(file, "./"+file.Filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
	}
	c.String(http.StatusOK, fmt.Sprintf("upload file %d success", len(files)))
}

func main() {
	GetDevOpsToken(DevOpsBaseURL + DevOpsLoginPath)
	// 创建一个路由(带有 日志 和 恢复 中间件)
	r := gin.Default()
	// 创建一个不包含中间件的路由
	// r := gin.New()

	v1 := r.Group("/v1")
	{
		v1.GET("/login/:name", login)
		v1.POST("/loginJson", loginJson)
		v1.GET("/read", read)
		v1.POST("/addUser", addUser)
		v1.POST("/upload", uploadSingleFile)
		v1.POST("/uploadFiles", uploadFiles)
		v1.GET("/rediect", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
		})
	}

	// 创建一个Get路由绑定
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	// /user/zhangsan可以被匹配到 /user/zhangsan/不会被匹配
	r.GET("/user/:name", User)

	// 查询字符串query string, /query?strings=hello 默认值
	r.GET("/query", func(c *gin.Context) {
		strings := c.DefaultQuery("strings", "hello")
		c.String(http.StatusOK, fmt.Sprintf("string is  %s", strings))
	})

	// 启动路由
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}

}

const (
	DevOpsBaseURL   = "xxx.com"
	DevOpsLoginPath = "/api/api-auth/"
	SecretID        = "xxx"
	SecretKey       = "xxx"
)

type DevOpsLoginPar struct {
	SecretID  string `json:"secretId"`
	Timestamp string `json:"timestamp"`
	Signature string `json:"signature"`
}

func NewDevOpsLoginPar() DevOpsLoginPar {
	signatureStr := SecretID + strconv.FormatInt(time.Now().UnixMilli(), 10)

	h := hmac.New(md5.New, []byte(SecretKey))
	h.Write([]byte(signatureStr))
	signature := hex.EncodeToString(h.Sum(nil))

	return DevOpsLoginPar{
		SecretID:  SecretID,
		Timestamp: strconv.FormatInt(time.Now().UnixMilli(), 10),
		Signature: signature,
	}

}

func GetDevOpsToken(url string) {
	devLoginJson, err := json.Marshal(NewDevOpsLoginPar())
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	// 创建请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(devLoginJson))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	// 发起请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// 打印响应
	fmt.Printf("Response status: %s\n", resp.Status)
	fmt.Printf("Response body: %s\n", string(body))
}
