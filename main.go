package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/foreveryouyou/angego/cmd/commands/web"
)

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		os.Exit(2)
		return
	}

	switch args[0] {
	case "api":
		fmt.Println("api")
	case "web":
		web.Web()
	case "help":
		fmt.Println("help")
	default:
		fmt.Println("无效参数")
	}
}
