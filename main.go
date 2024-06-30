package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Shobhit-Nagpal/html-lsp/internal/rpc"
)

func main() {
	//Set up scanner for reading input through stdin
	logger := getLogger("/home/shbhtngpl/personal/learn/html-lsp/log.txt")
	logger.Println("LSP server started")
  fmt.Println("LSP server started")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	//Make RPC connection
	for scanner.Scan() {
		chunk := scanner.Bytes()
		method, content, err := rpc.Decode(chunk)
		if err != nil {
			logger.Printf("Got an error: %s", err)
			continue
		}
    rpc.HandleRequest(method, content, logger)
	}
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	return log.New(logfile, "[html-lsp]", log.Ldate|log.Ltime)
}
