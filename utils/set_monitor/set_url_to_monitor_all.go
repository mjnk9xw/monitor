package set_monitor

import (
	"context"
	"fmt"

	"go.etcd.io/etcd/clientv3"
)

func SetURLToMonitorAll(address, name, url string) {
	cfg := clientv3.Config{Endpoints: []string{address}}
	c, err := clientv3.New(cfg)
	defer c.Close()
	if err != nil {
		fmt.Println(c)
		fmt.Println("err: ", err)
	}

	fmt.Println("Check 1")
	kapi := clientv3.NewKV(c)
	fmt.Println("Check 2")
	resp, err := kapi.Put(context.Background(), name, url)
	fmt.Println(resp, err)
	return
}
