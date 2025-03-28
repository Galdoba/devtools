package main

import "fmt"

func main() {
	var version string
	version = "0.0.0 [build 4]" //#gvc: version control token
	fmt.Printf("I am Hello App with version: '%v'", version)
}
