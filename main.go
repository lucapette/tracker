package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"os/signal"
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

func run() string {
	cmd := exec.Command("osascript", "-")
	cmd.Stdin = bytes.NewBufferString(script)
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(output)
}

func main() {
	go func() {
		c := time.Tick(1 * time.Second)
		for range c {
			log.Printf(run())
		}
	}()

	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt, os.Kill)
	<-s
}
