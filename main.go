package main

import (
	"api-openlive/config"
	"flag"
	"fmt"
	"os"
)

var Environment *string

func init() {
	Environment = flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
}

func main() {
	CronRemoveTemp()
	config.Init(*Environment)
	Init()
}
