package syncx

import "sync"

type Map[K comparable, V any] struct {
	syncMap sync.Map
}

func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{}
}

func (m *Map[K, V]) Delete(key K) { m.syncMap.Delete(key) }

func (m *Map[K, V]) Load(key K) (value V, ok bool) {
	v, ok := m.syncMap.Load(key)
	if v == nil {
		return value, ok
	}

	return v.(V), ok
}

func (m *Map[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	v, loaded := m.syncMap.LoadAndDelete(key)
	if !loaded {
		return value, loaded
	}
	return v.(V), loaded
}

func (m *Map[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	a, loaded := m.syncMap.LoadOrStore(key, value)
	return a.(V), loaded
}

func (m *Map[K, V]) Range(f func(key K, value V) bool) {
	m.syncMap.Range(func(key, value any) bool { return f(key.(K), value.(V)) })
}

func (m *Map[K, V]) Store(key K, value V) { m.syncMap.Store(key, value) }
