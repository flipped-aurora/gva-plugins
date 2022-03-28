package response

type GetHostModel struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Output           []string `json:"output"`
		SelectInterfaces []string `json:"selectInterfaces"`
	} `json:"params"`
	ID   int    `json:"id"`
	Auth string `json:"auth"`
}
