package myGin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
 * 一个对Gin最简单的封装，为了封装而封装。
 */

type MyGin struct {
	Ge *gin.Engine
}

type MyHandlerFunc func(c *MyContext)

// 对Gin的注册HandlerFunc过程进行封装，仍然返回一个 gin.HandlerFunc函数对象
// 该函数对象负责执行 h 函数.
func Handle(h MyHandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &MyContext{
			c,
		}
		h(ctx)
	}
}

type MyContext struct {
	*gin.Context
}

func NewMyGin() *MyGin {
	return &MyGin{
		Ge: gin.Default(),
	}
}

func (m *MyGin) Run() error {
	return m.Ge.Run()
}

func (c *MyContext) SuccessResponse(data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"retcode": 200,
		"msg":     "success",
		"data":    data,
	})
}

// 一个handle sample，可以不放在package myGin中
// 这样封装一下的好处是，当返回值需要统一格式化时，如 SuccessReponse()，可统一处理。
func HandleHello(c *MyContext) {
	// Todo.

	c.SuccessResponse(gin.H{
		"name":  "Cruise",
		"hobby": "Sport",
	})
}
