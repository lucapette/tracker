package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Category struct {
	Name  string
	Score int
}

var Communication = Category{Name: "Communication", Score: 0}
var Development = Category{Name: "Development", Score: 1}
var Social = Category{Name: "Social", Score: -1}
var Uncategorized = Category{Name: "Uncategorized", Score: 0}

var categories map[string]Category

type Activity struct {
	Name string
	Category
}

var UnknownActivity = Activity{Name: "Unknown", Category: Uncategorized}

func (a Activity) Store() error {
	b := bytes.NewBufferString(fmt.Sprintf("activity,category=%s,score=%d value=\"%s\"", a.Category.Name, a.Category.Score, a.Name))
	req, err := http.NewRequest("POST", "http://localhost:8086/write", b)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "")
	req.Header.Set("User-Agent", "lucapette/t")

	params := req.URL.Query()
	params.Set("db", "me")
	req.URL.RawQuery = params.Encode()

	c := &http.Client{Timeout: 100 * time.Millisecond}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		var err = fmt.Errorf(string(body))
		return err
	}

	return nil
}

func NewActivity(frontApp string) Activity {
	if category, ok := categories[frontApp]; ok {
		return Activity{Name: frontApp, Category: category}
	}

	url, err := url.Parse(frontApp)
	if err != nil {
		return UnknownActivity
	}

	if category, ok := categories[url.Host]; ok {
		return Activity{Name: url.Host, Category: category}
	}

	return UnknownActivity
}

func init() {
	categories = map[string]Category{
		"iTerm2":      Development,
		"twitter.com": Social,
		"airmail":     Communication,
	}
}
