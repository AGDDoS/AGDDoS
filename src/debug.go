package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
)

type DebugFunc struct {
}

func (d *DebugFunc) Run(args []string) {
	var isConnectOk bool
	switch len(args) {
	case 0:
		isConnectOk = debug("https://github.com")
	case 1:
		os.Exit(0)
	default:
		fmt.Println("Invalid args for debug. Use -h to get more information.")
	}
	if isConnectOk {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

func debugConnection(url string) bool {
	fmt.Print("Test connection...")
	response, err := http.Head(url)
	if err != nil {
		fmt.Println("Failed")
		fmt.Println("Response create failed\n", err)
		return false
	}
	if response.StatusCode != http.StatusOK {
		fmt.Println("Failed")
		return false
	} else {
		fmt.Println("Success")
		return true
	}
}

func debug(url string) bool {
	if url != "--help" && url != "-h" {
		fmt.Println(
			"AGDDoS Debug Command Line Tool\n" +
				"===============================")
		if !(strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) {
			url = "http://" + url
		}
		fmt.Println("Remote Address:", url)
		fmt.Print("IP Address: ")

		addr, err := net.LookupIP(removeHttpAndHttps(url))

		if err != nil {
			fmt.Println("Unknown")
		} else {
			fmt.Println(addr)
		}

		fmt.Print("Local Address: ")
		resp, err := http.Get("https://api.pig2333.workers.dev/ip") // Thanks the api by @xiaozhu2007

		if err != nil {
			fmt.Println("Unknown ->", err)
			defer resp.Body.Close()
		} else {
			s, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Unknown ->", err)
				defer resp.Body.Close()
			} else {
				fmt.Printf("[%s]\n", strings.Replace(string(s), "\n", "", -1))
			}
		}
		return debugConnection(url)
	} else {
		fmt.Println(debugHelpMsg)
		return true
	}
}
