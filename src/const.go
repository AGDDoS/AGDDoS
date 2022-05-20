package main

// 定义各种常量、变量等
const (
	WelcomeMsg = "" +
		"+----------------------------------------------------------------+ \n" +
		"| Welcome to use AGDDOS                                          | \n" +
		"| Code by AGDDoS Team                                            | \n" +
		"| If you have some problem when you use the tool,                | \n" +
		"| please submit issue at : https://github.com/AGDDoS/AGDDoS      | \n" +
		"+----------------------------------------------------------------+"
	debugHelpMsg = "AGDDoS Debug Command Line Tool\n" +
		"===============================\n" +
		"SYNTAX\n" +
		"    ./AGDDoS debug [--help|-h]\n" +
		"    ./AGDDoS debug [URL<string>]\n" +
		"REMARKS\n" +
		"    URL is an optional parameter\n" +
		"    We debug https://github.com by default\n" +
		"    If you want to debug another URL, enter URL param\n" +
		"EXAMPLE\n" +
		"    ./AGDDoS debug\n" +
		"    ./AGDDoS debug https://fastgit.org"
)

var (
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
		"AGDDoS/" + version + "(" + timestamp + ")",
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
	Debug = false
)
