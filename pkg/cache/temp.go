package cache

import (
	"github.com/defany/govk/pkg/maps"
	"time"
)

const maxExtraLifeIncrease = 3

type Temp[V any] struct {
	stop bool

	saveTime time.Duration

	expirationTimes map[string]time.Time
	values          map[string]V

	// isActiveLonger
	//
	// This param means, that when key will be got from cache his expiration will be increased to make life longer
	// but not bigger than maxExtraLifeIncrease X from save time
	isActiveLonger bool
}

func NewTemp[V any](saveTime time.Duration, isActiveLonger bool) *Temp[V] {
	temp := &Temp[V]{
		saveTime:        saveTime,
		expirationTimes: make(map[string]time.Time),
		values:          make(map[string]V),
		isActiveLonger:  isActiveLonger,
	}

	temp.startRefresh()

	return temp
}

func (t *Temp[V]) Save(key string, value V) {
	t.values[key] = value

	t.expirationTimes[key] = time.Now().Add(t.saveTime)
}

func (t *Temp[V]) Get(key string) (V, bool) {
	v, ok := t.values[key]
	if !ok {
		return v, false
	}

	if t.isActiveLonger {
		if t.expirationTimes[key].UnixMilli()-time.Now().UnixMilli() < t.saveTime.Milliseconds()*maxExtraLifeIncrease {
			t.expirationTimes[key] = t.expirationTimes[key].Add(t.saveTime)
		}
	}

	return v, true
}

func (t *Temp[V]) Stop() {
	t.stop = true
}

func (t *Temp[V]) startRefresh() {
	ticker := time.Tick(t.saveTime)

	go func(ticker <-chan time.Time) {
		for {
			select {
			case <-ticker:
				t.clearOld()
			default:
				if !t.stop {
					continue
				}

				t.clear()
			}
		}
	}(ticker)
}

func (t *Temp[V]) clear() {
	t.values = nil
	t.expirationTimes = nil
}

func (t *Temp[V]) clearOld() {
	for key := range t.values {
		createdAt, ok := t.expirationTimes[key]
		if !ok {
			t.delete(key)

			continue
		}

		if time.Now().Unix() >= createdAt.Unix() {
			continue
		}

		t.delete(key)
	}
}

func (t *Temp[V]) delete(key string) {
	defer func() {
		t.values = maps.Copied(t.values)
		t.expirationTimes = maps.Copied(t.expirationTimes)
	}()

	delete(t.values, key)
	delete(t.expirationTimes, key)
}
