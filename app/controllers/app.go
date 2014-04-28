package controllers

import (
	"github.com/revel/revel"
	//"github.com/snikch/revel-redis/app"
	"io/ioutil"
	"log"
	"net/http"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

type ProxyQuery struct {
	url string
}

func FatalLog(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c App) ApiProxy(pxyurl string) revel.Result {
	revel.INFO.Println("ApiProxyTest")
	//var pxyurl string
	//c.Params.Bind(&pxyurl, "pxyurl")
	c.Validation.Required(pxyurl).Message("proxyurlAPI needs originAPIURL")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.RenderText("API Validation Errors Occur")
	}
	ch := make(chan string)
	go func() {
		AsyncGetApi(pxyurl, ch)
	}()
	res := <-ch

	revel.INFO.Println(res)
	c.Response.ContentType = "application/json"
	return c.RenderText(res)
}

func AsyncGetApi(url string, ch chan string) {
	resp, err := http.Get(url)
	FatalLog(err)
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	FatalLog(err)
	ch <- string(bodyBytes)
}
