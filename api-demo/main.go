package main

import (
	"apidemo/src/gtwenty"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

//URL is the REGnosys reporting API endpoint
const URL = "https://g20techsprint.prod.rosetta-technology.io/api/promoted/reg-report/%s"

func main() {
	projectFlag := flag.String("project", "", "The username together with the project and report name. E.g. joe.bloggs_auth0_123/g20/latest/MAS/SFA/MAS_2013")
	authFlag := flag.String("auth", "", "The authorisation token")
	cookieFlag := flag.String("cookie", "cookie", "The affinity cookie. E.g. ROSETTA_SESSION=0123456789.123.1234.123456")
	flag.Parse()
	empty := false
	flag.VisitAll(func(f *flag.Flag) {
		if f.Value.String() == "" {
			empty = true
		}
	})
	if empty {
		flag.Usage()
		os.Exit(1)
	}
	err := os.Mkdir("output", os.ModePerm)
	if err != nil {
		if !strings.Contains(err.Error(), "file exists") {
			log.Fatal(err)
		}
	}
	list := gtwenty.UnmarshallProcessList("input/mas-regulatory-reporting-data-descriptor.json")
	if len(list) < 1 {
		log.Fatal("Empty list passed")
	}
	reports := make([]gtwenty.Report, 0)
	for _, v := range list[0].Data {
		client := gtwenty.RegnosysClient{
			Client: &http.Client{},
			Url:    fmt.Sprintf(URL, *projectFlag),
			Auth:   *authFlag,
			Cookie: *cookieFlag,
		}
		reportItems := gtwenty.Query(client, "input/"+v.Input)
		reports = append(reports, gtwenty.Report{Name: v.Name,
			Items: reportItems})
	}
	gtwenty.GenerateHtml(reports, "output/report.html", "src/gtwenty/data/template.html", "src/gtwenty/data/headers.txt")
}
