// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"gopl.io/ch4/github"
	"html/template"
	"log"
	"net/http"
	"os"
)

// exercise 4.10
//func main() {
//	result, err := github.SearchIssues(os.Args[1:])
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("%d issues:\n", result.TotalCount)
//
//	for _, item := range result.Items {
//		//fmt.Println(item)
//		// exercise 4.10
//		diff := time.Now().Unix() - item.CreatedAt.Unix()
//
//		if day := diff / 3600 / 24; day <= 365 {
//			fmt.Printf("#%-5d %9.9s %.30s %.55s\n",
//				item.Number, item.User.Login, item.CreatedAt, item.Title)
//		}
//
//	}
//}

//!-

/*
//!+textoutput
$ go build gopl.io/ch4/issues
$ ./issues repo:golang/go is:open json decoder
13 issues:
#5680    eaigner encoding/json: set key converter on en/decoder
#6050  gopherbot encoding/json: provide tokenizer
#8658  gopherbot encoding/json: use bufio
#8462  kortschak encoding/json: UnmarshalText confuses json.Unmarshal
#5901        rsc encoding/json: allow override type marshaling
#9812  klauspost encoding/json: string tag not symmetric
#7872  extempora encoding/json: Encoder internally buffers full output
#9650    cespare encoding/json: Decoding gives errPhase when unmarshalin
#6716  gopherbot encoding/json: include field name in unmarshal error me
#6901  lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384    joeshaw encoding/json: encode precise floating point integers u
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#4237  gjemiller encoding/base64: URLEncoding padding is optional
//!-textoutput
*/

// exercise 4.12

//type Xkcd struct {
//	Title string // 漫画标题
//	Img   string // 图片链接
//	// Transcript string // 漫画的描述 ： 用于搜索
//}
//
//var results = make(map[string]string)
//
//func main() {
//	// 请求下载url
//	for i := 1; i < 150; i++ {
//		fmt.Printf("正在请求 %d \n", i)
//		url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
//		if result, err := Download(url); err != nil { // 下载链接
//			fmt.Println("[-]请求失败  ：", err)
//		} else {
//			fmt.Printf("下载链接 : %v \t %v\n", result.Title, result.Img)
//		}
//	}
//	// 处理输入实现检索
//	if os.Args[1:] != nil {
//		for _, arg := range os.Args[1:] {
//			if val, ok := results[arg]; ok {
//				fmt.Printf("恭喜你查询成功！\n Title: %s\t Url:%s \n", arg, val)
//			} else {
//				fmt.Println("查无此title！")
//			}
//		}
//	}
//}
//
//
//func Download(url string) (*Xkcd, error) {
//	// 请求
//	// fmt.Println("正在请求")
//	resp, err := http.Get(url)
//	if err != nil {
//		return nil, err
//	}
//	// 错误代码处理
//	if resp.StatusCode != http.StatusOK {
//		resp.Body.Close()                                 //关闭连接
//		return nil, fmt.Errorf(" query failed ： %s", err) // 错误重写
//	}
//	// 正确内容解析
//	var result Xkcd // z准备接收json内容
//	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
//		resp.Body.Close()
//		return nil, err
//	}
//	// 解析完毕，返回
//	resp.Body.Close()
//	results[result.Title] = result.Img
//	return &result, err
//}


// exercise 4.13

//type Movie struct {
//	Title string // 电影标题
//	Poster string // 电影海报url
//	//Img   string // 影片海报
//}
//
//const apikey = "87063c6d"
//
//func SearchMovie(title []string, apikey string) (*Movie, error) {
//	url := "http://www.omdbapi.com/?t="
//	//nameArray := strings.Split(title," ")
//	for i, val := range title {
//		url += val
//		if i != len(title) - 1 {
//			url += "+"
//		}
//	}
//	url += "&apikey="
//	url += apikey
//	fmt.Println(url)
//	resp, err := http.Get(url)
//	if err != nil {
//		return nil, err
//	}
//
//	// We must close resp.Body on all execution paths.
//	// (Chapter 5 presents 'defer', which makes this simpler.)
//	if resp.StatusCode != http.StatusOK {
//		resp.Body.Close()
//		return nil, fmt.Errorf("search query failed: %s", resp.Status)
//	}
//
//	var result Movie
//	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
//		resp.Body.Close()
//		return nil, err
//	}
//	resp.Body.Close()
//	return &result, nil
//}
//
//func downloadPoster(url string) {
//	resp, err := http.Get(url)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer resp.Body.Close()
//
//	bcontent, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	pos := strings.LastIndex(url, "/")
//	if pos == -1 {
//		fmt.Println("failed to parse the title of the poster")
//		return
//	}
//	f, err := os.Create(url[pos+1:])
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer f.Close()
//
//	_, err = f.Write(bcontent)
//	if err != nil {
//		fmt.Println(err)
//	}
//}
//
//func main() {
//
//	result, err := SearchMovie(os.Args[1:], apikey)
//	if err != nil {
//		log.Fatal(err)
//	}
//	//fmt.Printf("%d issues:\n", result.TotalCount)
//	fmt.Println(result)
//	downloadPoster(result.Poster)
//}
// exercise 4.14


var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func main() {
	////启动一个web服务器
	//http.HandleFunc("/", handle)
	//http.ListenAndServe("0.0.0.0:8000", nil)

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(w, result); err != nil {
		log.Fatal(err)
	}
}
