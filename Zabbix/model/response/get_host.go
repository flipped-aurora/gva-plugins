package response

type ZabbixLoginModel struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"params"`
	ID   int         `json:"id"`
	Auth interface{} `json:"auth"`
}

type GetHostModel struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Output                []string `json:"output"`
		SelectInterfaces      []string `json:"selectInterfaces"`
		SelectGroups          []string `json:"selectGroups"`
		SelectParentTemplates []string `json:"selectParentTemplates"`
	} `json:"params"`
	ID   int    `json:"id"`
	Auth string `json:"auth"`
}

type ResultModel struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  string `json:"result"`
	ID      int    `json:"id"`
}
