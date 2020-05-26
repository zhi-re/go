package main

/**
go爬取壁纸
*/
import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	var start int
	var end int
	fmt.Println("请输入爬取的开始页：")
	fmt.Scan(&start)
	fmt.Println("start:", start)
	fmt.Println("请输入爬取的结束页：")
	fmt.Scan(&end)
	fmt.Println("end:", end)
	for ; start <= end; start++ {
		fmt.Println("开始爬取第" + strconv.Itoa(start) + "数据")
		url := ""
		if start == 1 {
			url = "http://www.netbian.com/"
			fmt.Println(url)
		} else {
			// url := "http://www.netbian.com/index_2.htm"
			url = "http://www.netbian.com/index_" + strconv.Itoa(start) + ".htm"
			fmt.Println("不是第一页 " + url)
		}
		doc, err := goquery.NewDocument(url)
		if err != nil {
			fmt.Println(err)
		}
		var picUrls []string
		// 查找<div class="list">标签下的<li>标签
		doc.Find("div[class=list]").Find("li").Each(func(i int, selection *goquery.Selection) {
			// fmt.Println(selection.Html())
			// 查找<li>标签下的a标签
			aaa := selection.Find("a")
			// fmt.Println(aaa.Html()) 获取a标签下数据
			// 查找a标签的href 获取图片大图地址
			href, _ := aaa.Attr("href")
			//	fmt.Println(href)
			flag := strings.Contains(href, ".htm")
			if flag == true {
				picUrls = append(picUrls, href)
			}
		})
		// 遍历数组每个页面获取大图
		for e := range picUrls {
			fmt.Println(picUrls[e])
			doc, err := goquery.NewDocument("http://www.netbian.com/" + picUrls[e])
			if err != nil {
				fmt.Println(err)
			}
			// 创建管道
			c := make(chan string)
			doc.Find("div[class=pic]").Find("a").Each(func(i int, selection *goquery.Selection) {
				img := selection.Find("img")
				bigPicturehref, _ := img.Attr("src")
				fmt.Println(bigPicturehref)
				// 创建线程下载
				if len(bigPicturehref) != 0 {
					go downloadPicture(bigPicturehref, c)
					fmt.Println(<-c + "下载成功！")
				}
			})
		}
		fmt.Println("完成爬取第" + strconv.Itoa(start) + "数据")
	}
}

func downloadPicture(bigPicturehref string, c chan string) {
	urlBody, _ := http.Get(bigPicturehref)
	w, _ := ioutil.ReadAll(urlBody.Body)
	uuid, _ := uuid.NewV4()
	fileName := "D:/background/go/" + uuid.String() + ".jpg"
	read, _ := os.Create(fileName)
	read.Write(w)
	c <- fileName
}
