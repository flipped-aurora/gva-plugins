package response

type AlertTiggerModel struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  Params      `json:"params"`
	Auth    interface{} `json:"auth"`
	ID      int         `json:"id"`
}
type Filter struct {
	Value int `json:"value"`
}
type Params struct {
	Output      []string `json:"output"`
	Filter      Filter   `json:"filter"`
	Sortfield   string   `json:"sortfield"`
	Sortorder   string   `json:"sortorder"`
	SelectHosts string   `json:"selectHosts"`
}
