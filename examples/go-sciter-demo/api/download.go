package api

import (
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/Mr-ShiHuaYu/go-sciter/examples/go-sciter-demo/utils/panicutil"
	"github.com/Mr-ShiHuaYu/go-sciter/examples/go-sciter-demo/utils/pathutil"
	"github.com/Mr-ShiHuaYu/go-sciter/examples/go-sciter-demo/utils/regexputil"

	"github.com/Mr-ShiHuaYu/go-sciter"
)

var first int = 0

func reloadProgressJsFunc(url string, max int64, value int64) {
	if MainWin != nil {
		first++
		// 100次缓存更新刷新一次UI，减少性能损耗
		if first%100 == 0 || max == value {
			first = 0
			result := sciter.NewValue()
			result.Set("url", url)
			result.Set("max", max)
			result.Set("value", value)
			MainWin.Call("postCustomEvent", sciter.NewValue("reloadDownloadProgress"), result)
		}
	}
}

func BrowseSvn(args ...*sciter.Value) *sciter.Value {
	defer panicutil.TryPanic(func(err interface{}) {
		fmt.Println("[PANIC][BrowseSvn]", "catch panic:", err)
	})
	result := sciter.NewValue()
	root := args[0].String()
	fileName := args[1].String()
	username := args[2].String()
	pwd := args[3].String()
	if root == "" || fileName == "" {
		result.Set("status", "400")
		result.Set("msg", "content is null")
		return result
	}
	lines := regexputil.RegexMatchAll(fileName, `.\d{9}.zip`, false)
	date := ""
	if len(lines) == 1 {
		date = strings.TrimRight(lines[0], ".zip")
		date = date[1:7]
	}
	root = strings.TrimSuffix(root, "/")
	url := fmt.Sprintf("%s/20%s", root, date)
	indexXml, err := ListDir(url, username, pwd)
	if err != nil {
		result.Set("status", "400")
		result.Set("msg", err.Error())
		return result
	}
	// fmt.Println(root, fileName, url, indexXml.FileList, err)
	bytes, err := json.Marshal(indexXml.FileList)
	if err != nil {
		result.Set("status", "400")
		result.Set("msg", err.Error())
		return result
	}
	result.Set("status", "200")
	result.Set("rootUrl", url)
	result.Set("data", string(bytes))
	return result
}

func DownloadFile(args ...*sciter.Value) *sciter.Value {
	defer panicutil.TryPanic(func(err interface{}) {
		fmt.Println("[PANIC][DownloadFile]", "catch panic:", err)
	})
	result := sciter.NewValue()
	url := args[0].String()
	saveDir := args[1].String()
	username := args[2].String()
	pwd := args[3].String()
	if url == "" || saveDir == "" {
		result.Set("status", "400")
		result.Set("msg", "content is null")
		return result
	}
	filename := pathutil.GetFileName(url)
	fmt.Println("DownloadFile URL=", url)
	go Download(url, saveDir+filename, username, pwd)
	result.Set("status", "200")
	result.Set("msg", "")
	return result
}

type Downloader struct {
	io.Reader
	FileUrl string
	Total   int64
	Current int64
}

func (d *Downloader) Read(p []byte) (n int, err error) {
	n, err = d.Reader.Read(p)
	d.Current += int64(n)
	reloadProgressJsFunc(d.FileUrl, d.Total, d.Current)
	fmt.Printf("\r正在下载，下载进度：%.2f%%", float64(d.Current*10000/d.Total)/100)
	if d.Current == d.Total {
		fmt.Printf("\r下载完成，下载进度：%.2f%%", float64(d.Current*10000/d.Total)/100)
	}
	return
}

func Download(url string, filePath string, username string, password string) (string, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true, MinVersion: tls.VersionTLS10},
	}
	client := http.Client{
		Transport: tr,
	}
	/****
	* request 对象包含SetBasicAuth方法逻辑如下
	 */
	//auth := USERNAME + ":" + PASSWORD
	//baseAuth := base64.StdEncoding. ([]byte(auth))
	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("出现错误了", err)
		return "http download connect error", err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	pathutil.MakeDir(path.Dir(filePath))

	file, _ := os.Create(filePath)
	defer func() {
		_ = file.Close()
	}()

	downloader := &Downloader{
		FileUrl: url,
		Reader:  resp.Body,
		Total:   resp.ContentLength,
	}
	if _, err := io.Copy(file, downloader); err != nil {
		fmt.Println(err)
		return "http download error", err
	}

	return "success", nil
}

type RootXml struct {
	XMLName xml.Name `xml:"svn"`
	Index   IndexXml `xml:"index"`
}

type IndexXml struct {
	DirList  []DirFileXml `xml:"dir"`
	FileList []DirFileXml `xml:"file"`
}

type DirFileXml struct {
	Name string `xml:"name,attr"`
	Href string `xml:"href,attr"`
}

func ListDir(url string, username string, password string) (IndexXml, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true, MinVersion: tls.VersionTLS10},
	}
	client := http.Client{
		Transport: tr,
	}
	/****
	* request 对象包含SetBasicAuth方法逻辑如下
	 */
	//auth := USERNAME + ":" + PASSWORD
	//baseAuth := base64.StdEncoding.EncodeToString([]byte(auth))
	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("request error:", err)
		return IndexXml{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return IndexXml{}, err
	}

	root := RootXml{}
	err = xml.Unmarshal(body, &root)
	if err != nil {
		return IndexXml{}, err
	}
	return root.Index, nil
}
