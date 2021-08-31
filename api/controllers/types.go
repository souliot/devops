package controllers

type ServiceMeta struct {
	Id             string
	Path           string
	Typ            string
	Address        string
	OutAddress     string
	Version        string
	MetricsType    string
	MetricsAddress string
	Ext            interface{}
	Status         bool
}
