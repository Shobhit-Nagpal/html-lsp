package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func Split(data []byte, _ bool) (advance int, token []byte, err error) {
	header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return 0, nil, nil
	}
	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return 0, nil, err
	}

	if len(content) < contentLength {
		return 0, nil, nil
	}

	totalLength := len(header) + 4 + contentLength
	return totalLength, data[:totalLength], nil
}

func Encode(resp interface{}) string {
	msg, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(msg), msg)
}

func Decode(data []byte) (string, []byte, error) {
	header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return "", nil, errors.New("Couldn't find separator")
	}
	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return "", nil, err
	}

	var dat RequestMessage
	err = json.Unmarshal(content[:contentLength], &dat)
	if err != nil {
		return "", nil, err
	}

	return dat.Method, content[:contentLength], nil
}

func HandleRequest(method string, content []byte, logger *log.Logger) {
	switch method {
	case "initialize":
		var req InitializeRequest
		err := json.Unmarshal(content, &req)
		if err != nil {
			logger.Printf("Error parsing initialize request: %s\n", err)
		}
		logger.Print(method)
		resp := NewInitializeResponse(req.Id)
		writeResponse(os.Stdout, resp)
		logger.Print("Responded to initialize request")
	}
}

func NewInitializeResponse(id int) InitializeResponse {
	return InitializeResponse{
		ResponseMessage{
			Jsonrpc: "2.0",
			Id:      &id,
		},
		InitializeResult{
			Capabilities: ServerCapabilities{
				TextDocumentSync: 1,
			},
			ServerInfo: &ServerInfo{
				Name:    "lsp",
				Version: "0.0.0.0-beta",
			},
		},
	}
}

func writeResponse(writer io.Writer, resp interface{}) {
	msg := Encode(resp)
	writer.Write([]byte(msg))
}
