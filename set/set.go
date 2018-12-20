package set

import "sync"

type Type interface{}
type Item Type

type ItemSet struct {
	items map[Item]bool
	lock  sync.RWMutex
}

func (s *ItemSet) Add(t Item) *ItemSet {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.items == nil {
		s.items = make(map[Item]bool)
	}
	_, ok := s.items[t]
	if !ok {
		s.items[t] = true
	}
	return s
}

func (s *ItemSet) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = make(map[Item]bool)
}

func (s *ItemSet) Remove(item Item) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, ok := s.items[item]
	if ok {
		delete(s.items, item)
	}
	return ok
}

func (s *ItemSet) Has(item Item) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	_, ok := s.items[item]
	return ok
}

func (s *ItemSet) Items() []Item {
	s.lock.RLock()
	defer s.lock.RUnlock()
	items := []Item{}
	for i := range s.items {
		items = append(items, i)
	}
	return items
}

func (s *ItemSet) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.items)
}
