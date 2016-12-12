package cache

import (
    "github.com/coreos/etcd/clientv3"
)

type EtcdCache struct {

}

func NewEtcdCache(clientv3 clientv3.Client, prefix string) *Cache {

}
