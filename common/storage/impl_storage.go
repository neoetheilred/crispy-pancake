package storage

import (
	"sort"
	"sync"
	"sync/atomic"
)

type Storage[T IStorable] struct {
	storage   map[int64]T
	mu        sync.Mutex
	idCounter atomic.Int64
}

func NewStorage[T IStorable]() IStorage[T] {
	return &Storage[T]{
		storage: map[int64]T{},
	}
}

func (s *Storage[T]) Add(value T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.idCounter.Add(1)
	s.storage[s.idCounter.Load()] = value
}

func (s *Storage[T]) Delete(value T) {
	s.DeleteById(value.GetID())
}

func (s *Storage[T]) DeleteById(id int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.storage[id]; ok {
		delete(s.storage, id)
	}
}

func (s *Storage[T]) Exists(value T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.storage[value.GetID()]
	return ok
}

func (s *Storage[T]) ExistsKey(id int64) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.storage[id]
	return ok
}

func (s *Storage[T]) Update(id int64, value T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.storage[id]; !ok {
		return false
	}
	value.SetID(id)
	s.storage[id] = value
	return true
}

func (s *Storage[T]) GetAll() []T {
	res := make([]T, 0, len(s.storage))
	for _, value := range s.storage {
		res = append(res, value)
	}
	sort.Slice(res, func(i, j int) bool { return res[i].GetID() < res[j].GetID() })
	return res
}
