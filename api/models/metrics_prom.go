package models

import (
	"context"
	"devops/pkg/resp"
	"fmt"
	"strings"
	"time"

	"github.com/prometheus/common/model"
	"github.com/shopspring/decimal"
)

// 主控节点运行信息
const (
	PROM_HOST_MAX_FDS        = `loghub_process_max_fds{job="%s"}`
	PROM_HOST_CPU_COUNT      = `system_cpu_count{job="%s",type="logical"}`
	PROM_HOST_CPU_PERCENT    = `system_cpu_percent{job="%s",cpu="all"}`
	PROM_HOST_TEMP           = `system_host_temperature{job="%s",sensorKey="coretemp_physical_id_0",type="temperature"}`
	PROM_HOST_DISK_FREE      = `system_disk_usage{job="%s",path="/",type=~"free|total|used|usedPercent|inodesUsedPercent"}`
	PROM_HOST_MEM_FREE       = `system_mem_virtual{job="%s",type=~"available|free|total|used|usedPercent"}`
	PROM_HOST_DISK_IO_STATUS = `irate(system_disk_iocounter{job="%s",path=~"sda",type=~"readBytes|writeBytes"}[5m])`
	PROM_HOST_NET_IO_STATUS  = `irate(system_net_info{job="%s",type=~"bytesRecv|bytesSent"}[5m])`
)

type MetricsProm struct{}

type HostInfo struct {
	Host              string
	Fds               model.SampleValue `json:"Fds,omitempty"`
	Temp              model.SampleValue `json:"Temp,omitempty"`
	CPUCount          model.SampleValue `json:"CPUCount,omitempty"`
	CPUPercent        model.SampleValue `json:"CPUPercent,omitempty"`
	MemFree           model.SampleValue `json:"MemFree,omitempty"`
	MemTotal          model.SampleValue `json:"MemTotal,omitempty"`
	MemAvailable      model.SampleValue `json:"MemAvailable,omitempty"`
	MemUsed           model.SampleValue `json:"MemUsed,omitempty"`
	MemPercent        model.SampleValue `json:"MemPercent,omitempty"`
	DiskFree          model.SampleValue `json:"DiskFree,omitempty"`
	DiskTotal         model.SampleValue `json:"DiskTotal,omitempty"`
	DiskUsed          model.SampleValue `json:"DiskUsed,omitempty"`
	DiskPercent       model.SampleValue `json:"DiskPercent,omitempty"`
	DiskInodesPercent model.SampleValue `json:"DiskInodesPercent,omitempty"`
}

func (m *MetricsProm) HostInfo() (ls *List, errC *resp.Response, err error) {
	job := "system"
	cache := make(map[model.LabelValue]*HostInfo)
	ls = new(List)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fds, err := getVector(fmt.Sprintf(PROM_HOST_MAX_FDS, job), ctx)
	if err != nil {
		return
	}
	temp, err := getVector(fmt.Sprintf(PROM_HOST_TEMP, job), ctx)
	if err != nil {
		return
	}
	cpu_count, err := getVector(fmt.Sprintf(PROM_HOST_CPU_COUNT, job), ctx)
	if err != nil {
		return
	}
	cpu_percent, err := getVector(fmt.Sprintf(PROM_HOST_CPU_PERCENT, job), ctx)
	if err != nil {
		return
	}

	for i, v := range fds {
		var host *HostInfo
		var ok bool
		if host, ok = cache[v.Metric["instance"]]; ok {
			host.Fds = v.Value
		} else {
			host = &HostInfo{
				Fds: v.Value,
			}
			cache[v.Metric["instance"]] = host
		}

		if i < len(temp) {
			host.Temp = temp[i].Value
		}
		if i < len(cpu_count) {
			host.CPUCount = cpu_count[i].Value
		}
		if i < len(cpu_percent) {
			per, _ := decimal.NewFromFloat(float64(cpu_percent[i].Value)).Round(2).Float64()
			host.CPUPercent = model.SampleValue(per)
		}
	}
	cs := make([]*HostInfo, 0)
	for k, v := range cache {
		host := string(k)
		hosts := strings.Split(host, ":")
		if len(hosts) > 0 {
			v.Host = hosts[0]
		} else {
			v.Host = string(k)
		}
		cs = append(cs, v)
	}
	ls.Lists = cs
	ls.Total = int64(len(cs))
	return
}
