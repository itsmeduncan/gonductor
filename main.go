/*
Package main implements a simple CLI tool for checking train statuses.
*/
package main

import (
	"encoding/xml"
	"fmt"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type service struct {
	ResponseCode string `xml:"responsecode"`
	Timestamp    string `xml:"timestamp"`
	Subway       subway `xml:"subway"`
}

type subway struct {
	Line []line `xml:"line"`
}

type line struct {
	Name   string `xml:"name"`
	Status string `xml:"status"`
	Text   string `xml:"text"`
	Date   string
	Time   string
}

var availableLines = []string{
	"123",
	"456",
	"7",
	"ACE",
	"BDFM",
	"G",
	"JZ",
	"L",
	"NQR",
	"S",
	"SIR",
}

var colorForTmux = map[string]string{
	"GOOD SERVICE":   "#[fg=green]√",
	"PLANNED WORK":   "#[fg=colour3]−",
	"SERVICE CHANGE": "#[fg=colour214]−",
	"DELAYS":         "#[fg=red]☓",
}

var tmux = false

func queryForStatusOf(line string) {
	url := "http://web.mta.info/status/serviceStatus.txt"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	service := service{}
	body, err := ioutil.ReadAll(res.Body)
	xml.Unmarshal(body, &service)

	for _, l := range service.Subway.Line {
		if l.Name == line {
			if tmux {
				fmt.Printf("%s", colorForTmux[l.Status])
			} else {
				fmt.Printf("%s %s\n", l.Name, l.Status)
			}
		}
	}
}

func statusOf(line string) {
	for _, availableLine := range availableLines {
		if line == availableLine {
			queryForStatusOf(line)
		}
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "gonductor"
	app.Usage = "Simple tool for MTA subway status"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "line,l",
			Usage: "subway line to check the status of",
		},
		cli.BoolFlag{
			Name:  "tmux,t",
			Usage: "turn tmux colorization on",
		},
	}
	app.Action = func(c *cli.Context) {
		line := c.String("line")
		tmux = c.Bool("tmux")
		statusOf(line)
	}

	app.Run(os.Args)
}
