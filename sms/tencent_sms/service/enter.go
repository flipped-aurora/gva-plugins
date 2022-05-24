package service

type ServiceGroup struct {
	TencentSmsService
}

var ServiceGroupApp = new(ServiceGroup)
