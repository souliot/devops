package test

import (
	"devops/pkg/fileutil"
	"html/template"
	"os"
	"path/filepath"
	"testing"
)

const (
	PROM_CONF_TPL = `global:
	scrape_interval:     60s
	evaluation_interval: 60s

scrape_configs:{{ range . }}{{ $le:=len .Targets }}{{ if and (eq .ConfigsType "static") (gt $le 0) }}{{ $le = add $le -1 }}
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

type Job struct {
	JobName     string   `json:"job_name"`
	MetricsPath string   `json:"metrics_path"`
	Scheme      string   `json:"scheme"`
	ConfigsType string   `json:"configs_type"`
	Targets     []string `json:"targets"`
	Url         string   `json:"url"`
}

func add(a, b int) int {
	return a + b
}

func TestTpl(t *testing.T) {
	tmpl, err := template.New("prom").Funcs(template.FuncMap{
		"add": add,
	}).Parse(PROM_CONF_TPL)
	if err != nil {
		t.Fatal(err)
	}
	jobs := make([]*Job, 0)
	j1 := &Job{
		JobName:     "minio",
		MetricsPath: "/minio/prometheus/metrics",
		Scheme:      "http",
		ConfigsType: "static",
		Targets:     []string{"192.168.0.4:9000", "192.168.0.35:9000"},
	}
	jobs = append(jobs, j1)
	j2 := &Job{
		JobName:     "node_export",
		MetricsPath: "/metrics",
		Scheme:      "http",
		ConfigsType: "http",
		Url:         "http://192.168.0.252:8080/v1/export/type/system",
	}
	jobs = append(jobs, j2)

	path := "tmpl"
	err = fileutil.TouchDirAll(path)
	if err != nil {
		t.Fatal(err)
	}
	file := filepath.Join(path, "prometheus.yaml")
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	err = tmpl.Execute(f, jobs)
	if err != nil {
		t.Fatal(err)
	}
}
