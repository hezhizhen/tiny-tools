package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/hezhizhen/tiny-tools/utilz"
)

func main() {
	specifiedDate := flag.String("date", "", "specified date. e.g.: 2016.01.02")
	editor := flag.String("editor", "mvim", "specified editor. e.g.: macvim, vimr, atom, code")
	flag.Parse()
	now := time.Now()
	var err error
	if *specifiedDate != "" {
		now, err = time.Parse("2006.01.02", *specifiedDate)
		utilz.Check(err)
	}
	if *editor == "macvim" {
		*editor = "mvim"
	}
	if *editor == "code" {
		*editor = "code-insiders"
	}
	// get right directory
	dir := "/Users/hezhizhen/Dropbox/Diary"
	filename := fmt.Sprintf("%4d-%02d-%02d.md", now.Year(), now.Month(), now.Day())
	filepath := fmt.Sprintf("%s/%s", dir, filename)
	_, err = os.Stat(filepath)
	if os.IsNotExist(err) {
		f, err := os.Create(filepath)
		utilz.Check(err)
		// write title
		f.WriteString(fmt.Sprintf("# %4d.%02d.%02d\n", now.Year(), now.Month(), now.Day()))
		f.Close()
	}
	// open file
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	utilz.Check(err)
	defer f.Close()
	// append current time
	_, err = f.WriteString(fmt.Sprintf("\n## %02d:%02d\n", now.Hour(), now.Minute()))
	utilz.Check(err)
	// open in macvim
	cmd := exec.Command(*editor, filepath)
	utilz.Check(cmd.Run())
}