package controllers

import (
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
	secretKey := "**************"
	// the request url(post)
	uri := "https://www.google.com/recaptcha/api/siteverify"
	// post data key:value
	params := url.Values{
		"secret":   {secretKey},
		"response": {c.GetString("g-recaptcha-response", "")},
	}

	// create a http request
	req, err := http.NewRequest("POST", uri, strings.NewReader(params.Encode()))
	// return error info if encounter any error
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	// set request header for post request
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// create a http client and emit the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	// here, res.body is not nil, so close resource when return
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	// here using gjson package for analyzing the request result
	if !gjson.Get(string(content), "success").Bool() {
		// render a error page if failed
		c.TplName = "error.html"
		return
	}

	// so render a success page if validation pass
	c.TplName = "success.html"

	// validation example:

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
