package pool

import (
	"fmt"
	"testing"
)

func TestRemotePool_GetPools(t *testing.T) {
	var r remotePool
	r.GetPools()
}
func TestGetTest(t *testing.T) {
	//GetTest()
	client := RedisClient{"192.168.1.115", "6379", "",1}
	pool, err := NewPool(client)
	if err != nil {
		fmt.Println(err)
		return
	}
	count:=pool.GetPools()
	fmt.Println(count)
}

func TestHash(t *testing.T) {
	Hash()
}