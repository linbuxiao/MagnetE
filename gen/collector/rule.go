package collector

import (
	"fmt"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/go-resty/resty/v2"
)

var spider *spiderRepo

type spiderRepo struct {
	client *resty.Client
}

type itemCollectorResult struct {
	Title  string
	Magnet string
	Size   string
}

type listCollectorResult struct {
	Title string
	Data  []*itemCollectorResult
}

func init() {
	client := getDefaultClient()
	spider = &spiderRepo{
		client: client,
	}
}

func getDefaultClient() *resty.Client {
	c := resty.New()
	c.SetHeader("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36")
	return c
}
func (s *spiderRepo) btsow(name string) (*listCollectorResult, error) {
	url := strings.ReplaceAll("http://btsow.bar/search/{query}", "{query}", name)
	resp, err := s.client.R().Get(url)
	if err != nil {
		return nil, err
	}
	b := resp.Body()
	doc, err := htmlquery.Parse(strings.NewReader(string(b)))
	if err != nil {
		return nil, err
	}
	list := htmlquery.Find(doc, "//div[@class='data-list']/div[@class='row']")
	res := &listCollectorResult{
		Title: name,
	}
	for _, v := range list {
		var title, magnet, size string
		if len("//a/@title") != 0 {
			titleNode := htmlquery.FindOne(v, "//a/@title")
			title = htmlquery.InnerText(titleNode)
		}
		if len("//a/@href") != 0 {
			magnetNode := htmlquery.FindOne(v, "//a/@href")
			magnet = htmlquery.InnerText(magnetNode)
		}
		if len("//div/div[1]") != 0 {
			sizeNode := htmlquery.FindOne(v, "//div/div[1]")
			size = htmlquery.InnerText(sizeNode)
		}
		item := &itemCollectorResult{
			Title:  title,
			Magnet: magnet,
			Size:   size,
		}
		res.Data = append(res.Data, item)
	}
	return res, nil
}

func (s *spiderRepo) torkitty(name string) (*listCollectorResult, error) {
	url := strings.ReplaceAll("http://www.torkitty.com/search/{query}", "{query}", name)
	resp, err := s.client.R().Get(url)
	if err != nil {
		return nil, err
	}
	b := resp.Body()
	doc, err := htmlquery.Parse(strings.NewReader(string(b)))
	if err != nil {
		return nil, err
	}
	list := htmlquery.Find(doc, "//table[@id='archiveResult']//tr/td[@class='action']/a[2]")
	res := &listCollectorResult{
		Title: name,
	}
	for _, v := range list {
		var title, magnet, size string
		if len("//@title") != 0 {
			titleNode := htmlquery.FindOne(v, "//@title")
			title = htmlquery.InnerText(titleNode)
		}
		if len("//@href") != 0 {
			magnetNode := htmlquery.FindOne(v, "//@href")
			magnet = htmlquery.InnerText(magnetNode)
		}
		if len("") != 0 {
			sizeNode := htmlquery.FindOne(v, "")
			size = htmlquery.InnerText(sizeNode)
		}
		item := &itemCollectorResult{
			Title:  title,
			Magnet: magnet,
			Size:   size,
		}
		res.Data = append(res.Data, item)
	}
	return res, nil
}

func (s *spiderRepo) (name string) (*listCollectorResult, error) {
	url := strings.ReplaceAll("http://sukebei.nyaa.si/?f=0&c=0_0&q={query}", "{query}", name)
	resp, err := s.client.R().Get(url)
	if err != nil {
		return nil, err
	}
	b := resp.Body()
	doc, err := htmlquery.Parse(strings.NewReader(string(b)))
	if err != nil {
		return nil, err
	}
	list := htmlquery.Find(doc, "")
	res := &listCollectorResult{
		Title: name,
	}
	for _, v := range list {
		var title, magnet, size string
		if len("") != 0 {
			titleNode := htmlquery.FindOne(v, "")
			title = htmlquery.InnerText(titleNode)
		}
		if len("") != 0 {
			magnetNode := htmlquery.FindOne(v, "")
			magnet = htmlquery.InnerText(magnetNode)
		}
		if len("") != 0 {
			sizeNode := htmlquery.FindOne(v, "")
			size = htmlquery.InnerText(sizeNode)
		}
		item := &itemCollectorResult{
			Title:  title,
			Magnet: magnet,
			Size:   size,
		}
		res.Data = append(res.Data, item)
	}
	return res, nil
}

func Get(e string, q string) (*listCollectorResult, error) {
	switch e {
		case "btsow":
		return spider.btsow(q)
		case "torkitty":
		return spider.torkitty(q)
		case "":
		return spider.(q)
	}
	return nil, nil
}
