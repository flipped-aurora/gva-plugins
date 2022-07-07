package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/es/api"
	"github.com/gin-gonic/gin"
)

type EsRouter struct{}

func (s *EsRouter) InitEsRouter(Router *gin.RouterGroup) {
	esRouter := Router.Use(middleware.OperationRecord())
	EsAdd := api.ApiGroupApp.ElasticsearchApi.EsAdd
	DelEs := api.ApiGroupApp.ElasticsearchApi.DelEs
	UpdateEs := api.ApiGroupApp.ElasticsearchApi.UpdateEs
	GetEsId := api.ApiGroupApp.ElasticsearchApi.GetEsId
	GetEsQuery := api.ApiGroupApp.ElasticsearchApi.GetEsQuery
	GetEsMultipleQuery := api.ApiGroupApp.ElasticsearchApi.GetEsMultipleQuery
	UpdateEsLocation := api.ApiGroupApp.ElasticsearchApi.UpdateEsLocation
	GetEsNearby := api.ApiGroupApp.ElasticsearchApi.GetEsNearby
	{
		esRouter.POST("esAdd", EsAdd)                          // 新增
		esRouter.DELETE("delEs", DelEs)                        // 删除
		esRouter.PUT("updateEs", UpdateEs)                     // 修改
		esRouter.GET("getEsId", GetEsId)                       // 根据ID查询
		esRouter.GET("getEsQuery", GetEsQuery)                 // 全局查询
		esRouter.GET("getEsMultipleQuery", GetEsMultipleQuery) // 全局多字段模糊查询
		esRouter.POST("updateEsLocation", UpdateEsLocation)    // 更新经纬度
		esRouter.GET("getEsNearby", GetEsNearby)               // 搜索附近的数据
	}
}
