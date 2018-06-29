package wallet

import "sync"

type nonceGuard *struct {
	lock  *sync.Mutex
	nonce uint64
}

var nonceMgr = make(map[string]nonceGuard)
