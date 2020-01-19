package pool

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/TT527/ProxyIP/cache"
	"github.com/TT527/ProxyIP/util"
	"github.com/go-redis/redis"
	"net/http"
	"strconv"
	"sync"
)

type remotePool struct {
	redis *redis.Client
}

type RedisClient struct {
	addr string
	port string
	pass string
	db   int
}

type proxyIP struct {
	Addr string `json:"addr"`
	Port string `json:"port"`
	Area string `json:"area"`
	Time string `json:"time"`
}

func (r *remotePool) saveRedis(ip string, hBody string) error {

	err := r.redis.HSet("ippool", ip, hBody).Err()
	if err != nil {
		return err
	}

	err = r.redis.SAdd("ippoolkey", ip).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *remotePool) GetPools() int {

	var wg sync.WaitGroup
	var count int
	for i := 1; i < 80; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			request, err := http.NewRequest("GET", fmt.Sprintf(util.PoolUrl, strconv.Itoa(i)), nil)
			if err != nil {
				fmt.Println("err1--Page_"+strconv.Itoa(i)+":", err)
				return
			}
			request.Header.Set("User-Agent", util.RandUA())

			cli1 := &http.Client{}
			res, err := cli1.Do(request)

			if err != nil {
				fmt.Println("err2--Page_"+strconv.Itoa(i)+":", err)
				//切换ip获取

				return
			}

			doc, err := goquery.NewDocumentFromReader(res.Body)
			if err != nil {
				fmt.Println("err3--Page_"+strconv.Itoa(i)+":", err)
				return
			}
			doc.Find("tbody").Each(func(i int, s *goquery.Selection) {

				s.Find("tr").Each(func(i int, s2 *goquery.Selection) {
					addr := util.CompressStr(s2.Find("td:first-child").Text())
					port := util.CompressStr(s2.Find("td:nth-child(2)").Text())
					area := util.CompressStr(s2.Find("td:nth-child(3)").Text())
					time := util.CompressStr(s2.Find("td:nth-child(5)").Text())

					info := proxyIP{
						addr,
						port,
						area,
						time,
					}
					hBody, _ := json.Marshal(info)

					err := r.saveRedis(addr+":"+port, string(hBody))
					if err != nil {
						fmt.Println(err)
						return
					}
					count++

				})

			})

		}(i)
	}

	wg.Wait()

	return count
}

func NewPool(r RedisClient) (Pool, error) {

	initRedis, err := cache.InitRedis(r.addr+":"+r.port, r.pass, r.db)
	if err != nil {
		return nil, err
	}
	p := &remotePool{
		redis: initRedis,
	}
	return p, nil
}
