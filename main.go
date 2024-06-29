package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Shobhit-Nagpal/html-lsp/internal/rpc"
)

func main() {
  //Set up scanner for reading input through stdin
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(rpc.Split)

  //Make RPC connection
  fmt.Println("LSP server started")
  for scanner.Scan() {
    chunk := scanner.Text()
    fmt.Println(chunk)
  }
}
