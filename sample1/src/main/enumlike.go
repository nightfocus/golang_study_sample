package main

import (
	"fmt"
	"sort"
)

type Platform int

const (
	Ubuntu Platform = 1 + iota
	CentOS
	Windows
	iOS
	Android
	Dos
)

var PlatformMap = map[Platform]string{
	Ubuntu:  "Linux-Ubuntu",
	CentOS:  "Linux-CentOS",
	Windows: "Microsoft-Win",
	iOS:     "Apple-iOS",
	Android: "Google-Android",
	Dos:     "Microsoft-Dos",
}

// 将int类型的枚举，转换成字符串输出
func (p Platform) ShowText() string {
	return PlatformMap[p]
}

// 可以用上面定义map的方式，扩展任何对枚举的输出

// KeyMap 输出模型
// 标签在这个示例中没有作用
type KeyMap struct {
	Key   string `json:"k"`
	Val   int    `json:"v"`
	NoUse int    `json:"nouse"`
}

type IntKeyMap struct {
	Key int `json:"k"`
	Val int `json:"v"`
}

// KeyMapSlice keymap sort
/*
自定义类型需要排序，需要实现以下三个interface
func Len() int {… }
func Swap(i, j int) {… }
func Less(i, j int) bool {… }
*/
type KeyMapSlice []KeyMap

func (s KeyMapSlice) Len() int {
	return len(s)
}
func (s KeyMapSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s KeyMapSlice) Less(i, j int) bool {
	return s[i].Val < s[j].Val
}

// 将上面的PlatformMap 转成struct切片 列表输出
func (m Platform) List() []KeyMap {
	km := make([]KeyMap, 0)
	for k, v := range PlatformMap {
		km = append(km, KeyMap{Key: fmt.Sprintf("%v", v), Val: int(k)})
	}
	sort.Sort(KeyMapSlice(km))
	return km
}
