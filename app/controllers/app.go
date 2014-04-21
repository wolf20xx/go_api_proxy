package controllers

import (
    "github.com/revel/revel"
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

func (c App) ApiProxy() revel.Result {
    revel.INFO.Println("ApiProxyTest")
    var pxyurl string
    c.Params.Bind(&pxyurl,"pxyurl")
    resp, err := http.Get(pxyurl)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    revel.INFO.Println(resp.Body)
    //htmtexts := []struct {
    //    Text string
    //}{}
    revel.INFO.Println(resp)
    //err = json.NewDecoder(ret.Body).Decode(&htmtexts)
    bodyBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    revel.INFO.Println(string(bodyBytes))
    c.Response.ContentType = "application/json"
    return c.RenderText(string(bodyBytes))
}
