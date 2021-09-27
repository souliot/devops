package models

// 主控节点运行信息
const (
	PROM_HOST_MAX_FDS        = `loghub_process_max_fds{job="%s"}`
	PROM_HOST_CPU_COUNT      = `system_cpu_count{job="%s",type="logical"}`
	PROM_HOST_CPU_PERCENT    = `system_cpu_percent{job="%s",cpu="all"}`
	PROM_HOST_DISK_FREE      = `system_disk_usage{job="%s",path="/",type=~"free|total|used|usedPercent|inodesUsedPercent"}`
	PROM_HOST_TEMP           = `system_host_temperature{job="%s",sensorKey="coretemp_physical_id_0",type="temperature"}`
	PROM_HOST_MEM_FREE       = `system_mem_virtual{job="%s",type=~"available|free|total|used|usedPercent"}`
	PROM_HOST_DISK_IO_STATUS = `irate(system_disk_iocounter{job="%s",path=~"sda",type=~"bytesRecv|bytesSent"}[5m])`
	PROM_HOST_NET_IO_STATUS  = `irate(system_net_info{job="%s",type=~"bytesRecv|bytesSent"}[5m])`
)

type MetricsProm struct{}

func (m *MetricsProm) HostInfo() {
	//
}
