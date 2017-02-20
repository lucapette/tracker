package tracker

import (
	"bytes"
	"os/exec"
	"strings"
)

const script = `
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

if activeURL is not "" then
  activeURL
else
  frontApp
end if
`

func GetActivityName() (string, error) {
	cmd := exec.Command("osascript", "-")
	cmd.Stdin = bytes.NewBufferString(script)
	output, err := cmd.Output()

	return strings.Replace(string(output), "\n", "", -1), err
}
