package reptile

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type DreamExplain struct {
	Title   string
	Content []string
}

func getDom(url string) (dom *goquery.Document, err error) {
	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println(url, "请求失败！网络出现异常！", err)
		return
	}
	client := http.Client{}
	response, err := client.Do(request)
	//fmt.Println(err)
	//bytes := make([]byte, 100000)
	//response.Body.Read(bytes)
	//fmt.Println(string(bytes))
	if err != nil {
		fmt.Println(url, "响应失败！网络出现异常！", err)
		return
	}
	d, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		fmt.Println(url, "解析失败！", err)
		return
	}
	return d, nil
}

func handleResult(index int, dom *goquery.Document) (dream DreamExplain) {
	url := "https://m.zgjm.org"
	element := dom.Find(".lib_text").Find("li").Eq(index).Find("a")
	if element.Text() == "" {
		return
	}
	location, exists := element.Attr("href")
	title, _ := element.Attr("title")
	if !exists {
		fmt.Println("搜索不到资源")
		return
	}
	dream.Title = title
	document, err := getDom(url + location)
	if err != nil {
		return
	}
	document.Find(".read-content").Find("p").Each(func(i int, selection *goquery.Selection) {
		if i == 0 || i == 1 {
			return
		}
		if selection.Find("strong").Text() == "" {
			text := selection.Text()
			if textFiltered(text) {
				return
			}
			dream.Content = append(dream.Content, strings.TrimSpace(text))
		}
	})
	return
}

func ExplainDream(condition string) (content []DreamExplain, err error) {
	url := "https://m.zgjm.org/search/?wd=" + condition + "_"
	dom, err := getDom(url)
	if err != nil {
		return
	}
	for i := 0; i < 5; i++ {
		result := handleResult(i, dom)
		if reflect.DeepEqual(result, DreamExplain{}) {
			return
		}
		content = append(content, result)
	}
	return
}

func textFiltered(content string) bool {
	if strings.Contains(content, "ZGJM") {
		return true
	}
	if strings.Contains(content, "《") {
		return true
	}
	if strings.Contains(content, "：") {
		return true
	}
	return false
}
