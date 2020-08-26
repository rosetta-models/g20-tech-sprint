package gtwenty

import (
	"bufio"
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
	body := bytes.NewReader(ReadFile(input))
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

func GenerateHtml3(reports []Report, output string, templateFile string, headersFile string) {
	headers := BuildArrayOfLines(ReadFile(headersFile))
	var templateReport ReportType3
	templateReport.Headers = headers
	templateReport.Reports = make([]ReportItemType2, 0)
	for _, report := range reports {
		headerValueMap := make(map[string]string)
		for _, header := range headers {
			headerValueMap[header] = ""
		}
		headerValueMap["Use Case"] = report.Name
		for _, v := range report.Items[0].UseCases[0].Fields {
			headerValueMap[v.Name] = v.Value
		}
		items := make([]string, 0)
		for _, v := range headers {
			items = append(items, headerValueMap[v])
		}
		templateReport.Reports = append(templateReport.Reports, ReportItemType2{Items: items})
	}
	tlt, err := template.New("webpage").Parse(string(ReadFile(templateFile)))
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(output)
	if err != nil {
		log.Fatal("create file: ", err)
	}
	log.Printf("Writing file: %s\n", output)
	err = tlt.Execute(f, templateReport)
	if err != nil {
		log.Fatal(err)
	}
}

func UnmarshallReportItems(useCase string) []ReportItem {
	var expected []ReportItem
	err := json.Unmarshal(ReadFile(useCase), &expected)
	if err != nil {
		log.Fatal(err)
	}
	return expected
}

func UnmarshallProcessList(list string) []ReportProcessList {
	var expected []ReportProcessList
	err := json.Unmarshal(ReadFile(list), &expected)
	if err != nil {
		log.Fatal(err)
	}
	return expected
}

func BuildArrayOfLines(b []byte) []string {
	var s []string
	aReader := bufio.NewScanner(bytes.NewReader(b))
	aReader.Split(bufio.ScanLines)
	for aReader.Scan() {
		s = append(s, aReader.Text())
	}
	return s
}

func ReadFile(file string) []byte {
	log.Printf("Reading file: %s\n", file)
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return b
}
