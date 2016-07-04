package main

import mcloud "github.com/ldinc/mcloudru-lib"
import (
	"fmt"
)

func PrintEndpoints() {
	fmt.Println("> endpoints:")
	fmt.Println("domain: " + mcloud.Domain)
	fmt.Println("cloud: " + mcloud.Cloud)
	fmt.Println("auth: " + mcloud.Auth)
	fmt.Println("pub: " + mcloud.Public)
}

func PrintMimes() {
	fmt.Println("> mime:")
	fmt.Println(mcloud.GetContentType(".asm"))
	fmt.Println(mcloud.GetContentType(".txt"))
	fmt.Println(mcloud.GetContentType(".m4a"))
	fmt.Println(mcloud.GetContentType(".alac"))
	fmt.Println(mcloud.GetContentType(".cs"))
}

func main() {
	fmt.Println("testing client ...")
	PrintEndpoints()
	PrintMimes()
	mcloud.Hello()
}
