package utils

import sync2 "sync"

var rate map[string]float64

func GetRate() map[string]float64 {
	if rate == nil {
		return make(map[string]float64)
	}

	return rate
}

func SetRate(res map[string]float64, mu *sync2.RWMutex) {
	mu.Lock()
	defer mu.Unlock()
	if rate == nil {
		rate = make(map[string]float64)
	}

	for k, v := range res {
		rate[k] = v
	}
}
