package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"time"
)

var script = `
tell application "System Events"
  set frontApp to name of first application process whose frontmost is true

  set activeURL to ""
  if frontApp is "Google Chrome" then
    tell application "Google Chrome"
      set normalWindows to (windows whose mode is not "incognito")

            if length of normalWindows is greater than 0 then
                set activeURL to (get URL of active tab of (first item of normalWindows))
            end if
    end tell
  end if
end tell

frontApp & "," & (activeURL) as String
`

type category struct {
	name  string
	score int
}

type activity struct {
	name string
	category
}

func (a activity) store() error {
	b := bytes.NewBufferString(fmt.Sprintf("activity,category=%s,score=%d value=\"%s\"", a.category.name, a.category.score, a.name))
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

func findCategory(activity string) category {
	return category{name: "Stuff", score: 42}
}

func currentActivity() (ac activity) {
	cmd := exec.Command("osascript", "-")
	cmd.Stdin = bytes.NewBufferString(script)
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	values := strings.Split(string(output), ",")

	for i := range values {
		values[i] = strings.Replace(values[i], "\n", "", -1)
	}

	if len(values[1]) == 0 {
		ac.name = values[0]
	} else {
		ac.name = values[1]
	}

	ac.category = findCategory(ac.name)

	return ac
}

func main() {
	go func() {
		tick := time.Tick(1 * time.Second)
		for range tick {
			a := currentActivity()
			err := a.store()
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt, os.Kill)
	<-s
}
