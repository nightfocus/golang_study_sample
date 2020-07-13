package main

/*
  对Gin库的auth进行测试
*/

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//* // sample from -> https://blog.csdn.net/fwhezfwhez/article/details/79258385
func main() {
	router := gin.Default()

	r2 := router.Group("/data")
	{
		r2.Use(Validate()) // 使用validate()中间件做身份验证
		r2.GET("uids", GetUids)
	}

	router.GET("/login", Service1) // 此方法不需要进行身份验证

	router.LoadHTMLGlob("html/*.tmpl") // 加载所有html/ 目录下的 .tmpl 模板文件
	router.Run(":18081")               // 127.0.0.1:18081
}

func GetUids(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{"message": "In GetUids"})
	c.HTML(200, "temp.tmpl", gin.H{
		"uids": "MMMMM000000000VC0000000",
	})

}

func Service1(c *gin.Context) {
	// 模拟已登录，写入cookie
	c.SetCookie("name", "ShiminLiAAAAAAAA", 3600, "/", "127.0.0.1", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "你好，登录成功，欢迎你"})
}

func Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从参数中读取，如 http://127.0.0.1:18081/data/uids?username=ft&password=123
		username := c.Query("username")
		password := c.Query("password")

		// 读取cookie
		cookie, _ := c.Cookie("name")
		if cookie == "ShiminLiAAAAAAAA" {
			// c.JSON(http.StatusOK, gin.H{"message": "cookie验证成功"})
			return
		}

		if username == "ft" && password == "123" {
			c.JSON(http.StatusOK, gin.H{"message": "身份验证成功"})
			c.Next() //该句可以省略，写出来只是表明可以进行验证下一步中间件，不写，也是内置会继续访问下一个中间件的
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "身份验证失败"})
			c.Abort()
			return // return也是可以省略的
		}
	}
}

/* //  Case 1:

// simulate some private data
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
	r := gin.Default()

	// Group using gin.BasicAuth() middleware
	// gin.Accounts is a shortcut for map[string]string
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	// /admin/secrets endpoint
	// hit "localhost:8080/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// get user, it was set by the BasicAuth middleware
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":18080")
}
*/
