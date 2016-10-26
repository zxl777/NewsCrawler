package main

/*
brew services restart selenium-server-standalone
brew services restart chromedriver

利用Chrome的查看功能，通过FindElements找到内容列表，

然后用FindElement在某一个内容区块中，继续搜索文字信息，图片信息。

用 selenium img 搜索，能获得更多关于怎么获得img的参考信息。
*/
import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"

	redis "gopkg.in/redis.v5"

	"github.com/tebeka/selenium"
)

var client *redis.Client

//GetMD5 获得字符串的md5
func GetMD5(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func initRedis() {
	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}

//ParseTweet 解析文本，输出tweet字符串，like数，转发数，回复数
func ParseTweet(txt string) (string, int) {
	s := strings.Split(txt, "\n")
	likes, _ := strconv.Atoi(s[len(s)-2])
	// retweets, _ := strconv.Atoi(s[len(s)-4])
	// replies, _ := strconv.Atoi(s[len(s)-6])

	// if err != nil {
	// 	return "", 0, 0, 0
	// }

	return s[len(s)-9], likes
}

func main6() {
	// FireFox driver without specific version
	// *** Add gecko driver here if necessary (see notes above.) ***
	// caps := selenium.Capabilities{"browserName": "firefox"}
	caps := selenium.Capabilities{"browserName": "chrome"}

	wd, err := selenium.NewRemote(caps, "")
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// Get simple playground interface
	wd.Get("https://twitter.com/search?q=minecraft&src=typd")

	// tweets, _ := wd.FindElements(selenium.ByCSSSelector, ".AdaptiveMedia-photoContainer") //(selenium.ByCSSSelector, ".content")
	// fmt.Printf("Got: %s\n", tweets)
	// tweets, _ := wd.FindElements(selenium.ByTagName, "img") //(selenium.ByCSSSelector, ".content")

	// tweets, _ := wd.FindElements(selenium.ByCSSSelector, ".AdaptiveMedia.is-square")

	// for i, v := range tweets {
	// 	// tt, _ := v.Text()

	// 	tt, _ := v.GetAttribute("src")

	// 	fmt.Println("tweets[", i, "] =", tt)
	// }

	//#stream-item-tweet-789452502747852805 > div > div.content > div.AdaptiveMedia.is-square > div > div > div > img

	// body, _ := wd.FindElement(selenium.ByTagName, "body")
	// body.SendKeys(selenium.PageDownKey)
	// time.Sleep(time.Millisecond * 1000)
	// body.SendKeys(selenium.PageDownKey)
	// time.Sleep(time.Millisecond * 1000)
	// body.SendKeys(selenium.PageDownKey)
	// time.Sleep(time.Millisecond * 1000)
	// body.SendKeys(selenium.PageDownKey)
	// time.Sleep(time.Millisecond * 1000)
	// body.SendKeys(selenium.PageDownKey)
	// time.Sleep(time.Millisecond * 1000)
	// body.SendKeys(selenium.PageDownKey)
	// time.Sleep(time.Millisecond * 1000)
	// body.SendKeys(selenium.PageDownKey)
	// time.Sleep(time.Millisecond * 1000)
	// body.SendKeys(selenium.PageDownKey)
	// time.Sleep(time.Millisecond * 1000)
	// body.SendKeys(selenium.PageDownKey)
	// time.Sleep(time.Millisecond * 1000)
	// body.SendKeys(selenium.PageDownKey)
	// time.Sleep(time.Millisecond * 1000)
	// body.SendKeys(selenium.PageDownKey)
	// time.Sleep(time.Millisecond * 1000)
	//<div class="content">
	tweets, _ := wd.FindElements(selenium.ByCSSSelector, ".content")

	fmt.Printf("tweets 条数 ＝ %d\n", len(tweets))

	for i, v := range tweets {
		txt, _ := v.Text()
		// _ = txt
		// _ = i
		fmt.Println("tweets[", i, "] =", txt)

		v0, _ := v.FindElement(selenium.ByCSSSelector, ".AdaptiveMedia.is-square")

		if v0 != nil {
			/*
				<img data-aria-label-part="" src="https://pbs.twimg.com/media/CvXwPVAUIAASnAB.jpg" alt="" style="width: 100%; top: -84px;">
			*/

			v1, _ := v0.FindElement(selenium.ByTagName, "img")
			tt, _ := v1.GetAttribute("src")

			fmt.Println("img  = ", tt)
		}

		//<a href="/nfrance09/status/789801462968623104" class="tweet-timestamp js-permalink js-nav js-tooltip" data-original-title="5:11 AM - 22 Oct 2016"><span class="_timestamp js-short-timestamp js-relative-timestamp" data-time="1477138317" data-time-ms="1477138317000" data-long-form="true" aria-hidden="true">2h</span><span class="u-hiddenVisually" data-aria-label-part="last">2 hours ago</span></a>
		//<span class="_timestamp js-short-timestamp js-relative-timestamp" data-time="1477138317" data-time-ms="1477138317000" data-long-form="true" aria-hidden="true">2h</span>
		timeline, _ := v.FindElement(selenium.ByCSSSelector, "._timestamp")
		timestamp, _ := timeline.GetAttribute("data-time")
		fmt.Println("时间戳：", timestamp)

		urlline, _ := v.FindElement(selenium.ByCSSSelector, "a[href*='status']")
		url, _ := urlline.GetAttribute("href")
		fmt.Println("源地址：", url)

		fmt.Printf("-------------------\n")
		//
		//
		//
	}

	// Get the result #stream-item-tweet-789304486896427009 > div > div.content > div.js-tweet-text-container
	// div, _ := wd.FindElement(selenium.ByCSSSelector, ".content") //".js-tweet-text-container")
	// output, _ := div.Text()
	// // oo, _ := div.GetAttribute
	// fmt.Printf("Got: %s\n", output)

	// divv, _ := wd.FindElement(selenium.ByClassName, "AdaptiveMedia-photoContainer js-adaptive-photo ")
	// outputv, _ := divv.Text()
	// // oo, _ := div.GetAttribute
	// fmt.Printf("Got: %s\n", outputv)

	fmt.Printf("--- END ---")
}

/*
定时抓取20条 twitter
只处理like数超过5的tweet
用文字的md5做id
写入hmset
写入zset，排序值算法（新的排序值高，时间戳）（先用时间戳，升级版再用reddit算法，需要增加两个用户投票按钮，相应响应程序）
*/

func timer4Twitter() {
	crawlTwitter()
	c := time.Tick(200 * time.Second)
	for now := range c {
		fmt.Println("tick", now)
		crawlTwitter()
	}
}

func crawlTwitter() {
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, "")
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	wd.Get("https://twitter.com/search?q=minecraft&src=typd")
	// wd.Get("https://twitter.com/zxl777")
	// wd.Get("https://twitter.com/search?q=minecraft%20Wallpaper&src=typd")

	timeline, _ := wd.FindElement(selenium.ByCSSSelector, "#timeline") //只在timeline－content中找tweets
	tweets, _ := timeline.FindElements(selenium.ByCSSSelector, ".content")

	fmt.Printf("tweets 条数 ＝ %d\n", len(tweets))

	for i, v := range tweets { //遍历搜索到的20条结果
		fmt.Printf("Tweet [%d]\n", i)

		//获取时间戳
		timeline, _ := v.FindElement(selenium.ByCSSSelector, "._timestamp")
		timestamp, _ := timeline.GetAttribute("data-time")
		fmt.Println("时间戳：", timestamp)

		//获取源地址
		urlline, _ := v.FindElement(selenium.ByCSSSelector, "a[href*='status']")
		url, _ := urlline.GetAttribute("href")
		fmt.Println("源地址：", url)

		//获取图片地址
		var img string
		imgline, err := v.FindElement(selenium.ByCSSSelector, ".AdaptiveMedia")
		if err == nil {
			imgline2, err := imgline.FindElement(selenium.ByTagName, "img")
			if err == nil {
				img, _ = imgline2.GetAttribute("src")
				fmt.Println("img  = ", img)
			}
		}

		//获取tweet文字
		tweetline, _ := v.FindElement(selenium.ByCSSSelector, ".js-tweet-text-container")
		tweet, _ := tweetline.Text()
		fmt.Println("Tweet:", tweet)

		// 获取likes数
		likeline, _ := v.FindElement(selenium.ByCSSSelector, ".ProfileTweet-action.ProfileTweet-action--favorite.js-toggleState")
		likeline2, _ := likeline.FindElement(selenium.ByCSSSelector, ".IconTextContainer")
		likeStr, _ := likeline2.Text()

		likes, _ := strconv.Atoi(likeStr)

		fmt.Println("Likes:", likes)

		if likes > 5 && img != "" {
			fmt.Println("hash~", "news:"+GetMD5(tweet))
			client.HMSet("news:"+GetMD5(tweet), map[string]string{
				"tweet":     tweet,
				"img":       img,
				"url":       url,
				"timestamp": timestamp,
			})

			timestamp, _ := strconv.Atoi(timestamp)
			news := "news:" + GetMD5(tweet)
			client.ZAdd("MineNews:twitter", redis.Z{Score: float64(timestamp), Member: news})
		}

		fmt.Printf("-------------------\n")
	}

	fmt.Printf("--- END --- 10-26")
}

func main() {
	initRedis()
	timer4Twitter()
}
