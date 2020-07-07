package main

import (
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"
)

var deviceNames = []string{
	"Wacom HID 5173 Pen stylus",
	"Wacom HID 5173 Pen eraser",
	"Wacom HID 5173 Finger touch",
}

func init() {
	sort.Strings(deviceNames)
	if !sort.StringsAreSorted(deviceNames) {
		log.Fatalln("assert failed: deviceNames are not sorted")
	}
}

func runCommandCutByLineBreak(path string, args ...string) (result []string, err error) {
	cmd := exec.Command(path, args...)
	cmd.Env = os.Environ()

	ret, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	return strings.Split(string(ret), "\n"), nil
}

func main() {
	names, err := runCommandCutByLineBreak("xinput", "list", "--name-only")
	if err != nil {
		log.Fatalln(err)
	}
	ids, err := runCommandCutByLineBreak("xinput", "list", "--id-only")
	if err != nil {
		log.Fatalln(err)
	}
	for i, name := range names {
		if ind := sort.SearchStrings(deviceNames, name); ind < len(deviceNames) && deviceNames[ind] == name {
			err := exec.Command("xinput", "disable", ids[i]).Run()
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}
