package main

import (
	"fmt"
	"log"
)

const (
	bannerText = `
 __  __     __     ______
/\ \_\ \   /\ \   /\__  _\
\ \  __ \  \ \ \  \/_/\ \/
 \ \_\ \_\  \ \_\    \ \_\
  \/_/\/_/   \/_/     \/_/
`

	usageText = `
	Usage:
	- url: HTTP server URL to make requests (required)
	- n: Number of requests to make
	- c: Concurrency level`
)

func banner() string { return bannerText[1:] }
func usage() string { return usageText[1:] }

func main() {
	var f flags
	if err := f.parse(); err != nil {
		fmt.Println(usage())
		log.Fatal(err)
	}
	fmt.Println(banner())
	fmt.Printf("Making %d requests to %s with a concurrency level of %d.\n", f.n, f.url, f.c)
	fmt.Println(usage())
}
