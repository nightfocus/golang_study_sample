package main

// import uuid "github.com/satori/go.uuid"
import (
	"math/rand"
	"time"
)

// 使用interface对外，内部的mywidget实现是隐藏的.
// 也称为简单工厂模式
type Widget interface {
	// GetId 返回这个 Widget 的唯一标识符
	GetId() string
}

// RandString 生成随机字符串
func randString(len int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

type mywidget struct {
	id string
}

// NewWidget() 返回一个新的 Widget 实例
func NewWidget() Widget {
	return mywidget{
		// id: uuid.NewV4().String(),
		id: randString(16), // 16是长度
	}
}

func (w mywidget) GetId() string {
	return w.id
}
