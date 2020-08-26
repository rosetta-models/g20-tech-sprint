package gtwenty

import "net/http"

//ReportItem output
type ReportItem struct {
	Identifier struct {
		RegRegime              string   `json:"regRegime"`
		Mandates               []string `json:"mandates"`
		Name                   string   `json:"name"`
		GeneratedJavaClassName string   `json:"generatedJavaClassName"`
	} `json:"identifier"`
	UseCases []struct {
		UseCase string `json:"useCase"`
		Fields  []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
			Issue string `json:"issue"`
		} `json:"fields"`
	} `json:"useCases"`
}

type Report struct {
	Name  string
	Items []ReportItem
}

type ReportType2 struct {
	Headers []string
	Items   []string
}

type ReportItemType2 struct {
	Items []string
}

type ReportType3 struct {
	Headers []string
	Reports []ReportItemType2
}

type ReportProcessList struct {
	DataSetName       string   `json:"dataSetName"`
	InputType         string   `json:"inputType"`
	ApplicableReports []string `json:"applicableReports"`
	Data              []struct {
		Name  string `json:"name"`
		Input string `json:"input"`
	} `json:"data"`
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type RegnosysClient struct {
	Client HttpClient
	Url    string
	Auth   string
	Cookie string
}
