package gtwenty

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

type ClientMock struct {
	body io.ReadCloser
}

func (c *ClientMock) Do(_ *http.Request) (*http.Response, error) {
	return c.MockResponse(), nil
}

func (c *ClientMock) MockResponse() *http.Response {
	return &http.Response{Body: c.body}
}

func TestQuery(t *testing.T) {
	queryTester(t, "testdata/input-single.json", "testdata/output-single.json")
}

func TestTransform(t *testing.T) {
	file := "test-output.html"
	reportItems := UnmarshallReportItems("testdata/input-all.json")
	list := UnmarshallProcessList("../../input/mas-regulatory-reporting-data-descriptor.json")
	reports := make([]Report, 0)
	for i, v := range list[0].Data {
		items := make([]ReportItem, 0)
		items = append(items, reportItems[i])
		reports = append(reports, Report{Name: v.Name,
			Items: items})
	}
	GenerateHtml3(reports, file, "data/template.html", "data/headers.txt")
	actual := ReadFile(file)
	expected := ReadFile("testdata/output-all.html")
	a := BuildArrayOfLines(actual)
	e := BuildArrayOfLines(expected)
	for i, v := range e {
		if i == len(a) {
			t.Fatal(fmt.Sprintf("GenerateHtml() = nil, expected %s", v))
		}
		t.Logf("comparing: %s with: %s", a[i], v)
		if strings.TrimSpace(a[i]) != strings.TrimSpace(v) {
			t.Error(fmt.Sprintf("GenerateHtml() = %s, expected %s", a[i], v))
		}
	}
}

func TestReportsProcessList(t *testing.T) {
	output := "testdata/output-single.json"
	expected := UnmarshallReportItems(output)
	list := UnmarshallProcessList("../../input/mas-regulatory-reporting-data-descriptor.json")
	for _, v := range list[0].Data {
		actual := Query(buildMock(output), "../../input/"+v.Input)
		t.Logf("comparing: %s with: %s", actual, expected)
		if !reflect.DeepEqual(actual, expected) {
			t.Error(fmt.Sprintf("Query() = %s, expected %s", actual, expected))
		}
	}
}

func queryTester(t *testing.T, input string, output string) {
	expected := UnmarshallReportItems(output)
	actual := Query(buildMock(output), input)
	t.Logf("comparing: %s with: %s", actual, expected)
	if !reflect.DeepEqual(actual, expected) {
		t.Error(fmt.Sprintf("Query() = %s, expected %s", actual, expected))
	}
}

func buildMock(file string) RegnosysClient {
	b := ReadFile(file)
	mock := RegnosysClient{
		Client: &ClientMock{
			body: ioutil.NopCloser(bytes.NewReader(b)),
		},
		Auth:   "",
		Cookie: "",
	}
	return mock
}
