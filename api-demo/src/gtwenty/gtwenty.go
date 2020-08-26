package gtwenty

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Query MAS
func Query(client RegnosysClient, input string) []ReportItem {
	b, err := ioutil.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}
	body := bytes.NewReader(b)
	contentType := "application/json"
	var reportItems []ReportItem
	req, err := http.NewRequest(http.MethodPost, client.Url, body)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-type", contentType)
	req.Header.Add("Authorization", client.Auth)
	req.Header.Add("Cookie", client.Cookie)
	resp, err := client.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()
	fmt.Printf("%s\n", responseBody)
	err = json.Unmarshal(responseBody, &reportItems)
	if err != nil {
		log.Fatal(err)
	}
	return reportItems
}

func GenerateHtml(reportItems []ReportItem, name string, output string, templatePath string) {
	b, err := ioutil.ReadFile(templatePath + "template.html")
	if err != nil {
		log.Fatal(err)
	}
	report := Report{Name: name, Items: reportItems}
	tlt, err := template.New("webpage").Parse(string(b))
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(output + ".html")
	if err != nil {
		log.Fatal("create file: ", err)
	}
	err = tlt.Execute(f, report)
	if err != nil {
		log.Fatal(err)
	}
}

func UnmarshallReportItems(useCase string) []ReportItem {
	var expected []ReportItem
	b, err := ioutil.ReadFile(useCase)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(b, &expected)
	if err != nil {
		log.Fatal(err)
	}
	return expected
}

func UnmarshallProcessList(list string) []ReportProcessList {
	var expected []ReportProcessList
	b, err := ioutil.ReadFile(list)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(b, &expected)
	if err != nil {
		log.Fatal(err)
	}
	return expected
}
