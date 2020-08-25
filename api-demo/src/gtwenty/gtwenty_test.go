package gtwenty

import (
	"bufio"
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

func (c *ClientMock) Do(req *http.Request) (*http.Response, error) {
	return c.MockResponse(), nil
}

func (c *ClientMock) MockResponse() *http.Response {
	return &http.Response{Body: c.body}
}

func TestQuery1(t *testing.T) {
	queryTester(t, "testdata/usecases1.json")
}

func TestQuery2(t *testing.T) {
	queryTester(t, "testdata/usecases2.json")
}

func TestTransform(t *testing.T) {
	file := "test-output"
	reportItems := UnmarshallReportItems("testdata/usecases2.json")
	GenerateHtml(reportItems, "test", file, "data/")
	actual, err := ioutil.ReadFile(file + ".html")
	if err != nil {
		t.Fatal(err)
	}
	expected, err := ioutil.ReadFile("testdata/template1.html")
	if err != nil {
		t.Fatal(err)
	}
	a := buildArrayOfWords(actual)
	e := buildArrayOfWords(expected)
	for i := 0; i < len(a); i++ {
		t.Logf("comparing: %s with: %s", a[i], e[i])
		if strings.TrimSpace(a[i]) != strings.TrimSpace(e[i]) {
			t.Error(fmt.Sprintf("GenerateHtml() = %s, expected %s", a[i], e[i]))
		}
	}
}

func TestReportsProcessList(t *testing.T) {
	expected := UnmarshallReportItems("testdata/usecases1.json")
	list := UnmarshallProcessList("../../input/mas-regulatory-reporting-data-descriptor.json")
	for _, v := range list[0].Data {
		actual := Query(buildMock(t, "testdata/usecases1.json"), "../../input/"+v.Input)
		t.Logf("comparing: %s with: %s", actual, expected)
		if !reflect.DeepEqual(actual, expected) {
			t.Error(fmt.Sprintf("Query() = %s, expected %s", actual, expected))
		}
	}
}

func queryTester(t *testing.T, useCase string) {
	expected := UnmarshallReportItems(useCase)
	actual := Query(buildMock(t, useCase), "testdata/input1.json")
	t.Logf("comparing: %s with: %s", actual, expected)
	if !reflect.DeepEqual(actual, expected) {
		t.Error(fmt.Sprintf("Query() = %s, expected %s", actual, expected))
	}
}

func buildMock(t *testing.T, file string) RegnosysClient {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}
	mock := RegnosysClient{
		Client: &ClientMock{
			body: ioutil.NopCloser(bytes.NewReader(b)),
		},
		Auth:   "",
		Cookie: "",
	}
	return mock
}

func buildArrayOfWords(b []byte) []string {
	var s []string
	aReader := bufio.NewScanner(bytes.NewReader(b))
	aReader.Split(bufio.ScanLines)
	for aReader.Scan() {
		s = append(s, aReader.Text())
	}
	return s
}
