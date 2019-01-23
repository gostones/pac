package main

import (
	"flag"
	"fmt"

	"github.com/gostones/pac"
	"os"
)

// pac tries set the proxy on networking settings.
// For now, this feature is only supported for:
//
//   - OSX
//   - Linux-gnome
//   - Windows
//
// Usage:
//
// `pac`
//

var usage = `
Supported systems:

-linux
-darwin
-windows

`

func main() {
	proxyURL := flag.String("proxy", "", "proxy auto config url")
	if *proxyURL == "" {
		fmt.Println(usage)
		os.Exit(1)
	}
	err := pac.SetupProxy(*proxyURL, false)
	fmt.Println(err)
}
