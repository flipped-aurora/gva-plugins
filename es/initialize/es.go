package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	es_global "github.com/flipped-aurora/gin-vue-admin/server/plugin/es/global"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
)

func LoadService() {
	connectElasticsearch()
}

//连接Elasticsearch
func connectElasticsearch() {
	ElasticsearchCfg := es_global.GlobalConfig

	host := ElasticsearchCfg.Host
	port := ElasticsearchCfg.Port
	user := ElasticsearchCfg.User
	pass := ElasticsearchCfg.Pass

	var err error
	es_global.Elasticsearch, err = elastic.NewClient(
		//改变golang代码初始化client时的参数.   client, err := elastic.NewClient(elastic.SetSniff(false),elastic.SetURL(host…))  新增参数 elastic.SetSniff(false), 用于关闭 Sniff
		elastic.SetSniff(false),
		elastic.SetURL("http://"+host+":"+port),
		elastic.SetBasicAuth(user, pass),
	)
	if err != nil {
		global.GVA_LOG.Error("连接Elasticsearch出错：(连接失败,请检查ES服务是否开启！), err:", zap.Any("err", err))
		panic("连接Elasticsearch出错：" + err.Error())
	}

}
