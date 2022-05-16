package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

type ElasticSearchSearch struct {
	Title string `json:"title" form:"title"`
	request.PageInfo
}
