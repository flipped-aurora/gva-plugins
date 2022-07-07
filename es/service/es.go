package service

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/es/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/es/model/response"
	"strconv"
)

type ElasticsearchService struct{}

//
// CreateElasticsearch
//  @Description: 新增Elasticsearch数据
//  @param name 索引名
//  @param esId 索引ID
//  @param data 数据
//  @return err 错误
//  @return ret 返回值
//
func (e *ElasticsearchService) CreateElasticsearch(name string, esId string, data interface{}) (err error, ret interface{}) {

	get, err := global.Elasticsearch.Index().Index(name).Id(esId).BodyJson(data).Do(context.Background())
	if err != nil {
		return err, nil
	}
	return err, get
}

//
// DeleteElasticsearch
//  @Description: 删除Elasticsearch数据
//  @param name 索引名
//  @param esId 索引ID
//  @return err
//  @return ret
//
func (e *ElasticsearchService) DeleteElasticsearch(name string, esId string) (err error, ret interface{}) {

	get, err := global.Elasticsearch.Delete().Index(name).Id(esId).Do(context.Background())
	if err != nil {
		return err, nil
	}
	return err, get
}

//
// UpdateElasticsearch
//  @Description:
//  @param name 索引名
//  @param esId 索引ID
//  @param data 数据
//  @return err 错误
//  @return ret 返回值
//
func (e *ElasticsearchService) UpdateElasticsearch(name string, esId string, data interface{}) (err error, ret interface{}) {

	get, err := global.Elasticsearch.Update().Index(name).Id(esId).Doc(data).Do(context.Background())
	if err != nil {
		return err, nil
	}
	return err, get
}

//
// GetIdElasticsearch
//  @Description: 通过id查找Elasticsearch
//  @param name 索引名
//  @param esId 索引ID
//  @return err 错误
//  @return ret 返回值
//
func (e *ElasticsearchService) GetIdElasticsearch(name string, esId string) (err error, ret interface{}) {
	get, err := global.Elasticsearch.Get().Index(name).Id(esId).Do(context.Background())
	if err != nil {
		return err, nil
	}
	return err, get
}

//
// GetCountElasticsearch
//  @Description: 当前ID索引数量
//  @param name 索引名
//  @return ret 返回值
//
func (e *ElasticsearchService) GetCountElasticsearch(name string) (ret interface{}) {
	list, err := global.Elasticsearch.Count(name).Do(context.Background())
	if err != nil {
		return err
	}
	return list
}

//
// GetQueryElasticsearch
//  @Description: 查询Elasticsearch
//  @param info 查询条件
//  @param name 索引名
//  @param text 查询字段
//  @return err 错误
//  @return ret 返回值
//
func (e *ElasticsearchService) GetQueryElasticsearch(info response.ElasticSearchSearch, name, searchField string) (err error, ret interface{}) {
	size := info.PageSize
	page := info.Page
	//根据name索引查询Elasticsearch数据
	boolQ := elastic.NewQueryStringQuery(info.Title)
	boolQ = boolQ.Field(searchField)
	get, err := global.Elasticsearch.Search(name).
		Query(boolQ). // specify the query
		//Sort("id", true). //按字段"age"排序，升序排列
		Size(size). // 分页，单页显示10条
		From((page - 1) * size).
		//FetchSourceContext(fsc).//只取对应字段
		Do(context.Background()) // 执行
	if err != nil {
		return err, ""
	}
	return err, get
}

//
// GetQueryMultipleElasticsearch
//  @Description: 多字段查询Elasticsearch
//  @param info 查询条件
//  @param name 索引名
//  @param text 查询字段
//  @return err 错误
//  @return ret 返回值
//
func GetQueryMultipleElasticsearch(info response.ElasticSearchSearch, name string, fields ...string) (err error, ret interface{}) {
	size := info.PageSize
	page := info.Page
	if info.Type == 1 {
		//Elasticsearch精准搜索
		//高亮加粗
		hig := elastic.NewHighlight()
		hig = hig.Field("barCode")
		hig = hig.PreTags("<font color='red'>")
		hig = hig.PostTags("</font>")
		get, err := global.Elasticsearch.Search(name).
			Query(elastic.NewMatchPhraseQuery("barCode", info.Title)).
			Highlight(hig).
			Size(size). // 分页，单页显示10条
			From((page - 1) * size).
			Do(context.Background())
		if err != nil {
			return err, ""
		}
		return err, get

	} else {
		// 模糊搜索多字段-需要用分词器analysis-ik
		if len(fields) > 0 {
			get, err := global.Elasticsearch.Search(name).
				Query(elastic.NewBoolQuery().Should(elastic.NewMultiMatchQuery(info.Title, fields).
					Fuzziness("AUTO")).MinimumShouldMatch("1")).
				From((page - 1) * size).Size(size).Do(context.Background())
			if err != nil {
				return err, ""
			}
			return err, get
		}
	}
	return err, nil
}

//
// 	UpdateEsLocation
//  @Description:
//  @param id 比如用户UID
//  @param name 索引名
//  @param latitude 纬度
//  @param longitude 经度
//  @return err 错误
//  @return ret 返回值
//
func (e *ElasticsearchService) UpdateElasticsearchLocation(id int, name, latitude, longitude string) (err error, ret interface{}) {
	// 更新经纬度到es
	data := fmt.Sprintf(`{
    "uid": "%d",
    "location": "%s,%s"
    }`, id, latitude, longitude)
	get, err := global.Elasticsearch.Index().Index(name).Id(strconv.Itoa(int(id))).BodyJson(data).Do(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	return err, get
}

//
// GetElasticsearchNearby
//  @Description: 查询附近的人
//  @param info 查询条件
//  @param name 索引名
//  @param text 查询字段
//  @return err 错误
//  @return ret 返回值
//
func (e *ElasticsearchService) GetElasticsearchNearby(info response.ElasticSearchSearch, name string) (err error, ret interface{}) {
	// 距离范围，默认100
	if info.Distance == "" {
		distance = "100"
	}
	// 分页
	size := info.PageSize
	page := info.Page
	from := (page - 1) * size
	//查询附近的人的位置
	query := elastic.NewGeoDistanceQuery("location").Distance(distance + "km").Lat(info.Latitude).Lon(info.Longitude)
	sort := elastic.NewGeoDistanceSort("location").Point(info.Latitude, info.Longitude).Asc().DistanceType("arc").Unit("km")
	get, err := global.Elasticsearch.Search(name).
		Query(query).SortBy(sort).
		From(from).Size(size).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		return err, nil
	}
	//下面拿到查询结果，可以循环es结果去数据库根据UID或者传入商户ID查询你自己的数据
	return err, get
}
