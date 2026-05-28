package handlerforpayandsavemoney

import "sync"

type Balance struct {
	mu    sync.Mutex
	Bank  int
	Money int
}
