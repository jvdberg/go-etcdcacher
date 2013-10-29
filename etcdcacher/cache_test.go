package etcdcacher

import (
	"testing"
	"github.com/coreos/go-etcd/etcd"
	"time"
)

func TestCache(t *testing.T) {
	c := etcd.NewClient([]string{"http://127.0.0.1:4001"})

	c.Set("/test/a","test",0)

	cache := NewCache(c,"/test/")
	
	result, err := cache.Get("/test/a")
	if err != nil || result[0].Key != "/test/a" || result[0].Value != "test" {
		if err != nil {
			t.Fatal(err)
		}
		t.Fatalf("Watch failed with %s %s %v %v", result[0].Key, result[0].Value, result[0].TTL, result[0].Index)
	}

	c.Set("/test/a","test2",0)
	time.Sleep(1* time.Second)

	result, err = cache.Get("/test/a")
	if err != nil || result[0].Key != "/test/a" || result[0].Value != "test2" {
		if err != nil {
			t.Fatal(err)
		}
		t.Fatalf("Watch2 failed with %s %s %v %v", result[0].Key, result[0].Value, result[0].TTL, result[0].Index)
	}

	c.Delete("/test/a")
	time.Sleep(1* time.Second)

	result, err = cache.Get("/test/a")
	if err.Error() != "101: Key not found" {
		t.Fatal("Should throw error 101: Key not found")
	}
}
