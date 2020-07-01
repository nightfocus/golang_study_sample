package common

import (
	"bufio"
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	LOGDebug   *log.Logger // 记录所有日志
	LOGInfo    *log.Logger // 一般的信息
	LOGWarning *log.Logger // 需要注意的信息
	LOGError   *log.Logger // 严重的问题
)

/*
type MyLog struct {
	LOG2 *log.Logger
}

func (ml MyLog) WriteLog(format string, slst ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	format2 := "%s:%d " + format
	slst2 := make([]interface{}, 2)
	slst2[0] = file
	slst2[1] = line
	slst2 = append(slst2, slst)
	ml.LOG2.Printf(format2, slst2...)
}
*/

func initSrvLog() {
	var iomw io.Writer

	logFile, err := os.OpenFile("meg_mst.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		DbgPrint("Failed to open error log file: %s", err)
		iomw = io.MultiWriter(os.Stdout)
		return
	} else {
		// 日志输出：标准输出 + 文件输出
		iomw = io.MultiWriter(logFile, os.Stdout)
	}

	// 日志自动前缀时间戳，文件名，行号，级别信息
	logFlags := log.Ldate | log.Lmicroseconds | log.Lshortfile | log.Lmsgprefix

	LOGDebug = log.New(iomw, "[Debug] ", logFlags)
	LOGInfo = log.New(iomw, "[INFO] ", logFlags)
	LOGWarning = log.New(iomw, "[WARNING] ", logFlags)
	LOGError = log.New(iomw, "[ERROR] ", logFlags)
}

func init() {
	initSrvLog()
}

// 一个简单的printf输出函数的封装，会输出时间和文件行号信息。
func DbgPrint(format string, slst ...interface{}) {

	// 获得调用该DbgPrint() 的文件全名和行号
	if _, srcFile, srcLine, ok := runtime.Caller(1); ok {
		// 只保留文件名
		idx := strings.LastIndex(srcFile, "/")
		if idx != -1 {
			idx++
			srcFile = srcFile[idx:]
		}

		format2 := "%s %s:%d " + format // 在前面先输出 时间 文件名:行号
		// 先分配3个长度，存放三个预置值，再用append()追加调用者参数传入
		// 将max置为10，在append()时，只有当slst个数超过7个后，才会重新分配内存。
		slst2 := make([]interface{}, 3, 10)
		slst2[0] = time.Now().Format("2006-01-02 15:04:05")
		slst2[1] = srcFile
		slst2[2] = srcLine
		slst2 = append(slst2, slst...)

		fmt.Printf(format2, slst2...)
	} else {
		fmt.Printf(format, slst...)
	}
}

// 获取协程的id
func GetGoroutineID() uint64 {
	b := make([]byte, 64)
	runtime.Stack(b, false)
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

// 计算HMac-SHA1签名
func GetHmacCode(s string, key string) string {
	h := hmac.New(sha1.New, []byte(key))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

/*
func getSha1Code(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}
*/

type Ids struct {
	DidKey string
	Uid    string
	UidKey string
}

/*  从自定义的文件中加载数据
    文件格式为：
#P2P Server
47.95.253.138;112.126.83.206

#DID,DIDKey,UID,UIDKey
ABCDE0001,KEYABCDE0001,MEGOT00AA1LKYKS,abCDefGH1234567890
ABCDE0002,KEYABCDE0002,MEGOT00AA2UPETC,abCDefGH1234567890
ABCDE0003,KEYABCDE0003,MEGOT00AA3UPETC,abCDefGH1234567890

	返回
  		string: p2p server addr
		map: key 是 DID, value 是Ids struct
*/
func LoadIdsDb(fileName string) (string, map[string]Ids, error) {

	p2pServerAddr := ""
	idsMap := make(map[string]Ids)

	if file, err := os.Open(fileName); err != nil {
		return p2pServerAddr, idsMap, err
	} else {
		defer file.Close()

		firstLine := true
		/*
			bufio.NewScanner()的参数是io.Reader类型，只有一个接口:
			Read(p []byte) (n int, err error)
			os.File 类型实现了上述的Read接口，所以能直接传递给NewScanner()
		*/
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			tl := scanner.Text()
			// # 开头为注释行
			if len(tl) == 0 || tl[0] == '#' {
				continue
			}

			if firstLine {
				p2pServerAddr = tl
				firstLine = false
			} else {
				s := strings.Split(tl, ",")
				if len(s) == 4 {
					idsMap[s[0]] = Ids{DidKey: s[1], Uid: s[2], UidKey: s[3]}
				}
			}
		}

		return p2pServerAddr, idsMap, nil
	}
}
