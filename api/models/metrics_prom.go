package models

import (
	"context"
	"devops/pkg/resp"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/prometheus/common/model"
	"github.com/shopspring/decimal"
	logs "github.com/souliot/siot-log"
)

// 主控节点运行信息
const (
	PROM_HOST_INFO_BASIC = `{__name__=~"loghub_process_max_fds|system_cpu_count|system_cpu_percent|system_host_temperature|system_disk_usage|system_mem_virtual",job="%s",type=~"logical|temperature|free|total|used|usedPercent|inodesUsedPercent|available|()",cpu=~"all|()",sensorKey=~"coretemp_physical_id_0|()",path=~"/|()"}`
	PROM_HOST_IO_STATUS  = `irate({__name__=~"system_disk_iocounter|system_net_info",job="%s",path=~"sda|()",type=~"readBytes|writeBytes|bytesRecv|bytesSent"}[5m])`
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
	DiskIOReadRate    model.SampleValue `json:"DiskIOReadRate,omitempty"`
	DiskIOWriteRate   model.SampleValue `json:"DiskIOWriteRate,omitempty"`
	NetIOReadRate     model.SampleValue `json:"NetIOReadRate,omitempty"`
	NetIOWriteRate    model.SampleValue `json:"NetIOWriteRate,omitempty"`
}

func (m *MetricsProm) HostInfo() (ls *List, errC *resp.Response, err error) {
	t := time.Now()
	job := "system"
	cache := make(map[model.LabelValue]*HostInfo)
	ls = new(List)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	basic, err := getVector(ctx, fmt.Sprintf(PROM_HOST_INFO_BASIC, job), t)
	if err != nil {
		logs.Info(err)
		return
	}
	io, err := getVector(ctx, fmt.Sprintf(PROM_HOST_IO_STATUS, job), t)
	if err != nil {
		logs.Info(err)
		return
	}

	for _, v := range basic {
		name := v.Metric["__name__"]
		instance := v.Metric["instance"]
		switch name {
		case model.LabelValue("loghub_process_max_fds"):
			if host, ok := cache[instance]; ok {
				host.Fds = v.Value
			} else {
				host := &HostInfo{
					Fds: v.Value,
				}
				cache[instance] = host
			}
		case model.LabelValue("system_cpu_count"):
			if host, ok := cache[instance]; ok {
				host.CPUCount = v.Value
			} else {
				host := &HostInfo{
					CPUCount: v.Value,
				}
				cache[instance] = host
			}
		case model.LabelValue("system_cpu_percent"):
			per, _ := decimal.NewFromFloat(float64(v.Value)).Round(2).Float64()
			if host, ok := cache[instance]; ok {
				host.CPUPercent = model.SampleValue(per)
			} else {
				host := &HostInfo{
					CPUPercent: model.SampleValue(per),
				}
				cache[instance] = host
			}
		case model.LabelValue("system_host_temperature"):
			if host, ok := cache[instance]; ok {
				host.Temp = v.Value
			} else {
				host := &HostInfo{
					Temp: v.Value,
				}
				cache[instance] = host
			}
		case model.LabelValue("system_mem_virtual"):
			typ := v.Metric["type"]
			switch typ {
			case model.LabelValue("free"):
				if host, ok := cache[instance]; ok {
					host.MemFree = v.Value
				} else {
					host := &HostInfo{
						MemFree: v.Value,
					}
					cache[instance] = host
				}
			case model.LabelValue("total"):
				if host, ok := cache[instance]; ok {
					host.MemTotal = v.Value
				} else {
					host := &HostInfo{
						MemTotal: v.Value,
					}
					cache[instance] = host
				}
			case model.LabelValue("used"):
				if host, ok := cache[instance]; ok {
					host.MemUsed = v.Value
				} else {
					host := &HostInfo{
						MemUsed: v.Value,
					}
					cache[instance] = host
				}
			case model.LabelValue("usedPercent"):
				per, _ := decimal.NewFromFloat(float64(v.Value)).Round(2).Float64()
				if host, ok := cache[instance]; ok {
					host.MemPercent = model.SampleValue(per)
				} else {
					host := &HostInfo{
						MemPercent: model.SampleValue(per),
					}
					cache[instance] = host
				}
			case model.LabelValue("available"):
				if host, ok := cache[instance]; ok {
					host.MemAvailable = v.Value
				} else {
					host := &HostInfo{
						MemAvailable: v.Value,
					}
					cache[instance] = host
				}
			}
		case model.LabelValue("system_disk_usage"):
			typ := v.Metric["type"]
			switch typ {
			case model.LabelValue("free"):
				if host, ok := cache[instance]; ok {
					host.DiskFree = v.Value
				} else {
					host := &HostInfo{
						DiskFree: v.Value,
					}
					cache[instance] = host
				}
			case model.LabelValue("total"):
				if host, ok := cache[instance]; ok {
					host.DiskTotal = v.Value
				} else {
					host := &HostInfo{
						DiskTotal: v.Value,
					}
					cache[instance] = host
				}
			case model.LabelValue("used"):
				if host, ok := cache[instance]; ok {
					host.DiskUsed = v.Value
				} else {
					host := &HostInfo{
						DiskUsed: v.Value,
					}
					cache[instance] = host
				}
			case model.LabelValue("usedPercent"):
				per, _ := decimal.NewFromFloat(float64(v.Value)).Round(2).Float64()
				if host, ok := cache[instance]; ok {
					host.DiskPercent = model.SampleValue(per)
				} else {
					host := &HostInfo{
						DiskPercent: model.SampleValue(per),
					}
					cache[instance] = host
				}
			case model.LabelValue("inodesUsedPercent"):
				per, _ := decimal.NewFromFloat(float64(v.Value)).Round(2).Float64()
				if host, ok := cache[instance]; ok {
					host.DiskInodesPercent = model.SampleValue(per)
				} else {
					host := &HostInfo{
						DiskInodesPercent: model.SampleValue(per),
					}
					cache[instance] = host
				}
			}
		}
	}

	for _, v := range io {
		instance := v.Metric["instance"]
		typ := v.Metric["type"]
		per, _ := decimal.NewFromFloat(float64(v.Value)).Round(2).Float64()
		switch typ {
		case model.LabelValue("readBytes"):
			if host, ok := cache[instance]; ok {
				host.DiskIOReadRate = model.SampleValue(per)
			} else {
				host := &HostInfo{
					DiskIOReadRate: model.SampleValue(per),
				}
				cache[instance] = host
			}
		case model.LabelValue("writeBytes"):
			if host, ok := cache[instance]; ok {
				host.DiskIOWriteRate = model.SampleValue(per)
			} else {
				host := &HostInfo{
					DiskIOWriteRate: model.SampleValue(per),
				}
				cache[instance] = host
			}
		case model.LabelValue("bytesRecv"):
			if host, ok := cache[instance]; ok {
				host.NetIOReadRate = model.SampleValue(per)
			} else {
				host := &HostInfo{
					NetIOReadRate: model.SampleValue(per),
				}
				cache[instance] = host
			}
		case model.LabelValue("bytesSent"):
			if host, ok := cache[instance]; ok {
				host.NetIOWriteRate = model.SampleValue(per)
			} else {
				host := &HostInfo{
					NetIOWriteRate: model.SampleValue(per),
				}
				cache[instance] = host
			}
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
	sort.Sort(HostList(cs))
	ls.Lists = cs
	ls.Total = int64(len(cs))
	return
}

type HostList []*HostInfo

func (m HostList) Len() int {
	return len(m)
}

func (m HostList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m HostList) Less(i, j int) bool {
	return lessIP(m[i].Host, m[j].Host)
}
