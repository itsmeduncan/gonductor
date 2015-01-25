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

type Service struct {
	ResponseCode string `xml:"responsecode"`
	Timestamp    string `xml:"timestamp"`
	Subway       Subway `xml:"subway"`
}

type Subway struct {
	Line []Line `xml:"line"`
}

type Line struct {
	Name   string `xml:"name"`
	Status string `xml:"status"`
	Text   string `xml:"text"`
	Date   string
	Time   string
}

var AvailableLines = []string{
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

var ColorForTmux = map[string]string{
	"GOOD SERVICE":   "#[fg=green]",
	"PLANNED WORK":   "#[fg=colour3]",
	"SERVICE CHANGE": "#[fg=colour214]",
	"DELAYS":         "#[fg=red]",
}

var tmux bool = false

func queryForStatusOf(l string) {
	url := "http://web.mta.info/status/serviceStatus.txt"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	service := Service{}
	body, err := ioutil.ReadAll(res.Body)
	xml.Unmarshal(body, &service)

	for _, line := range service.Subway.Line {
		if line.Name == l {
			if tmux {
				fmt.Printf("%s%s", ColorForTmux[line.Status], line.Name)
			} else {
				fmt.Println(line.Status)
			}
		}
	}
}

func statusOf(line string) {
	for _, availableLine := range AvailableLines {
		if line == availableLine {
			queryForStatusOf(line)
		}
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "gonductor"
	app.Usage = "Simple tool for MTA subway status"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "line,l",
			Value: "123",
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
