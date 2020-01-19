package pool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func GetTest()  {
	proxyAddr := "https://49.70.33.169:9999"
	httpUrl := "http://106.12.221.167:8085/"
	proxy, err := url.Parse(proxyAddr)
	if err != nil {
		log.Fatal(err)
	}
	netTransport := &http.Transport{
		Proxy:http.ProxyURL(proxy),
		MaxIdleConnsPerHost: 10,
		ResponseHeaderTimeout: time.Second * time.Duration(5),
	}
	httpClient := &http.Client{
		Timeout: time.Second * 10,
		Transport: netTransport,
	}
	res, err := httpClient.Get(httpUrl)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Println(err)
		return
	}
	c, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(c))
}
func Hash()  {

	// 创建一个结构体对象
	person := Person{"小明", 18}
	result, err := json.Marshal(&person)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result))

}

type Person struct {
	// 这里的两个字段名，首字母都要大写，否则无法转换
	name string
	age  int
}

func main() {

}
