package main

import mcloud "github.com/ldinc/mcloudru-lib"
import (
	"fmt"
)

func PrintEndpoints() {
	fmt.Println("endpoints:")
	fmt.Println("domain: " + mcloud.Domain)
	fmt.Println("cloud: " + mcloud.Cloud)
	fmt.Println("auth: " + mcloud.Auth)
	fmt.Println("pub: " + mcloud.Public)
}

func main() {
	fmt.Println("testing client ...")
	PrintEndpoints()
	mcloud.Hello()
}
