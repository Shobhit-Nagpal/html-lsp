package rpc

import (
	"bytes"
	"fmt"
)

func Split(data []byte, _ bool) (advance int, token []byte, err error) {
  content := bytes.Split(data, []byte("\r\n\r\n"))
  fmt.Println(content)
  return 0, data, nil
}
