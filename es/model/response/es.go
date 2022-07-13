package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

type ElasticSearchSearch struct {
	Title     string `json:"title" form:"title"`
	Distance  string `json:"distance" form:"distance"`
	Latitude  string `json:"latitude" form:"latitude"`
	Longitude string `json:"longitude" form:"longitude"`
	Type      uint   `json:"type" form:"type"`
	request.PageInfo
}
