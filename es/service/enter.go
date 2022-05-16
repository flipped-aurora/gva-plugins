package service

type ServiceGroup struct {
	ElasticsearchService
}

var ServiceGroupApp = new(ServiceGroup)
