package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Activity struct {
	Name string
	Category
}

var client *http.Client

func (a Activity) Store() error {
	measurement := fmt.Sprintf("activity,category=%s,app=\"%s\" value=1i,score=%di", a.Category.Name, a.Name, a.Category.Score)

	b := bytes.NewBufferString(measurement)
	req, err := http.NewRequest("POST", "http://localhost:8086/write", b)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "")
	req.Header.Set("User-Agent", "lucapette/tracker")

	params := req.URL.Query()
	params.Set("db", "me")
	req.URL.RawQuery = params.Encode()

	resp, err := client.Do(req)
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
	if activity, ok := Activities[frontApp]; ok {
		return activity
	}

	url, err := url.Parse(frontApp)
	if err != nil {
		return Activity{Name: frontApp, Category: Categories["Uncategorized"]}
	}

	if activity, ok := Activities[url.Hostname()]; ok {
		return activity
	}

	return Activity{Name: frontApp, Category: Categories["Uncategorized"]}
}

func init() {
	client = &http.Client{Timeout: 100 * time.Millisecond}
}
