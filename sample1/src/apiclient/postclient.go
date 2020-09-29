package apiclient

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

// 需要先注册megvii账号获得
var (
	meg_api_key    string = "a"
	meg_api_secret string = "a"
)

// 一个sample，调用以下服务的示例：
// https://console.faceplusplus.com.cn/documents/4888373
func PostMegFacepp() {
	postData := make(map[string]string)
	postData["api_key"] = meg_api_key
	postData["api_secret"] = meg_api_secret

	file, err := os.Open("D:\\21.jpg")
	if err != nil {
		// 本地文件打不开，用网络文件测试
		postData["image_url"] = "http://00.minipic.eastday.com/20170221/20170221212912_cbff414ccd6113e1d49401b874e438c6_9.jpeg"
	}
	defer file.Close()

	fileContent, err := ioutil.ReadAll(file)
	postData["image_base64"] = base64.StdEncoding.EncodeToString(fileContent)

	url := "https://api-cn.faceplusplus.com/facepp/v3/detect"

	// 如果图片中没有人脸，会输出：
	// {... "face_num":0}
	PostWithFormData("POST", url, &postData)
}

func PostWithFormData(method, url string, postData *map[string]string) {
	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	for k, v := range *postData {
		w.WriteField(k, v)
	}
	w.Close()

	// 方式一：通过Do方法选择POST方式调用
	// req, _ := http.NewRequest(method, url, body)
	// req.Close = true // 完成req后就关闭http连接
	// req.Header.Set("Content-Type", w.FormDataContentType())
	// resp, _ := http.DefaultClient.Do(req)

	// 方式二：直接用Post方法调用
	resp, err := http.DefaultClient.Post(url, w.FormDataContentType(), body)
	if err == nil && resp.StatusCode == 200 {
		data, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		fmt.Println(resp.StatusCode)
		fmt.Printf("%s", data)
	} else {
		fmt.Println("post err: ", err, resp)
	}
	// 这是关闭空闲的http连接
	// 否则底层tcp连接不会立即释放
	http.DefaultClient.CloseIdleConnections()

}
