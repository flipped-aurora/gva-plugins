package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	es_response "github.com/flipped-aurora/gin-vue-admin/server/plugin/es/model/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/es/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ElasticsearchApi
// @Description: ElasticsearchApi
//
type ElasticsearchApi struct {
}

type Employee struct {
	Type           int      `json:"type"`
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	BrandName      string   `json:"brandName"`
	OneClass       int      `json:"oneClass"`
	TwoClass       int      `json:"twoClass"`
	ThreeClass     int      `json:"threeClass"`
	BrandId        int      `json:"brandId"`
	SealId         int      `json:"sealId"`
	SealName       string   `json:"sealName"`
	OneClassName   string   `json:"oneClassName"`
	TwoClassName   string   `json:"twoClassName"`
	ThreeClassName string   `json:"threeClassName"`
	Price          string   `json:"price"`
	Img            string   `json:"img"`
	IsShow         int      `json:"isShow"`
	About          string   `json:"about"`
	Interests      []string `json:"interests"`
}

// EsAdd  新增ES数据
func (a *ElasticsearchApi) EsAdd(c *gin.Context) {

	data := Employee{
		1,
		3,
		"MSP430G2553IPMSPR",
		"NUVOTON(新唐) ",
		1,
		2,
		3,
		1,
		2,
		"LQMCUFP56",
		"处理器及微控制器",
		"单片机(MSP/MPU)",
		"",
		"10",
		"",
		2,
		"",
		[]string{""},
	}
	name := "test"
	err, list := service.ServiceGroupApp.CreateElasticsearch(name, "3", data)
	if err != nil {
		global.GVA_LOG.Error("新增失败!", zap.Any("err", err))
		response.FailWithMessage("新增失败", c)
	} else {
		response.OkWithData(list, c)
	}
}

// DelEs 删除ES数据
func (a *ElasticsearchApi) DelEs(c *gin.Context) {
	name := "test"
	err, list := service.ServiceGroupApp.DeleteElasticsearch(name, "3")
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Any("err", err))
		response.FailWithMessage("失败", c)
	} else {
		response.OkWithData(list, c)
	}
}

// UpdateEs 修改ES数据
func (a *ElasticsearchApi) UpdateEs(c *gin.Context) {
	name := "test"
	data := Employee{
		34,
		3,
		"MSP430G2553IPMSPR",
		"新唐",
		1,
		2,
		3,
		1,
		2,
		"LQMCUFP56",
		"处理器及微控制器",
		"单片机(MSP/MPU)",
		"",
		"10",
		"",
		2,
		"",
		[]string{""},
	}
	err, list := service.ServiceGroupApp.UpdateElasticsearch(name, "3", data)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Any("err", err))
		response.FailWithMessage("失败", c)
	} else {
		response.OkWithData(list, c)
	}
}

// GetEsId  通过id查找
func (a *ElasticsearchApi) GetEsId(c *gin.Context) {
	name := "test"
	err, list := service.ServiceGroupApp.GetIdElasticsearch(name, "3")
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Any("err", err))
		response.FailWithMessage("失败", c)
	} else {
		response.OkWithData(list, c)
	}
}

// GetEsQuery 搜索
func (a *ElasticsearchApi) GetEsQuery(c *gin.Context) {
	var pageInfo es_response.ElasticSearchSearch
	_ = c.ShouldBindQuery(&pageInfo)
	name := "test"
	if err, list := service.ServiceGroupApp.GetQueryElasticsearch(pageInfo, name, "brandName"); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(list, c)
	}
}

// GetEsMultipleQuery 多字段搜索
func (appHomeApi *AppHomeApi) GetEsMultipleQuery(c *gin.Context) {
	var pageInfo autocodeReq.ElasticSearchSearch
	_ = c.ShouldBindQuery(&pageInfo)
	name := "test"
	if err, list := service.ServiceGroupApp.GetQueryElasticsearch(pageInfo, name, "storeName", "barName", "brandName"); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(list, c)
	}
}

// GetEsMultipleQuery 多字段搜索
func (appHomeApi *AppHomeApi) GetEsMultipleQuery(c *gin.Context) {
	var pageInfo autocodeReq.ElasticSearchSearch
	_ = c.ShouldBindQuery(&pageInfo)
	name := "test"
	if err, list := service.ServiceGroupApp.GetQueryElasticsearch(pageInfo, name, "storeName", "barName", "brandName"); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(list, c)
	}
}

// UpdateEsLocation 更新经纬度
func (a *ElasticsearchApi) UpdateEsLocation(c *gin.Context) {
	name := "user_location"
	latitude := "30.5"
	longitude := "120.5"
	//用户ID或者商户ID
	id := "1"
	err, list := service.ServiceGroupApp.UpdateElasticsearchLocation(id, name, latitude, longitude)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Any("err", err))
		response.FailWithMessage("失败", c)
	} else {
		response.OkWithData(list, c)
	}
}

// GetEsNearby 查询附近的人
func (appHomeApi *AppHomeApi) GetEsNearby(c *gin.Context) {
	var pageInfo autocodeReq.ElasticSearchSearch
	_ = c.ShouldBindQuery(&pageInfo)
	name := "user_location"
	if err, list := service.ServiceGroupApp.GetElasticsearchNearby(pageInfo, name); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(list, c)
	}
}
