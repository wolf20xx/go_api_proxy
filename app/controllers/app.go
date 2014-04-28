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
	resp, err := http.Get(pxyurl)
	FatalLog(err)
	defer resp.Body.Close()
	revel.INFO.Println(resp.Body)
	//htmtexts := []struct {
	//    Text string
	//}{}
	revel.INFO.Println(resp)
	//err = json.NewDecoder(ret.Body).Decode(&htmtexts)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	FatalLog(err)
	revel.INFO.Println(string(bodyBytes))
	c.Response.ContentType = "application/json"
	return c.RenderText(string(bodyBytes))
}
