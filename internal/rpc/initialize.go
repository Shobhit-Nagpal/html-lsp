package rpc

type InitializeRequest struct {
	RequestMessage
  Params  InitializeRequestParams `json:"params"`
}

type InitializeRequestParams struct {
	ProcessId  *int        `json:"processId"`
	ClientInfo *ClientInfo `json:"clientInfo"`
}

type InitializeResponse struct {
  ResponseMessage
  InitializeResult  InitializeResult `json:"result"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   *ServerInfo        `json:"serverInfo"`
}

type ServerCapabilities struct {
  TextDocumentSync  int  `json:"textDocumentSync"`
}

type ServerInfo struct {
	Name    string  `json:"name"`
	Version string `json:"version"`
}

type ClientInfo struct {
	Name    string  `json:"name"`
	Version string `json:"version"`
}
