package request

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/parnurzeal/gorequest"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var UserAgentLists = []string{
	"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; AcooBrowser; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0; Acoo Browser; SLCC1; .NET CLR 2.0.50727; Media Center PC 5.0; .NET CLR 3.0.04506)",
	"Mozilla/4.0 (compatible; MSIE 7.0; AOL 9.5; AOLBuild 4337.35; Windows NT 5.1; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
	"Mozilla/5.0 (Windows; U; MSIE 9.0; Windows NT 9.0; en-US)",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Win64; x64; Trident/5.0; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 2.0.50727; Media Center PC 6.0)",
	"Mozilla/5.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0; WOW64; Trident/4.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 1.0.3705; .NET CLR 1.1.4322)",
	"Mozilla/4.0 (compatible; MSIE 7.0b; Windows NT 5.2; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.2; .NET CLR 3.0.04506.30)",
	"Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN) AppleWebKit/523.15 (KHTML, like Gecko, Safari/419.3) Arora/0.3 (Change: 287 c9dfb30)",
	"Mozilla/5.0 (X11; U; Linux; en-US) AppleWebKit/527+ (KHTML, like Gecko, Safari/419.3) Arora/0.6",
	"Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.2pre) Gecko/20070215 K-Ninja/2.1.1",
	"Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; rv:1.9) Gecko/20080705 Firefox/3.0 Kapiko/3.0",
	"Mozilla/5.0 (X11; Linux i686; U;) Gecko/20070322 Kazehakase/0.4.5",
	"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.0.8) Gecko Fedora/1.9.0.8-1.fc10 Kazehakase/0.5.6",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_3) AppleWebKit/535.20 (KHTML, like Gecko) Chrome/19.0.1036.7 Safari/535.20",
	"Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; fr) Presto/2.9.168 Version/11.52",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.11 (KHTML, like Gecko) Chrome/20.0.1132.11 TaoBrowser/2.0 Safari/536.11",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/21.0.1180.71 Safari/537.1 LBBROWSER",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E; LBBROWSER)",
	"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; QQDownload 732; .NET4.0C; .NET4.0E; LBBROWSER)",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.84 Safari/535.11 LBBROWSER",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E)",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E; QQBrowser/7.0.3698.400)",
	"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; QQDownload 732; .NET4.0C; .NET4.0E)",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Trident/4.0; SV1; QQDownload 732; .NET4.0C; .NET4.0E; 360SE)",
	"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; QQDownload 732; .NET4.0C; .NET4.0E)",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E)",
	"Mozilla/5.0 (Windows NT 5.1) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/21.0.1180.89 Safari/537.1",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/21.0.1180.89 Safari/537.1",
	"Mozilla/5.0 (iPad; U; CPU OS 4_2_1 like Mac OS X; zh-cn) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8C148 Safari/6533.18.5",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:2.0b13pre) Gecko/20110307 Firefox/4.0b13pre",
	"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:16.0) Gecko/20100101 Firefox/16.0",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11",
	"Mozilla/5.0 (X11; U; Linux x86_64; zh-CN; rv:1.9.2.10) Gecko/20100922 Ubuntu/10.10 (maverick) Firefox/3.6.10",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
}
var UserProxyLists = []string{
	"",
	"112.85.129.175:9999",
	"112.85.130.75:9999",
	"110.52.235.198:9999",
	"110.52.235.182:9999",
	"112.85.169.72:9999",
	"219.138.47.221:9999",
	"122.193.244.244:9999",
	"125.126.206.156:9999",
	"125.125.211.84:9999",
	"110.52.235.136:9999",
	"110.52.235.211:9999",
	"223.241.78.195:9999",
	"112.85.128.211:9999",
	"112.85.168.233:9999",
	"117.91.232.196:9999",
	"175.148.73.240:1133",
	"111.177.107.3:9999",
	"110.52.235.64:9999",
	"110.52.235.121:9999",
	"125.126.199.168:9999",
	"221.233.46.216:9999",
	"110.52.235.90:9999",
	"183.148.135.37:9999",
	"112.85.173.61:9999",
	"112.85.170.91:9999",
	"112.87.68.35:9999",
	"110.52.235.202:9999",
	"218.87.239.121:9999",
	"112.85.130.3:9999",
	"119.99.45.123:9999",
	"116.209.55.141:9999",
	"112.85.130.212:9999",
	"110.52.235.214:9999",
	"119.99.45.40:9999",
	"110.52.235.246:9999",
	"115.53.22.4:9999",
	"125.126.223.56:9999",
	"110.52.235.78:9999",
	"111.177.183.235:9999",
	"117.81.136.85:9999",
	"111.177.177.68:9999",
	"183.148.143.254:9999",
	"111.177.187.172:9999",
	"183.148.135.204:9999",
	"111.177.162.241:9999",
	"117.62.38.225:9999",
	"114.230.69.14:9999",
	"110.52.235.3:9999",
	"163.204.241.9:9999",
	"114.139.32.27:9999",
	"223.241.116.196:9999",
	"117.91.255.59:9999",
	"61.235.24.126:8123",
}

type Requester struct {
	Client *ethclient.Client
}

func Initialize(url string) (*Requester, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("initialize eth client err:%v", err))
	}

	return &Requester{Client: client}, nil
}

func (r *Requester) RequestContract(url string) (full, code, abi, bin string, err error) {
	fmt.Println("request url:", url)
	res := gorequest.New().Proxy(UserProxyLists[rand.Intn(len(UserProxyLists))]).Set("user-agent", UserAgentLists[rand.Intn(len(UserAgentLists))])
	ret, body, errs := res.Timeout(time.Second * 5).Retry(5,time.Second,http.StatusRequestTimeout,http.StatusBadRequest).Get(url).End()
	if errs != nil || ret.StatusCode != http.StatusOK {
		req, err := res.MakeRequest()
		if err == nil && req != nil && ret != nil {
			fmt.Printf("request status:%d, body:%+v\n", ret.StatusCode, req)
		}
		var errStr string
		for _, e := range errs {
			errStr += e.Error()
		}
		return "", "", "", "", errors.New(errStr)
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		return "", "", "", "", err
	}
	codeSelector := doc.Find("#dividcode")
	code = codeSelector.Find("pre[class='js-sourcecopyarea editor']").Text()
	abi = codeSelector.Find("pre[class='wordwrap js-copytextarea2']").Text()
	bin = codeSelector.Find("#verifiedbytecode2").Text()

	return codeSelector.Text(), code, abi, bin, nil
}

func GetIpList() ([]string, error) {
	res := gorequest.New().Set("user-agent", UserAgentLists[rand.Intn(len(UserAgentLists))])
	ret, body, errs := res.Timeout(time.Second * 5).Get("https://www.xicidaili.com/nn/7").End()
	if errs != nil || ret.StatusCode != http.StatusOK {
		req, err := res.MakeRequest()
		if err == nil {
			fmt.Printf("request body:%+v\n", req)
		}
		var errStr string
		for _, e := range errs {
			errStr += e.Error()
		}
		return nil, errors.New(errStr)
	}
	fmt.Println("body:", body)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		fmt.Println("new doc err:", err)
		return nil, err
	}
	doc.Find("div").Each(func(i int, selection *goquery.Selection) {
		fmt.Println("td:", selection.Text())
	})

	return nil, nil
}
