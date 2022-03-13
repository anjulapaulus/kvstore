package memory

import (
	"errors"
	"sync"
	"time"
)

type memoryStore struct {
	name        string
	records     map[interface{}]Record
	expireAfter time.Duration
	mtx         sync.RWMutex
}

type Record struct {
	Key, Value interface{}
	createdAt  time.Time
}

type memOptions struct {
	expiry time.Duration
}

type memOps func(*memOptions)

func WithTTL(expiry time.Duration) memOps {
	return func(opts *memOptions) {
		opts.expiry = expiry
	}
}

func applyDefault() *memOptions {
	ops := new(memOptions)
	ops.expiry = 0
	return ops
}

func NewMemoryStore(name string, options ...memOps) *memoryStore {
	opts := applyDefault()
	for _, opt := range options {
		opt(opts)
	}
	m := &memoryStore{
		name:        name,
		records:     make(map[interface{}]Record),
		expireAfter: opts.expiry * time.Nanosecond,
	}

	if m.expireAfter > 0 {
		go m.cleanup()
	}

	return m
}

func (m *memoryStore) Name() string {
	return m.name
}

func (m *memoryStore) Set(key, value interface{}) error {
	r := Record{
		Key:       key,
		Value:     value,
		createdAt: time.Now(),
	}
	m.mtx.Lock()
	m.records[key] = r
	m.mtx.Unlock()

	return nil
}

func (m *memoryStore) Get(key interface{}) (interface{}, error) {
	m.mtx.RLock()
	v, ok := m.records[key]
	m.mtx.RUnlock()
	if ok {
		return v, nil
	}
	return nil, nil
}

func (m *memoryStore) Delete(key interface{}) error {
	m.mtx.Lock()
	delete(m.records, key)
	m.mtx.Unlock()
	return nil
}

func (m *memoryStore) CreatedAt(key interface{}) time.Time {
	m.mtx.RLock()
	v, ok := m.records[key]
	m.mtx.RUnlock()
	if ok {
		return v.createdAt
	}
	return time.Time{}
}

func (m *memoryStore) NewIterator() (*Iterator, error) {
	copyRecs := m.copy()
	if len(copyRecs) < 1 {
		return nil, errors.New("no records")
	}

	return &Iterator{
		records: copyRecs,
		index:   len(copyRecs),
	}, nil
}

func (m *memoryStore) copy() []Record {
	var recs []Record
	m.mtx.RLock()
	for _, v := range m.records {
		recs = append(recs, v)
	}
	m.mtx.RUnlock()
	return recs
}

func (m *memoryStore) cleanup() {
	records := m.copy()
	for _, v := range records {
		if time.Since(v.createdAt).Nanoseconds() > m.expireAfter.Nanoseconds() {
			m.Delete(v.Key)
		}
	}
}
