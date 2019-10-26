package man

import (
	"flag"
	"fmt"
	"os"
)

func Usage() {
	fmt.Fprintf(os.Stderr, `Usage: screenshot [-h]
Options:
`)
	flag.PrintDefaults()
}

func yesOrNo() bool {
	var yn string
	fmt.Scanln(&yn)
	if yn == "y" || yn == "Y" {
		return true
	}
	return false
}

func HelpAsk() {
	fmt.Printf("确定Chrome/Chromium已安装? y or n: ")
	if !yesOrNo() {
		os.Exit(0)
	}
	fmt.Printf("确定VPN已关闭? y or n: ")
	if !yesOrNo() {
		os.Exit(0)
	}
}
