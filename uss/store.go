package uss

import (
	"fmt"
	"sync"
	"sync/atomic"
	"github.com/itchyny/base58-go"
)

// using an in-memory map and counter for POC
var urlmap sync.Map
var ctr atomic.Int64

func Init() {
	ctr.Store(10000)
}

func Save(url string) string {
	shortId := generateShortId(ctr.Add(1))
	urlmap.Store(shortId, url)
	return shortId
}

func Get(shortId string) (string, bool) {
	url, ok := urlmap.Load(shortId)
	return url.(string), ok
}

func Delete(shortId string) {
	urlmap.Delete(shortId)
}

func generateShortId(n int64) string {
	enc := base58.BitcoinEncoding
    s, err := enc.Encode([]byte(fmt.Sprintf("%d", n)))
    if err != nil {
        panic(err)
    }
    return string(s)
}