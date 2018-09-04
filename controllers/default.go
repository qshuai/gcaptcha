package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/astaxie/beego"
	"github.com/tidwall/gjson"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
}

func (c *MainController) Post() {
	// your secret key
	secretKey := "****"
	uri := "https://www.google.com/recaptcha/api/siteverify"
	params := url.Values{
		"secret":   {secretKey},
		"response": {c.GetString("g-recaptcha-response", "") + "lsjfaj"},
	}
	req, err := http.NewRequest("POST", uri, strings.NewReader(params.Encode()))
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON(true)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		c.Data["json"] = err.Error()
		fmt.Println("hello")
		c.ServeJSON(true)
		return
	}

	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
		fmt.Println("world")
		c.ServeJSON(true)
		return
	}

	if !gjson.Get(string(content), "success").Bool() {
		c.TplName = "error.html"
	}

	c.TplName = "success.html"

	// success:
	// {
	// 		"success": true,
	// 		"challenge_ts": "2018-09-04T14:26:23Z",
	// 		"hostname": "localhost"
	// }

	// failed:
	// {
	// 		success": false,
	// 		"error-codes": ["invalid-input-response"]
	// }
}
