package etcdcacher

import (
	"github.com/coreos/go-etcd/etcd"
	"errors"
	"sync"
)

type Cache struct {
	sync.RWMutex
	client *etcd.Client
	path string
	items map[string]*etcd.Response
	
}

func NewCache(client *etcd.Client, path string) *Cache {
	stopChan := make(chan bool)
	receiverChan := make(chan *etcd.Response)
	go client.Watch(path, 0, receiverChan, stopChan)

	response, _ := client.Get(path)
	cachedItems :=  map[string]*etcd.Response{ }
	for _,element := range response {
		cachedItems[element.Key] = element
	}
	
	cache := &Cache{
		client: client,
		path: path,
		items: cachedItems,
	}

	go cache.Watcher(receiverChan, stopChan)

	return cache
}

func (c *Cache) Watcher(receiver chan *etcd.Response, stop chan bool) {
	for {
		resp := <- receiver
		c.Lock()
		if(resp.Action=="DELETE") {
			delete(c.items,resp.Key)
		} else {
			c.items[resp.Key] = resp
		}
		c.Unlock()
	}
}

func (c *Cache) Get(key string) ([]*etcd.Response, error) {
	c.RLock()
	data, ok := c.items[key]
	c.RUnlock()
	if(ok) {
		return []*etcd.Response{data}, nil
	} else {
		return nil, errors.New("101: Key not found")
	}
}