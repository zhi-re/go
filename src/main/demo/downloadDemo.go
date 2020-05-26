package main

/**
下载demo
*/
import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	doc, err := goquery.NewDocument("http://www.netbian.com/desk/22676.htm")
	if err != nil {
		fmt.Println(err)
	}
	c := make(chan string)
	doc.Find("div[class=pic]").Find("a").Each(func(i int, selection *goquery.Selection) {
		//fmt.Println(selection.Html())
		img := selection.Find("img")
		bigPicturehref, _ := img.Attr("src")
		fmt.Println(bigPicturehref)
		if len(bigPicturehref) != 0 {
			go download(bigPicturehref, c)
			result := <-c
			fmt.Println(result + "下载成功")
		}

	})

}
func download(bigPicturehref string, c chan string) {
	urlBody, _ := http.Get(bigPicturehref)
	w, err := ioutil.ReadAll(urlBody.Body)
	if err != nil {
		fmt.Println("http.get err", err)
	}
	read, err := os.Create("D:/background/go/44.jpg")
	if err != nil {
		fmt.Println("创建文件失败", err)
	}

	read.Write(w)
	c <- "D:/background/go/3.jpg"

}
