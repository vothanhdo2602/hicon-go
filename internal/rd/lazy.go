package rd

import (
	"context"
	"github.com/vothanhdo2602/hicon/external/config"
	"github.com/vothanhdo2602/hicon/external/util/log"
	"sync"
	"time"
)

type HMap struct {
	mu   sync.RWMutex
	data map[string]map[string]time.Duration
}

func (m *HMap) Add(key, field string, dur time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.data[key]; ok {
		if _, ok = m.data[key][field]; !ok {
			m.data[key][field] = dur
		}
	} else {
		m.data[key] = map[string]time.Duration{field: dur}
	}
}

// Delete xóa một key từ map
func (m *HMap) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, exists := m.data[key]
	if exists {
		delete(m.data, key)
		clear(m.data)
	}
}

func (m *HMap) DeleteField(key, field string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, exists := m.data[key]
	if exists {
		if _, ok := m.data[key][field]; ok {
			delete(m.data[key], field)
			clear(m.data[key])
		}
	}
}

var (
	hSetMap = &HMap{data: make(map[string]map[string]time.Duration)}
)

func LazySet(ctx context.Context) {
	var (
		logger = log.WithCtx(ctx)
	)

	for {
		if err := config.ConfigurationUpdated(); err == nil && client != nil {
			pipe := client.Pipeline()

			for k, v := range hSetMap.data {
				for field, dur := range v {
					pipe.HExpire(ctx, k, dur, field)
				}

				hSetMap.Delete(k)
			}

			if _, pipeErr := pipe.Exec(ctx); err != nil {
				logger.Error(pipeErr.Error())
			}
		}

		time.Sleep(9 * time.Minute)
	}

}
