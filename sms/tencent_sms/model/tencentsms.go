package model

type TencentModel struct {
	Phones        []string `json:"phones"`
	TemplateId    string   `json:"templateId"`
	TemplateParam []string `json:"templateParam"`
}
