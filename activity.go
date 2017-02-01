package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Category struct {
	Name  string
	Score int
}

type Activity struct {
	Name string
	Category
}

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

func findCategory(name string) Category {
	return Category{Name: "Stuff", Score: 42}
}

func NewActivity(name string) Activity {
	return Activity{
		Name:     name,
		Category: findCategory(name),
	}
}
