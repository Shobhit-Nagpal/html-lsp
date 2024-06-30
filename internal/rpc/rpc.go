package rpc

type RequestMessage struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Method  string `json:"method"`
}

type ResponseMessage struct {
  Jsonrpc string `json:"jsonrpc"`
	Id     *int           `json:"id"`
	Result *string        `json:"result"`
	Error  *ResponseError `json:"error"`
}

type ResponseError struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    *string `json:"data"`
}

type NotificationMessage struct {
  Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
}
