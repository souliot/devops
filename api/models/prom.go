package models

import (
	"devops/pkg/fileutil"
	"devops/pkg/resp"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

const (
	PROM_CONF_TPL = `global:
  scrape_interval:     {{ .ScrapeInterval }}s
  evaluation_interval: {{ .EvaluationInterval }}s

scrape_configs:{{ range .Jobs }}{{ $le:=len .Targets }}{{ if and (eq .ConfigsType "static") (gt $le 0) }}{{ $le = add $le -1 }}
  - job_name: {{ .JobName }}
    metrics_path: {{ .MetricsPath }}
    scheme: {{ .Scheme }}
    static_configs: 
      - targets: [{{ range $k,$v:=.Targets }}'{{$v}}'{{ if ne $k $le }},{{ end }}{{ end }}]
{{ else if and (eq .ConfigsType "http") (ne .Url "") }}
  - job_name: {{ .JobName }}
    metrics_path: {{ .MetricsPath }}
    scheme: {{ .Scheme }}
    http_sd_configs:
      - url: "{{ .Url }}"
{{ end }}
{{ end }}`
)

type Prom struct {
	Jobs               []*PromJob
	Path               string `json:"path"`
	ScrapeInterval     int    `json:"scrapeInterval"`
	EvaluationInterval int    `json:"evaluationInterval"`
}

func (m *Prom) BuildConfiger() (errC *resp.Response, err error) {
	ls, errC, err := (&PromJob{}).All()
	if err != nil {
		return
	}
	m.Jobs = ls.Lists.([]*PromJob)
	tmpl, err := template.New("prom").Funcs(template.FuncMap{
		"add": add,
	}).Parse(PROM_CONF_TPL)
	if err != nil {
		errC = resp.ErrNewTmpl
		errC.MoreInfo = err.Error()
		return
	}

	m.Path = strings.TrimPrefix(m.Path, "/")
	if m.Path == "" {
		m.Path = "prom"
	}

	err = fileutil.TouchDirAll(m.Path)
	if err != nil {
		errC = resp.ErrNewTmpl
		errC.MoreInfo = err.Error()
		return
	}
	file := filepath.Join(m.Path, "prometheus.yml")
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		errC = resp.ErrOpenFile
		errC.MoreInfo = err.Error()
		return
	}
	defer f.Close()
	err = tmpl.Execute(f, m)
	if err != nil {
		errC = resp.ErrTransferTmpl
		errC.MoreInfo = err.Error()
		return
	}
	return
}

func (m *Prom) Reload() (errC *resp.Response, err error) {
	url := fmt.Sprintf("http://%s/-/reload", PromAddress)
	res, err := httpCli.R().Post(url)
	if err != nil {
		errC = resp.ErrHttpRequest
		errC.MoreInfo = err.Error()
		return
	}
	if res.StatusCode() != 200 {
		err = fmt.Errorf("%v", res.Error())
		errC = resp.ErrHttpResponseData
		errC.MoreInfo = err.Error()
		return
	}
	return
}

func add(a, b int) int {
	return a + b
}
