package main

import "fmt"

func Log(grindex int, i int, responseStatus string) {
	fmt.Printf("[INFO][AGDDoS#%d/%d]%s \033[0m \n", grindex, i, responseStatus) // 默认 [INFO][AGDDoS#当前线程序号/第i次重复] (Get "https://1.1.1.1": dial tcp 1.1.1.1:443: i/o timeout)
}

func PrintError(grindex int, i int, responseStatus string) {
	//fmt.Printf("[Error][AGDDoS#%d/%d] \033[1;31;40m (%s) \033[0m \n", grindex, i, responseStatus) // 默认
}
