package global

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/es/config"
	"github.com/olivere/elastic/v7"
)

//var GlobalConfig = new(config.Elasticsearch)

var (
	GlobalConfig  = new(config.Elasticsearch)
	Elasticsearch *elastic.Client
)
