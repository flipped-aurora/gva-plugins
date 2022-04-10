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
